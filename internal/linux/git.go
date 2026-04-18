package linux

import "fmt"

// InstallGit installs Git via apt.
func InstallGit() error {
	if err := AptUpdate(); err != nil {
		return err
	}

	return runCmdInteractive("$ sudo apt install git", "sudo", "apt", "install", "git")
}

// VerifyGit prints the installed Git version.
func VerifyGit() error {
	return runCmd("$ git --version", "git", "--version")
}

// SetupGit configures git global user.name and user.email.
func SetupGit(name, email string) error {
	if err := runCmd(`$ git config --global user.name "`+name+`"`, "git", "config", "--global", "user.name", name); err != nil {
		return err
	}

	return runCmd(`$ git config --global user.email "`+email+`"`, "git", "config", "--global", "user.email", email)
}

// VerifyGitConfig prints the current git global configuration.
func VerifyGitConfig() error {
	return runCmd("$ git config --list", "git", "config", "--list")
}

// SetupGitInteractive prompts the user for a name and email, confirms, and configures Git globally.
func SetupGitInteractive() error {
	for {
		name, err := PromptUsername()
		if err != nil || name == "" {
			return fmt.Errorf("no name provided")
		}

		email, err := PromptEmail()
		if err != nil || email == "" {
			return fmt.Errorf("no email provided")
		}

		printGitConfig(name, email)

		if err := ConfirmPrompt("Is this correct"); err == nil {
			if err := SetupGit(name, email); err != nil {
				return err
			}
			return VerifyGitConfig()
		}
	}
}

func printGitConfig(name, email string) {
	fmt.Printf("\n  name:  %s\n  email: %s\n\n", name, email)
}

// UninstallGit removes Git and its unused dependencies.
func UninstallGit() error {
	if err := runCmdInteractive("$ sudo apt remove git", "sudo", "apt", "remove", "git"); err != nil {
		return err
	}

	return runCmdInteractive("$ sudo apt autoremove", "sudo", "apt", "autoremove")
}
