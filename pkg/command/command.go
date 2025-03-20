package command

import (
	"os/exec"

	"github.com/julienlevasseur/goconfig/pkg/user"
)

type CommandOptions struct {
	User *user.User
	// Debug *bool
}

// func Exec(command string, args ...string) error {
func Exec(command string, args []string, opts *CommandOptions) error {
	if opts != nil {
		if opts.User != nil {
			command = "runuser"
			optsArgs := []string{
				"-u",
				opts.User.Username,
				command,
			}
			optsArgs = append(optsArgs, args...)
		}
	}

	cmd := exec.Command(command, args...)

	return cmd.Run()
}

func ExecWithOutput(command string, args []string, opts *CommandOptions) (string, error) {
	if opts != nil {
		if opts.User != nil {
			command = "runuser"
			optsArgs := []string{
				"-u",
				opts.User.Username,
				command,
			}
			optsArgs = append(optsArgs, args...)
		}
	}

	cmd := exec.Command(command, args...)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
