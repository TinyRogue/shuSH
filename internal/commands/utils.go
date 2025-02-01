package commands

import (
	"os/exec"
)

func hasProgram(program string) (*string, bool) {
	path, err := exec.LookPath(program)
	if err != nil {
		return nil, false
	}
	return &path, true
}
