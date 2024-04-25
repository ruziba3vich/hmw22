package repo

import (
	"os/exec"
	"strings"
)

func GetUserEmail() (string, error) {
	cmd := exec.Command("git", "config", "--global", "user.email")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	email := strings.TrimSpace(string(output))
	return email, nil
}
