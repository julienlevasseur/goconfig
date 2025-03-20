package command

import (
	"log"
	"os/exec"
	"strings"

	"github.com/julienlevasseur/goconfig/pkg/user"
)

// type CommandOptions struct {
// 	User *user.User
// 	// Debug *bool
// }

// func Exec(command string, args ...string) error {
func Exec(command string, args *[]string) error {
	// if opts != nil {
	// 	if opts.User != nil {
	// 		command = "runuser"
	// 		optsArgs := []string{
	// 			"-u",
	// 			opts.User.Username,
	// 			command,
	// 		}
	// 		optsArgs = append(optsArgs, args...)
	// 	}
	// }

	cmd := exec.Command(command, *args...)

	return cmd.Run()
}

func ExecWithOutput(command string, args *[]string) (string, error) {
	// if opts != nil {
	// 	if opts.User != nil {
	// 		command = "runuser"
	// 		optsArgs := []string{
	// 			"-u",
	// 			opts.User.Username,
	// 			command,
	// 		}
	// 		optsArgs = append(optsArgs, args...)
	// 	}
	// }

	cmd := exec.Command(command, *args...)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func ExecAs(u *user.User, command string, args *[]string) (string, error) {
	cmdArgs := []string{
		"-c",
		"sudo",
		"-H",
		"-u",
		u.Username,
		"bash",
		"-c",
		"-i",
		"'" + command,
	}
	cmdArgs = append(cmdArgs, *args...)

	var arguments string
	arguments = strings.Join(cmdArgs, " ")

	cmd := exec.Command(
		"/bin/sh",
		arguments+"'",
	)

	log.Printf("[DEBUG] %v %v\n", "sudo", arguments+"'")
	out, err := cmd.Output()
	log.Printf("[DEBUG] %v\n", out)
	log.Printf("[DEBUG] %v\n", err)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
