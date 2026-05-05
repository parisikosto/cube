package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

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

// Execute is the entry point for the CLI.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddGroup(
		&cobra.Group{ID: "setup", Title: "Setup:"},
		&cobra.Group{ID: "system", Title: "System:"},
		&cobra.Group{ID: "git", Title: "Git:"},
		&cobra.Group{ID: "docker", Title: "Docker:"},
		&cobra.Group{ID: "tips", Title: "Tips:"},
		&cobra.Group{ID: "info", Title: "Info:"},
	)
}
