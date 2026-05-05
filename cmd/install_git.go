package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var installGitCmd = &cobra.Command{
	GroupID: "git",
	Use:     "install-git",
	Short:   "Install Git version control system",
	Long:    `Installs Git via apt and verifies the installation by printing the installed version.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.SubCommand("> Installing Git...")

		if err := linux.InstallGit(); err != nil {
			ui.Error(fmt.Sprintf("Error installing Git: %v", err))
			os.Exit(1)
		}

		ui.SubCommand("> Verifying Git installation...")
		if err := linux.VerifyGit(); err != nil {
			ui.Error(fmt.Sprintf("Error verifying Git: %v", err))
			os.Exit(1)
		}

		ui.Success("Git installed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(installGitCmd)
}
