package cmd

import (
	"fmt"
	"os"

	"github.com/parisikosto/cube/internal/linux"
	"github.com/parisikosto/cube/internal/ui"
	"github.com/spf13/cobra"
)

var dockerAddUserCmd = &cobra.Command{
	GroupID: "docker",
	Use:     "docker-add-user",
	Short:   "Add the current user to the docker group",
	Long:    `Adds the current user to the docker group with usermod -aG docker. Log out and back in for the change to take effect.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.SubCommand("> Adding user to docker group...")

		user, err := linux.AddUserToDockerGroup()
		if err != nil {
			ui.Error(fmt.Sprintf("Error adding user to docker group: %v", err))
			os.Exit(1)
		}

		ui.Success(fmt.Sprintf("User '%s' added to docker group!", user))
		ui.Instruction("The change takes effect on your next login session.")
		ui.Instruction("After your next login, verify with: " + ui.InlineCommand("$ docker run hello-world"))
	},
}

func init() {
	rootCmd.AddCommand(dockerAddUserCmd)
}
