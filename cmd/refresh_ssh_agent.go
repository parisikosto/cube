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
	Short: "Re-add your GitHub SSH key to the agent in the current session",
	Long: `Session-level fallback for when the ssh-agent has died mid-session.

Scans ~/.ssh/ for GitHub SSH keys (id_rsa_github_*), lets you select one if multiple exist,
and runs ssh-add to load it into the active ssh-agent.

Under normal circumstances this command is not needed: setup-github-ssh configures
~/.bashrc to start the agent automatically on every login, and ~/.ssh/config loads
the key on first use via AddKeysToAgent yes.

If the agent is not running, source ~/.bashrc (or open a new terminal) instead.`,
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
