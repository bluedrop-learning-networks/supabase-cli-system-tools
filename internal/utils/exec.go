package utils

import (
	"io"
	"os"
	"os/exec"
	"syscall"
)

func Exec(env []string, cmd []string, stdout, stderr io.Writer) error {
	envsWithSystem := append(os.Environ(), env...)

	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = envsWithSystem
	command.Stderr = os.Stderr
	command.Stdout = stdout
	err := command.Run()
	// Ignore the "broken pipe" error when the supabase cli process' stdout pipe is closed
	// Example: `supabase db dump | head -n 20`
	if e, ok := err.(*exec.ExitError); ok {
		if status, ok := e.Sys().(syscall.WaitStatus); ok {
			if status.Signal() == syscall.SIGPIPE {
				return nil
			}
		}
	}
	return nil
}
