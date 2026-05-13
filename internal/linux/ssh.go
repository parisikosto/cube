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

const sshAgentAutoStartLine = `[ -z "$SSH_AUTH_SOCK" ] && eval "$(ssh-agent -s)" > /dev/null 2>&1`

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

// WriteSSHConfig writes an ~/.ssh/config Host block for github.com that sets
// AddKeysToAgent yes so the key is loaded automatically on first use.
// If a github.com entry already exists the file is left untouched.
func WriteSSHConfig(keyPath string) error {
	u, err := user.Current()
	if err != nil {
		return fmt.Errorf("could not get current user: %w", err)
	}

	sshDir := filepath.Join(u.HomeDir, ".ssh")
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		return fmt.Errorf("could not create ~/.ssh: %w", err)
	}

	configPath := filepath.Join(sshDir, "config")
	content, _ := os.ReadFile(configPath)
	if strings.Contains(string(content), "Host github.com") {
		ui.Instruction("  ~/.ssh/config already has a github.com entry, skipping.")
		return nil
	}

	block := fmt.Sprintf("\nHost github.com\n    AddKeysToAgent yes\n    IdentityFile %s\n", keyPath)

	f, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("could not open ~/.ssh/config: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(block); err != nil {
		return fmt.Errorf("could not write ~/.ssh/config: %w", err)
	}

	ui.Success("~/.ssh/config updated with AddKeysToAgent yes for github.com")
	return nil
}

// WriteSSHAgentAutoStart appends a one-liner to ~/.bashrc that ensures
// ssh-agent is running silently on every new login session.
// If the line is already present the file is left untouched.
func WriteSSHAgentAutoStart() error {
	u, err := user.Current()
	if err != nil {
		return fmt.Errorf("could not get current user: %w", err)
	}

	bashrc := filepath.Join(u.HomeDir, ".bashrc")
	content, _ := os.ReadFile(bashrc)
	if strings.Contains(string(content), sshAgentAutoStartLine) {
		ui.Instruction("  ~/.bashrc already has ssh-agent auto-start, skipping.")
		return nil
	}

	f, err := os.OpenFile(bashrc, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open ~/.bashrc: %w", err)
	}
	defer f.Close()

	line := "\n# ssh-agent auto-start (added by cube)\n" + sshAgentAutoStartLine + "\n"
	if _, err := f.WriteString(line); err != nil {
		return fmt.Errorf("could not write to ~/.bashrc: %w", err)
	}

	ui.Success("~/.bashrc configured: ssh-agent will start automatically on login")
	return nil
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

	ui.SubCommand("> Configuring ~/.ssh/config...")
	if err := WriteSSHConfig(keyPath); err != nil {
		ui.Warning(fmt.Sprintf("Could not update ~/.ssh/config: %v", err))
	}

	ui.SubCommand("> Configuring ssh-agent auto-start in ~/.bashrc...")
	if err := WriteSSHAgentAutoStart(); err != nil {
		ui.Warning(fmt.Sprintf("Could not update ~/.bashrc: %v", err))
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

// RefreshSSHAgent is a session-level fallback that loads a GitHub SSH key into
// a running ssh-agent. Useful when the agent died mid-session.
// Under normal circumstances the agent is started automatically by ~/.bashrc
// and the key is loaded on first use via AddKeysToAgent in ~/.ssh/config.
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

	// Try ssh-add directly — works if agent is alive in the current session.
	// Uses runCmdInteractive so the user can enter a passphrase if the key is protected.
	if err := runCmdInteractive("$ ssh-add "+keyPath, "ssh-add", keyPath); err == nil {
		ui.Success("Key added to ssh-agent!")
		ui.Instruction("Test the connection: " + ui.InlineCommand("$ ssh -T git@github.com"))
		return nil
	}

	// Agent is not running. A child process cannot export env vars to the
	// parent shell, so the only reliable fix is to reload ~/.bashrc in the
	// current shell session.
	ui.Warning("ssh-agent is not running in this session.")
	ui.Instruction("Restart the agent by reloading your shell config:")
	ui.Instruction("  " + ui.InlineCommand("$ source ~/.bashrc"))
	ui.Instruction("Or open a new terminal — the agent starts automatically on login.")
	return fmt.Errorf("ssh-agent not available in current session")
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
	ui.Instruction("  2. Reload your shell to start the ssh-agent:")
	ui.Instruction("     " + ui.InlineCommand("$ source ~/.bashrc") + "\n")
	ui.Instruction("  3. Test the connection (the key loads automatically on first use):")
	ui.Instruction("     " + ui.InlineCommand("$ ssh -T git@github.com"))
	ui.Instruction("\n  From now on, every new terminal starts the agent automatically.")
	ui.SubCommand("\n─────────────────────────────────────")
}
