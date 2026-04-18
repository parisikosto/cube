package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var ubuntuStandardSetupCmd = &cobra.Command{
	Use:   "ubuntu-standard-setup",
	Short: "Standard VPS setup as new user [2] (Ubuntu 24.04.4 LTS)",
	Long:  `Runs the complete standard server setup as the new user: system update, Git installation and configuration, Docker installation, and Docker group setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Command("> Starting standard server setup...")

		if err := linux.StandardSetup(); err != nil {
			ui.Error(fmt.Sprintf("Setup failed: %v", err))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(ubuntuStandardSetupCmd)
}
