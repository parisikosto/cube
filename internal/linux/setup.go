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

// StandardSetup runs the full standard server setup as the new user:
// system update → install git → setup git → install docker → add user to docker group → print next steps.
func StandardSetup() error {
	// Step 1: Update system
	ui.SubCommand("\n[Step 1/6] Updating system packages...")
	if err := UpdateSystem(); err != nil {
		return fmt.Errorf("system update failed: %w", err)
	}
	ui.Success("Step 1/6 complete.")

	// Step 2: Install Git
	ui.SubCommand("\n[Step 2/6] Installing Git...")
	if err := InstallGit(); err != nil {
		ui.Error(fmt.Sprintf("Failed to install Git: %v", err))
		ui.Suggestion("You can run this step separately later: cube install-git")

		if err := ConfirmPrompt("Skip and continue"); err != nil {
			return fmt.Errorf("setup stopped at Git installation")
		}
	} else {
		ui.Success("Step 2/6 complete.")
	}

	// Step 3: Setup Git
	ui.SubCommand("\n[Step 3/6] Configuring Git...")
	if err := SetupGitInteractive(); err != nil {
		ui.Error(fmt.Sprintf("Failed to configure Git: %v", err))
		ui.Suggestion("You can run this step separately later: cube setup-git")

		if err := ConfirmPrompt("Skip and continue"); err != nil {
			return fmt.Errorf("setup stopped at Git configuration")
		}
	} else {
		ui.Success("Step 3/6 complete.")
	}

	// Step 4: Install Docker
	ui.SubCommand("\n[Step 4/6] Installing Docker...")
	if err := InstallDocker(); err != nil {
		ui.Error(fmt.Sprintf("Failed to install Docker: %v", err))
		ui.Suggestion("You can run this step separately later: cube install-docker")

		if err := ConfirmPrompt("Skip and continue"); err != nil {
			return fmt.Errorf("setup stopped at Docker installation")
		}
	} else {
		ui.Success("Step 4/6 complete.")
	}

	// Step 5: Add user to Docker group
	ui.SubCommand("\n[Step 5/6] Adding user to Docker group...")
	user, err := AddUserToDockerGroup()
	if err != nil {
		ui.Error(fmt.Sprintf("Failed to add user to Docker group: %v", err))
		ui.Suggestion("You can run this step separately later: cube docker-add-user")

		if err := ConfirmPrompt("Skip and continue"); err != nil {
			return fmt.Errorf("setup stopped at Docker group configuration")
		}
	} else {
		ui.Success(fmt.Sprintf("Step 5/6 complete. User '%s' added to Docker group.", user))
	}

	// Step 6: Setup GitHub SSH key
	ui.SubCommand("\n[Step 6/6] Setting up GitHub SSH key...")
	if err := SetupGithubSSH(); err != nil {
		ui.Error(fmt.Sprintf("Failed to set up GitHub SSH key: %v", err))
		ui.Suggestion("You can run this step separately later: cube setup-github-ssh")

		if err := ConfirmPrompt("Skip and continue"); err != nil {
			return fmt.Errorf("setup stopped at GitHub SSH setup")
		}
	} else {
		ui.Success("Step 6/6 complete.")
	}

	printStandardNextSteps()

	return nil
}

func printStandardNextSteps() {
	ui.SubCommand("\n─────────────────────────────────────")
	ui.Instruction("Setup complete. Next steps:\n")
	ui.Instruction("  1. Log out and back in to apply Docker group changes:")
	ui.Instruction("     " + ui.InlineCommand("$ exit") + "\n")
	ui.Instruction("  2. After login, verify Docker works:")
	ui.Instruction("     " + ui.InlineCommand("$ docker run hello-world"))
	ui.SubCommand("\n─────────────────────────────────────")
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
