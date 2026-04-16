package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var timezoneCmd = &cobra.Command{
	Use:   "timezone",
	Short: "Get the current date in a given timezone",
	Long: `Displays the current date for the given timezone. Use the --date flag to customize the output format.

Example:
  cube timezone Europe/Athens
  cube timezone America/New_York --date "2006-01-02 15:04:05"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timezone := args[0]
		location, _ := time.LoadLocation(timezone)
		dateFlag, _ := cmd.Flags().GetString("date")
		var date string

		if dateFlag != "" {
			date = time.Now().In(location).Format(dateFlag)
		} else {
			date = time.Now().In(location).Format(time.RFC3339)[:10]
		}
		fmt.Printf("Current date in %v: %v\n", timezone, date)
	},
}

func init() {
	rootCmd.AddCommand(timezoneCmd)
	timezoneCmd.Flags().String("date", "", "Date for which to get the time (format: yyyy-mm-dd)")
}
