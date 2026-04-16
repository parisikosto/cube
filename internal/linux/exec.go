package linux

import (
	"os"
	"os/exec"

	"github.com/parisikosto/cube/internal/ui"
)

// runCmd prints the command label and executes it with live stdout/stderr output
func runCmd(label, name string, args ...string) error {
	ui.Command(label)
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
