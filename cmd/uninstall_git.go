package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var uninstallGitCmd = &cobra.Command{
	Use:   "uninstall-git",
	Short: "Uninstall Git and remove unused dependencies",
	Long:  `Removes Git with apt remove and cleans up unused dependencies with apt autoremove.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := linux.ConfirmPrompt("Are you sure you want to uninstall Git"); err != nil {
			return
		}

		ui.SubCommand("> Uninstalling Git...")
		if err := linux.UninstallGit(); err != nil {
			ui.Error(fmt.Sprintf("Error uninstalling Git: %v", err))
			os.Exit(1)
		}

		ui.Success("Git uninstalled successfully!")
	},
}

func init() {
	rootCmd.AddCommand(uninstallGitCmd)
}
