package linux

// UpdateSystem runs apt update, upgrade, and autoremove
func UpdateSystem() error {
	if err := runCmd("$ sudo apt update", "sudo", "apt", "update"); err != nil {
		return err
	}

	if err := runCmd("$ sudo apt upgrade -y", "sudo", "apt", "upgrade", "-y"); err != nil {
		return err
	}

	return runCmd("$ sudo apt autoremove -y", "sudo", "apt", "autoremove", "-y")
}
