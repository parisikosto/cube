package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Create a new system user and grant sudo privileges",
	Long:  `Runs adduser to create a new Linux user and usermod -aG sudo to grant admin privileges.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, err := linux.PromptUsername()
		if err != nil || username == "" {
			ui.Error("No username provided")
			os.Exit(1)
		}

		ui.SubCommand(fmt.Sprintf("> Creating user %s...", username))
		if err := linux.CreateUser(username); err != nil {
			ui.Error(fmt.Sprintf("Error creating user: %v", err))
			os.Exit(1)
		}

		ui.SubCommand(fmt.Sprintf("> Granting sudo privileges to %s...", username))
		if err := linux.GrantAdminPrivileges(username); err != nil {
			ui.Error(fmt.Sprintf("Error granting privileges: %v", err))
			os.Exit(1)
		}

		ui.Success(fmt.Sprintf("User %s created successfully with sudo privileges!", username))
	},
}

func init() {
	rootCmd.AddCommand(createUserCmd)
}
