package cmd

import (
	"fmt"

	"github.com/parisikosto/cube/build"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	GroupID: "info",
	Use:     "version",
	Short:   "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cube %s\n", build.Version)
		if build.GitCommit != "" {
			fmt.Printf("  commit:  %s\n", build.GitCommit)
		}
		if build.Time != "" {
			fmt.Printf("  built:   %s\n", build.Time)
		}
		if build.User != "" {
			fmt.Printf("  by:      %s\n", build.User)
		}
		if build.TargetOS != "" {
			fmt.Printf("  os:      %s\n", build.TargetOS)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
