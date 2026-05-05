package linux

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
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

// FindGithubSSHKeys scans ~/.ssh/ for private keys matching the id_rsa_github_* naming convention.
func FindGithubSSHKeys() ([]string, error) {
	u, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("could not get current user: %w", err)
	}

	matches, err := filepath.Glob(u.HomeDir + "/.ssh/id_rsa_github_*")
	if err != nil {
		return nil, err
	}

	var keys []string
	for _, f := range matches {
		if !strings.HasSuffix(f, ".pub") {
			if _, err := os.Stat(f + ".pub"); err == nil {
				keys = append(keys, f)
			}
		}
	}

	return keys, nil
}

// SelectSSHKey prompts the user to select a key from a list.
// If only one key exists, it is returned automatically.
func SelectSSHKey(keys []string) (string, error) {
	if len(keys) == 1 {
		return keys[0], nil
	}

	prompt := promptui.Select{
		Label: "Select SSH key to add",
		Items: keys,
	}

	_, selected, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("selection cancelled: %w", err)
	}

	return selected, nil
}

// startSSHAgent starts a new ssh-agent process, parses its socket path,
// and sets SSH_AUTH_SOCK in the current process. Returns the socket path.
func startSSHAgent() (string, error) {
	out, err := exec.Command("ssh-agent", "-s").Output()
	if err != nil {
		return "", fmt.Errorf("could not start ssh-agent: %w", err)
	}

	for _, line := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(line, "SSH_AUTH_SOCK=") {
			// line format: SSH_AUTH_SOCK=/tmp/ssh-.../agent.XXXXX; export SSH_AUTH_SOCK;
			sockPath := strings.SplitN(line, "=", 2)[1]
			sockPath = strings.SplitN(sockPath, ";", 2)[0]
			os.Setenv("SSH_AUTH_SOCK", strings.TrimSpace(sockPath))
			return sockPath, nil
		}
	}

	return "", fmt.Errorf("could not parse ssh-agent output")
}

// RefreshSSHAgent auto-detects GitHub SSH keys, lets the user select one,
// and loads it into the ssh-agent. Starts the agent automatically if needed.
func RefreshSSHAgent() error {
	keys, err := FindGithubSSHKeys()
	if err != nil || len(keys) == 0 {
		ui.Warning("No GitHub SSH keys found in ~/.ssh/")
		ui.Instruction("Generate one first with: " + ui.InlineCommand("$ cube setup-github-ssh"))
		return fmt.Errorf("no keys found")
	}

	keyPath, err := SelectSSHKey(keys)
	if err != nil {
		return err
	}

	ui.Instruction(fmt.Sprintf("  Key: %s", keyPath))

	// Step 1: try ssh-add directly (works if agent is already running)
	if err := runCmd("$ ssh-add "+keyPath, "ssh-add", keyPath); err == nil {
		ui.Success("Key added to ssh-agent!")
		ui.Instruction("Test the connection: " + ui.InlineCommand("$ ssh -T git@github.com"))
		return nil
	}

	// Step 2: agent not running or socket stale — start a fresh one
	ui.Warning("ssh-agent not available. Starting a new one...")
	agentSock, err := startSSHAgent()
	if err != nil {
		ui.Instruction("Start the agent and add the key manually:")
		ui.Instruction("  " + ui.InlineCommand(fmt.Sprintf(`$ eval "$(ssh-agent -s)" && ssh-add %s`, keyPath)))
		return err
	}

	// Step 3: retry ssh-add with the new agent socket
	if err := runCmd("$ ssh-add "+keyPath, "ssh-add", keyPath); err != nil {
		ui.Instruction("Add the key manually:")
		ui.Instruction("  " + ui.InlineCommand(fmt.Sprintf(`$ eval "$(ssh-agent -s)" && ssh-add %s`, keyPath)))
		return err
	}

	ui.Success("Key added to ssh-agent!")
	ui.Warning("A new agent was started. To use it in your current shell, run:")
	ui.Instruction("  " + ui.InlineCommand(fmt.Sprintf("$ export SSH_AUTH_SOCK=%s", agentSock)))
	ui.Instruction("Then test: " + ui.InlineCommand("$ ssh -T git@github.com"))

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
