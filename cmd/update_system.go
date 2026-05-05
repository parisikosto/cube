package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var updateSystemCmd = &cobra.Command{
	GroupID: "system",
	Use:     "update-system",
	Short:   "Update and upgrade all system packages",
	Long:    `Runs apt update, apt upgrade -y, apt dist-upgrade -y, and apt autoremove -y to keep the system fully up to date.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.SubCommand("> Starting system update...")

		if err := linux.UpdateSystem(); err != nil {
			ui.Error(fmt.Sprintf("Error updating system: %v", err))
			os.Exit(1)
		}

		ui.Success("System update completed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(updateSystemCmd)
}
