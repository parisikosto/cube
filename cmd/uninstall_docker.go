package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var uninstallDockerCmd = &cobra.Command{
	Use:   "uninstall-docker",
	Short: "Uninstall Docker CE and remove unused dependencies",
	Long:  `Removes Docker CE with apt remove and cleans up unused dependencies with apt autoremove.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := linux.ConfirmPrompt("Are you sure you want to uninstall Docker"); err != nil {
			return
		}

		ui.SubCommand("> Uninstalling Docker...")
		if err := linux.UninstallDocker(); err != nil {
			ui.Error(fmt.Sprintf("Error uninstalling Docker: %v", err))
			os.Exit(1)
		}

		ui.Success("Docker uninstalled successfully!")
	},
}

func init() {
	rootCmd.AddCommand(uninstallDockerCmd)
}
