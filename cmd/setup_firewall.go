package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var setupFirewallCmd = &cobra.Command{
	GroupID: "system",
	Use:     "setup-firewall",
	Short:   "Configure and enable the UFW firewall",
	Long:    `Allows OpenSSH through the firewall, enables UFW, and displays the current firewall status.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.SubCommand("> Checking registered firewall apps...")
		if err := linux.ListFirewallApps(); err != nil {
			ui.Warning(fmt.Sprintf("Warning: could not list firewall apps: %v", err))
		}

		ui.SubCommand("> Allowing SSH connections...")
		if err := linux.AllowSSH(); err != nil {
			ui.Error(fmt.Sprintf("Error allowing SSH: %v", err))
			os.Exit(1)
		}

		ui.SubCommand("> Enabling firewall...")
		if err := linux.EnableFirewall(); err != nil {
			ui.Error(fmt.Sprintf("Error enabling firewall: %v", err))
			os.Exit(1)
		}

		ui.Success("Firewall configured successfully!")
	},
}

func init() {
	rootCmd.AddCommand(setupFirewallCmd)
}
