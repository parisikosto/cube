package linux

// AptUpdate runs apt update.
func AptUpdate() error {
	return runCmd("$ sudo apt update", "sudo", "apt", "update")
}

// AptUpgrade runs apt upgrade -y.
func AptUpgrade() error {
	return runCmd("$ sudo apt upgrade -y", "sudo", "apt", "upgrade", "-y")
}

// AptDistUpgrade runs apt dist-upgrade -y, upgrading packages that require dependency changes.
func AptDistUpgrade() error {
	return runCmd("$ sudo apt dist-upgrade -y", "sudo", "apt", "dist-upgrade", "-y")
}

// AptAutoremove runs apt autoremove -y.
func AptAutoremove() error {
	return runCmd("$ sudo apt autoremove -y", "sudo", "apt", "autoremove", "-y")
}

// UpdateSystem runs apt update, upgrade, dist-upgrade, and autoremove.
func UpdateSystem() error {
	if err := AptUpdate(); err != nil {
		return err
	}

	if err := AptUpgrade(); err != nil {
		return err
	}

	if err := AptDistUpgrade(); err != nil {
		return err
	}

	return AptAutoremove()
}
