package linux

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/parisikosto/cube/internal/ui"
)

// CreateUser runs adduser <username>
func CreateUser(username string) error {
	return runCmdInteractive("$ adduser "+username, "adduser", username)
}

// GrantAdminPrivileges runs usermod -aG sudo <username>
func GrantAdminPrivileges(username string) error {
	return runCmd("$ usermod -aG sudo "+username, "usermod", "-aG", "sudo", username)
}

// PromptUsername interactively asks for a valid username
func PromptUsername() (string, error) {
	const pattern = `^[a-z_][a-z0-9_-]*$`
	re := regexp.MustCompile(pattern)

	ui.SubCommand("Enter a username for the new system user:")

	prompt := promptui.Prompt{
		Label: "username",
		Validate: func(input string) error {
			trimmed := strings.TrimSpace(input)
			if len(trimmed) == 0 {
				return errors.New("username cannot be empty")
			}
			if !re.MatchString(trimmed) {
				return fmt.Errorf("invalid username: must match %s", pattern)
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(result), nil
}
