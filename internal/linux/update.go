package linux

// AptUpdate runs apt update.
func AptUpdate() error {
	return runCmd("$ sudo apt update", "sudo", "apt", "update")
}

// AptUpgrade runs apt upgrade -y.
func AptUpgrade() error {
	return runCmd("$ sudo apt upgrade -y", "sudo", "apt", "upgrade", "-y")
}

// AptAutoremove runs apt autoremove -y.
func AptAutoremove() error {
	return runCmd("$ sudo apt autoremove -y", "sudo", "apt", "autoremove", "-y")
}

// UpdateSystem runs apt update, upgrade, and autoremove.
func UpdateSystem() error {
	if err := AptUpdate(); err != nil {
		return err
	}

	if err := AptUpgrade(); err != nil {
		return err
	}

	return AptAutoremove()
}
