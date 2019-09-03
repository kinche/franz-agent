package system

import "os/exec"

// Helper functions for OS info

// Uname runs uname -a command
func Uname() string {
	e := exec.Command("uname", "-a")
	s, err := e.Output()
	if err != nil {
		return ""
	}
	return string(s)
}
