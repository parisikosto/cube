package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var dockerCheckUserCmd = &cobra.Command{
	GroupID: "docker",
	Use:     "docker-check-user",
	Short:   "Verify the current user is in the docker group",
	Long:    `Prints the groups of the current user with id -nG. Run this after logging back in to confirm docker group membership.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.SubCommand("> Checking docker group membership...")

		if err := linux.CheckDockerGroup(); err != nil {
			ui.Error(fmt.Sprintf("Error checking docker group: %v", err))
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerCheckUserCmd)
}
