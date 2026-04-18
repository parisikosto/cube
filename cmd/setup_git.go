package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var setupGitCmd = &cobra.Command{
	Use:   "setup-git",
	Short: "Configure global Git user name and email",
	Long:  `Sets git config --global user.name and user.email, then prints the current git configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			ui.SubCommand("> Enter your Git user name:")
			name, err := linux.PromptUsername()
			if err != nil || name == "" {
				ui.Error("No name provided")
				os.Exit(1)
			}

			ui.SubCommand("> Enter your Git email:")
			email, err := linux.PromptEmail()
			if err != nil || email == "" {
				ui.Error("No email provided")
				os.Exit(1)
			}

			ui.Instruction(fmt.Sprintf("\n  name:  %s", name))
			ui.Instruction(fmt.Sprintf("  email: %s\n", email))

			if err := linux.ConfirmPrompt("Is this correct"); err == nil {
				ui.SubCommand("> Configuring Git...")
				if err := linux.SetupGit(name, email); err != nil {
					ui.Error(fmt.Sprintf("Error configuring Git: %v", err))
					os.Exit(1)
				}

				ui.SubCommand("> Verifying Git configuration...")
				if err := linux.VerifyGitConfig(); err != nil {
					ui.Error(fmt.Sprintf("Error verifying Git config: %v", err))
					os.Exit(1)
				}

				ui.Success("Git configured successfully!")
				break
			}

			ui.Warning("Let's try again...")
		}
	},
}

func init() {
	rootCmd.AddCommand(setupGitCmd)
}
