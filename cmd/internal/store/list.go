package store

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// 保存されたコマンドファイルのリストを取得
func List() ([]string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("invalid name")
	}

	dir := filepath.Join(home, ".gdeck", "requests")

	var filenames []string
	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// ディレクトリはスキップ
		if d.IsDir() {
			return nil
		}

		// 相対パス化
		rel, err := filepath.Rel(dir, path)
		filenames = append(filenames, rel)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return filenames, nil
}
