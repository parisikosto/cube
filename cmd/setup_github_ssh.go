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
	Short:   "Generate an SSH key pair for GitHub access",
	Long: `Generates an RSA 4096-bit SSH key pair for GitHub, prints the public key to add to your GitHub account,
and provides instructions to activate the ssh-agent and test the connection with ssh -T git@github.com.`,
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
