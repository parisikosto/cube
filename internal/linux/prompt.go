package linux

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
)

// ConfirmPrompt asks the user whether to continue. Returns nil if confirmed, error if declined.
func ConfirmPrompt(label string) error {
	if label == "" {
		label = "Do you want to continue"
	}

	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		fmt.Println("Stopped.")
		return err
	}

	return nil
}

// PromptEmail interactively asks for a valid email address.
func PromptEmail() (string, error) {
	const pattern = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)

	prompt := promptui.Prompt{
		Label: "email",
		Validate: func(input string) error {
			trimmed := strings.TrimSpace(input)
			if len(trimmed) == 0 {
				return errors.New("email cannot be empty")
			}
			if !re.MatchString(trimmed) {
				return errors.New("invalid email address")
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
