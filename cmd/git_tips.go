package cmd

import (
	"github.com/parisikosto/cube/internal/linux"
	"github.com/spf13/cobra"
)

var gitTipsCmd = &cobra.Command{
	GroupID: "tips",
	Use:     "git-tips",
	Short:   "Display useful Git command tips",
	Long:    `Prints a categorized reference of useful Git commands for config, repository management, branches, and changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		linux.GitTips()
	},
}

func init() {
	rootCmd.AddCommand(gitTipsCmd)
}
