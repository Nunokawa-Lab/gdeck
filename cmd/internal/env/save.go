package env

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func SaveEnv(path string, envs map[string]string) error {

	var lines []string
	for key, value := range envs {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}

	content := strings.Join(lines, "\n")

	// ディレクトリ作成
	//（MkdirAllはなければ作成しないため、ディレクトリ存在チェックは冗長で不要）
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(content), 0644)
}
