package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cube",
	Short: "A lightweight Go CLI for automating Ubuntu VPS provisioning and secure initial server setup",
	Long: `Cube is a lightweight Go CLI for automating secure
Ubuntu VPS provisioning and initial server setup.

It streamlines the standard server bootstrapping process
by guiding you through a step-by-step workflow—from
root initialization to user-level configuration.

Cube automates essential tasks such as system updates,
user creation with sudo privileges, SSH hardening,
and basic security setup, helping you turn a fresh Ubuntu server
into a production-ready environment quickly and reliably.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
