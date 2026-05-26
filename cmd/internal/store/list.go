package store

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/nunokawa/gdeck/cmd/internal/model"
)

// 保存されたコマンドファイルのリストを取得
func List() ([]model.RequestItem, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("invalid name")
	}

	dir := filepath.Join(home, ".gdeck", "requests")

	var requestItems []model.RequestItem
	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// ディレクトリはスキップ
		if d.IsDir() {
			return nil
		}

		// 隠しファイルはスキップ（主に .DS_Store）
		// Windows対応時は別の書き方にする必要あり
		if strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}

		// 相対パス化
		rel, err := filepath.Rel(dir, path)

		// メソッド取得
		loadFile, err := Load(rel)
		if err != nil {
			return err
		}
		method := loadFile[0].Method

		requestItems = append(
			requestItems,
			model.RequestItem{
				Method: method,
				Name:   rel,
			},
		)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return requestItems, nil
}
