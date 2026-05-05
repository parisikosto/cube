package linux

import "github.com/parisikosto/cube/internal/ui"

// LinuxTips prints useful Linux system and SSH command tips.
func LinuxTips() {
	ui.SubCommand("System:")
	ui.Suggestion("  Check OS version:    " + ui.InlineCommand("$ cat /etc/*-release"))
	ui.Suggestion("  Check users list:    " + ui.InlineCommand("$ less /etc/passwd"))
	ui.Suggestion("  Check disk usage:    " + ui.InlineCommand("$ df -h"))
	ui.Suggestion("  Check memory usage:  " + ui.InlineCommand("$ free -h"))
	ui.Suggestion("  Check uptime:        " + ui.InlineCommand("$ uptime"))

	ui.SubCommand("\nUsers:")
	ui.Suggestion("  Who am I:            " + ui.InlineCommand("$ whoami"))
	ui.Suggestion("  My groups:           " + ui.InlineCommand("$ id -nG"))

	ui.SubCommand("\nSSH Keys:")
	ui.Suggestion("  List SSH keys:       " + ui.InlineCommand("$ ls -al ~/.ssh"))
	ui.Suggestion("  Start ssh-agent:     " + ui.InlineCommand(`$ eval "$(ssh-agent -s)"`))
	ui.Suggestion("  Add SSH private key: " + ui.InlineCommand("$ ssh-add ~/.ssh/id_rsa"))
	ui.Suggestion("  Start agent + add:   " + ui.InlineCommand(`$ eval "$(ssh-agent -s)" && ssh-add ~/.ssh/id_rsa`))
	ui.Suggestion("  Test GitHub SSH:     " + ui.InlineCommand("$ ssh -T git@github.com"))
}

// GitTips prints useful Git command tips.
func GitTips() {
	ui.SubCommand("Config:")
	ui.Suggestion("  List all config:     " + ui.InlineCommand("$ git config --list"))
	ui.Suggestion("  Set username:        " + ui.InlineCommand(`$ git config --global user.name "Your Name"`))
	ui.Suggestion("  Set email:           " + ui.InlineCommand(`$ git config --global user.email "you@example.com"`))
	ui.Suggestion("  Set default branch:  " + ui.InlineCommand("$ git config --global init.defaultBranch main"))

	ui.SubCommand("\nRepository:")
	ui.Suggestion("  Init repo:           " + ui.InlineCommand("$ git init"))
	ui.Suggestion("  Clone repo:          " + ui.InlineCommand("$ git clone <url>"))
	ui.Suggestion("  Check status:        " + ui.InlineCommand("$ git status"))
	ui.Suggestion("  View log:            " + ui.InlineCommand("$ git log --oneline"))

	ui.SubCommand("\nBranches:")
	ui.Suggestion("  List branches:       " + ui.InlineCommand("$ git branch -a"))
	ui.Suggestion("  Create branch:       " + ui.InlineCommand("$ git checkout -b <branch>"))
	ui.Suggestion("  Switch branch:       " + ui.InlineCommand("$ git checkout <branch>"))

	ui.SubCommand("\nChanges:")
	ui.Suggestion("  Stage all:           " + ui.InlineCommand("$ git add ."))
	ui.Suggestion("  Commit:              " + ui.InlineCommand(`$ git commit -m "message"`))
	ui.Suggestion("  Push:                " + ui.InlineCommand("$ git push origin <branch>"))
	ui.Suggestion("  Pull:                " + ui.InlineCommand("$ git pull"))
}
