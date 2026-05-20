package env

import (
	"os"
	"path/filepath"
)

// envパス生成共通関数
func BuildEnvPath(name string) (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if name == "" {
		// 空の場合はデフォルトの.env
		return filepath.Join(home, ".gdeck", ".env"), nil
	}

	return filepath.Join(home, ".gdeck", "envs", name+".env"), nil
}
