package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var setupGithubSSHCmd = &cobra.Command{
	GroupID: "git",
	Use:     "setup-github-ssh",
	Short: "Generate an SSH key pair for GitHub and configure automatic agent loading",
	Long: `Generates an RSA 4096-bit SSH key pair for GitHub, then:
  - Writes an ~/.ssh/config Host block with AddKeysToAgent yes so the key loads on first use
  - Appends an ssh-agent auto-start line to ~/.bashrc so the agent runs on every login
  - Prints the public key to add to your GitHub account

After running this command once you never need to manually run eval ssh-agent or ssh-add again.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Command("> Starting GitHub SSH key setup...")

		if err := linux.SetupGithubSSH(); err != nil {
			ui.Error(fmt.Sprintf("SSH setup failed: %v", err))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(setupGithubSSHCmd)
}
