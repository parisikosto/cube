package linux

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

// UninstallGit removes Git and its unused dependencies.
func UninstallGit() error {
	if err := runCmdInteractive("$ sudo apt remove git", "sudo", "apt", "remove", "git"); err != nil {
		return err
	}

	return runCmdInteractive("$ sudo apt autoremove", "sudo", "apt", "autoremove")
}
