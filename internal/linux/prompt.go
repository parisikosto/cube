package linux

import (
	"fmt"

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
