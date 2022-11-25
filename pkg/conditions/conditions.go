package conditions

import (
	"os/exec"
)

func NotIf(command string) bool {
	//cmdAndArgs := strings.Split(command, " ")
	//fmt.Println(cmdAndArgs)
	//cmd := exec.Command(cmdAndArgs[0], cmdAndArgs[1:]...)
	cmd := exec.Command("bash", "-c", command)
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.Success()
		} else {
			panic(err)
		}
	}

	return true
}
