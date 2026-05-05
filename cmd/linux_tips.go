package cmd

import (
	"github.com/parisikosto/cube/internal/linux"
	"github.com/spf13/cobra"
)

var linuxTipsCmd = &cobra.Command{
	GroupID: "tips",
	Use:     "linux-tips",
	Short:   "Display useful Linux system and SSH command tips",
	Long:    `Prints a categorized reference of useful Linux commands for system info, user management, and SSH key management.`,
	Run: func(cmd *cobra.Command, args []string) {
		linux.LinuxTips()
	},
}

func init() {
	rootCmd.AddCommand(linuxTipsCmd)
}
