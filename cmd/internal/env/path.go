package env

import (
	"os"
	"path/filepath"
)

// envパス生成共通関数
func EnvPath() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".gdeck", ".env"), nil
}
