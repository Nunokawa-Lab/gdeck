package env

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnv(path string) (map[string]string, error) {

	envs := make(map[string]string)

	file, err := os.Open(path)
	if os.IsNotExist(err) {
		// ファイル自体が存在しない場合は空map返す
		return envs, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// scannerで一行ずつ処理（メモリリーク対策）
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// 空行は処理しない
			continue
		}

		if strings.HasPrefix(line, "#") {
			// コメント行は処理しない
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			// key=valueになっていない不正行は処理しない
			continue
		}

		key := parts[0]
		value := parts[1]

		envs[key] = value
	}

	return envs, scanner.Err()
}
