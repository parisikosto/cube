package linux

import (
	"fmt"

	"github.com/parisikosto/cube/internal/ui"
)

// InitialSetup runs the full initial server setup as root:
// system update → create user → configure firewall → print next steps.
func InitialSetup() error {
	username := "<new-username>"

	// Step 1: Update system
	ui.SubCommand("\n[Step 1/3] Updating system packages...")
	if err := UpdateSystem(); err != nil {
		return fmt.Errorf("system update failed: %w", err)
	}
	ui.Success("Step 1/3 complete.")

	// Step 2: Create user
	ui.SubCommand("\n[Step 2/3] Creating system user...")
	input, err := PromptUsername()
	if err != nil || input == "" {
		ui.Error("Could not read username.")
		ui.Suggestion("You can run this step separately later: cube create-user")

		if err := ConfirmPrompt("Skip user creation and continue"); err != nil {
			return fmt.Errorf("setup stopped at user creation")
		}
	} else {
		username = input

		if err := CreateUser(username); err != nil {
			ui.Error(fmt.Sprintf("Failed to create user: %v", err))
			ui.Suggestion("You can run this step separately later: cube create-user")

			if err := ConfirmPrompt("Skip and continue"); err != nil {
				return fmt.Errorf("setup stopped at user creation")
			}
		} else if err := GrantAdminPrivileges(username); err != nil {
			ui.Error(fmt.Sprintf("Failed to grant sudo privileges: %v", err))
			ui.Suggestion("You can run this step separately later: cube create-user")

			if err := ConfirmPrompt("Skip and continue"); err != nil {
				return fmt.Errorf("setup stopped at granting privileges")
			}
		} else {
			ui.Success("Step 2/3 complete.")
		}
	}

	// Step 3: Firewall
	ui.SubCommand("\n[Step 3/3] Configuring firewall...")
	ListFirewallApps()

	if err := AllowSSH(); err != nil {
		ui.Error(fmt.Sprintf("Failed to allow SSH: %v", err))
		ui.Suggestion("You can run this step separately later: cube setup-firewall")

		if err := ConfirmPrompt("Skip and continue"); err != nil {
			return fmt.Errorf("setup stopped at firewall configuration")
		}
	} else if err := EnableFirewall(); err != nil {
		ui.Error(fmt.Sprintf("Failed to enable firewall: %v", err))
		ui.Suggestion("You can run this step separately later: cube setup-firewall")

		if err := ConfirmPrompt("Skip and continue"); err != nil {
			return fmt.Errorf("setup stopped at enabling firewall")
		}
	} else {
		ui.Success("Step 3/3 complete.")
	}

	printNextSteps(username)

	return nil
}

func printNextSteps(username string) {
	ip, err := GetServerIP()
	if err != nil {
		ip = "<server-ip>"
	}

	ui.SubCommand("\n─────────────────────────────────────")
	ui.Instruction("Setup complete. Next steps:\n")
	ui.Instruction("  1. Exit root session:")
	ui.Instruction("     " + ui.InlineCommand("$ exit") + "\n")
	ui.Instruction("  2. Connect as new user:")
	ui.Instruction("     " + ui.InlineCommand(fmt.Sprintf("$ ssh %s@%s", username, ip)))
	ui.SubCommand("\n─────────────────────────────────────")
}
