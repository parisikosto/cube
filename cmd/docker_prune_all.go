package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var dockerPruneAllCmd = &cobra.Command{
	GroupID: "docker",
	Use:     "docker-prune-all",
	Short:   "Prune all unused Docker resources",
	Long:    `Removes all unused Docker resources by running docker system prune, container prune, volume prune, network prune, and image prune -a.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := linux.ConfirmPrompt("This will remove all unused Docker resources. Are you sure"); err != nil {
			return
		}

		ui.SubCommand("> Pruning all Docker resources...")

		if err := linux.PruneDocker(); err != nil {
			ui.Error(fmt.Sprintf("Error pruning Docker resources: %v", err))
			os.Exit(1)
		}

		ui.Success("All unused Docker resources removed!")
	},
}

func init() {
	rootCmd.AddCommand(dockerPruneAllCmd)
}
