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
