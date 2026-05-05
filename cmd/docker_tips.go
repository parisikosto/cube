package cmd

import (
	"github.com/parisikosto/cube/internal/linux"
	"github.com/spf13/cobra"
)

var dockerTipsCmd = &cobra.Command{
	GroupID: "tips",
	Use:     "docker-tips",
	Short:   "Display useful Docker command tips",
	Long:    `Prints a categorized reference of useful Docker commands for containers, images, and pruning.`,
	Run: func(cmd *cobra.Command, args []string) {
		linux.DockerTips()
	},
}

func init() {
	rootCmd.AddCommand(dockerTipsCmd)
}
