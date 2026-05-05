package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var installDockerCmd = &cobra.Command{
	GroupID: "docker",
	Use:     "install-docker",
	Short:   "Install Docker CE from the official Docker repository",
	Long:    `Installs Docker CE by adding the official Docker GPG key and repository, then installs docker-ce and verifies the service is running.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.SubCommand("> Installing Docker CE...")

		if err := linux.InstallDocker(); err != nil {
			ui.Error(fmt.Sprintf("Error installing Docker: %v", err))
			os.Exit(1)
		}

		ui.Success("Docker CE installed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(installDockerCmd)
}
