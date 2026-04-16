package linux

// ListFirewallApps lists the UFW registered application profiles.
func ListFirewallApps() error {
	return runCmd("$ ufw app list", "ufw", "app", "list")
}

// AllowSSH allows OpenSSH connections through the firewall.
func AllowSSH() error {
	return runCmd("$ ufw allow OpenSSH", "ufw", "allow", "OpenSSH")
}

// EnableFirewall enables UFW and prints the current status.
func EnableFirewall() error {
	if err := runCmdInteractive("$ ufw enable", "ufw", "enable"); err != nil {
		return err
	}

	return runCmd("$ ufw status", "ufw", "status")
}
