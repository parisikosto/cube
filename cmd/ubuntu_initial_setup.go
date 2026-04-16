package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var ubuntuInitialSetupCmd = &cobra.Command{
	Use:   "ubuntu-initial-setup",
	Short: "Initial VPS setup as root [1] (Ubuntu 24.04.1 LTS)",
	Long:  `Runs the complete initial server setup as root: system update, new user creation with sudo privileges, and firewall configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Command("> Starting initial server setup...")

		if err := linux.InitialSetup(); err != nil {
			ui.Error(fmt.Sprintf("Setup failed: %v", err))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(ubuntuInitialSetupCmd)
}
