package env

import (
	"fmt"
	"os"
	"strings"
)

func SaveEnv(envs map[string]string) error {

	path, err := EnvPath()
	if err != nil {
		return err
	}

	var lines []string
	for key, value := range envs {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}

	content := strings.Join(lines, "\n")

	return os.WriteFile(path, []byte(content), 0644)
}
