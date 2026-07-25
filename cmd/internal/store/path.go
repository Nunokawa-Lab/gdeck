package store

import (
	"os"
	"path/filepath"
	"strings"
)

// gdeckで使用するディレクトリを作成
func EnsureDirs() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	base := filepath.Join(home, ".gdeck")
    return os.MkdirAll(filepath.Join(base, "requests"), 0755)
}

func BuildRequestPath(name string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".gdeck", "requests")

	// 拡張子を除いたファイル名を取り出す
	base := filepath.Base(name)
	ext := filepath.Ext(name)
	filename := strings.TrimSuffix(base, ext)

	path := filepath.Join(dir, filepath.Dir(name), filename+".json")

	return path, nil
}
