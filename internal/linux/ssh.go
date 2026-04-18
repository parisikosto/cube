package linux

import (
	"fmt"
	"os/user"

	"github.com/parisikosto/cube/internal/ui"
)

// ListSSHKeys prints the SSH keys in the current user's .ssh directory.
func ListSSHKeys() error {
	u, err := user.Current()
	if err != nil {
		return fmt.Errorf("could not get current user: %w", err)
	}

	return runCmd("$ ls -al ~/.ssh", "ls", "-lah", u.HomeDir+"/.ssh")
}

// GenerateSSHKeyForGithub runs ssh-keygen to create an RSA 4096-bit key for GitHub.
// The key is saved to ~/.ssh/id_rsa_github_<email>. Returns the key file path.
func GenerateSSHKeyForGithub(email string) (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("could not get current user: %w", err)
	}

	keyPath := u.HomeDir + "/.ssh/id_rsa_github_" + email
	ui.Instruction("  Key will be saved to: " + ui.InlineCommand(keyPath))

	err = runCmdInteractive(
		fmt.Sprintf(`$ ssh-keygen -t rsa -b 4096 -C "%s" -f %s`, email, keyPath),
		"ssh-keygen", "-t", "rsa", "-b", "4096", "-C", email, "-f", keyPath,
	)

	return keyPath, err
}

// PrintSSHPublicKey prints the contents of the public key file at the given path.
func PrintSSHPublicKey(keyPath string) error {
	return runCmd("$ cat "+keyPath+".pub", "cat", keyPath+".pub")
}

// SetupGithubSSH runs the full GitHub SSH key setup flow:
// lists existing keys, prompts for email, generates key pair, prints public key, and prints next steps.
func SetupGithubSSH() error {
	ui.SubCommand("> Listing existing SSH keys...")
	if err := ListSSHKeys(); err != nil {
		ui.Warning("No SSH keys found or ~/.ssh does not exist yet.")
	}

	if err := ConfirmPrompt("Continue with SSH key generation for GitHub"); err != nil {
		return fmt.Errorf("SSH key generation cancelled")
	}

	ui.SubCommand("> Enter your GitHub email:")
	email, err := PromptEmail()
	if err != nil || email == "" {
		return fmt.Errorf("no email provided")
	}

	ui.SubCommand("> Generating SSH key pair...")
	keyPath, err := GenerateSSHKeyForGithub(email)
	if err != nil {
		return fmt.Errorf("key generation failed: %w", err)
	}

	ui.SubCommand("> Your public key:")
	if err := PrintSSHPublicKey(keyPath); err != nil {
		return fmt.Errorf("could not print public key: %w", err)
	}

	printGithubSSHInstructions(keyPath)

	return nil
}

func printGithubSSHInstructions(keyPath string) {
	ip, err := GetServerIP()
	if err != nil {
		ip = "<server-ip>"
	}

	ui.SubCommand("\n─────────────────────────────────────")
	ui.Instruction("Copy the public key above and add it to GitHub:\n")
	ui.Instruction("  1. Go to GitHub → Settings → SSH and GPG keys → New SSH key")
	ui.Instruction(fmt.Sprintf("     Title: <username>@%s", ip))
	ui.Instruction("     Key:   paste the public key printed above\n")
	ui.Instruction("  2. Activate the ssh-agent and add the key:")
	ui.Instruction("     " + ui.InlineCommand(fmt.Sprintf(`$ eval "$(ssh-agent -s)" && ssh-add %s`, keyPath)) + "\n")
	ui.Instruction("  3. Test the connection:")
	ui.Instruction("     " + ui.InlineCommand("$ ssh -T git@github.com"))
	ui.SubCommand("\n─────────────────────────────────────")
}
