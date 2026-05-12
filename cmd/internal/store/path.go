package store

import (
	"os"
	"path/filepath"
	"strings"
)

func BuildRequestPath(name string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".apictl", "requests")

	// 拡張子を除いたファイル名を取り出す
	base := filepath.Base(name)
	ext := filepath.Ext(name)
	filename := strings.TrimSuffix(base, ext)

	path := filepath.Join(dir, filepath.Dir(name), filename+".json")

	return path, nil
}
