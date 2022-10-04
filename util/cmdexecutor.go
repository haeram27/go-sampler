package util

import (
	"bufio"
	"os/exec"
	"strings"
)

type ErrExitCode struct {
	Exitcode int
	Stderr   []byte
}

// Get exit code from error result of (*exec.Cmd).Output()
func CmdExitCode(err error) (bool, ErrExitCode) {
	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			return true, ErrExitCode{Exitcode: e.ExitCode(), Stderr: e.Stderr}
		}
	}
	return false, ErrExitCode{}
}

// use (*exec.Cmd).StdoutPipe() to resolve multiline result
func CmdStdoutPipe(cmd *exec.Cmd) (string, error) {
	out, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	scbuffer := bufio.NewScanner(out)
	scanResult := make([]string, 0)
	for {
		scbuffer.Scan()
		if data := scbuffer.Text(); len(data) != 0 {
			scanResult = append(scanResult, data)
		} else {
			break
		}
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}

	result := strings.Join(scanResult, "\n")
	return result, err
}
