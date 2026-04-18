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
		ui.SubCommand("> Configuring Git...")

		if err := linux.SetupGitInteractive(); err != nil {
			ui.Error(fmt.Sprintf("Error configuring Git: %v", err))
			os.Exit(1)
		}

		ui.Success("Git configured successfully!")
	},
}

func init() {
	rootCmd.AddCommand(setupGitCmd)
}
