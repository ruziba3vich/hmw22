package repo

import (
    "os/exec"
    "strings"
)

func GetUserName() (string, error) {
    cmd := exec.Command("git", "config", "--global", "user.name")
    output, err := cmd.Output()
    if err != nil {
        return "", err
    }
    username := strings.TrimSpace(string(output))
    return username, nil
}
