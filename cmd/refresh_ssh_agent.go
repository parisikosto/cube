package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var refreshSSHAgentCmd = &cobra.Command{
	GroupID: "git",
	Use:     "refresh-ssh-agent",
	Short:   "Add your GitHub SSH key to the running ssh-agent",
	Long: `Scans ~/.ssh/ for GitHub SSH keys (id_rsa_github_*), lets you select one if multiple exist,
and runs ssh-add to load it into the active ssh-agent session.

If the agent is not running, the ready-to-copy command is printed instead.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.SubCommand("> Loading SSH key into agent...")

		if err := linux.RefreshSSHAgent(); err != nil {
			ui.Error(fmt.Sprintf("Failed to refresh ssh-agent: %v", err))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(refreshSSHAgentCmd)
}
