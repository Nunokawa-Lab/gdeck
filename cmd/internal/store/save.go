package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 入力されたコマンドの保存処理
func Save(name string, req interface{}) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dir := filepath.Join(home, ".apictl", "requests")
	if strings.Contains(name, "/") {
		// 名称にパスが含まれる場合はそのパスもディレクトリ作成対象とする
		dir = filepath.Join(dir, filepath.Dir(name))
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// ファイルパス生成
	if strings.Contains(name, "..") {
		// 不正なディレクトリアクセスを防ぐ
		return fmt.Errorf("invalid name")
	}

	// 拡張子を除いたファイル名を取り出す
	base := filepath.Base(name)
	ext := filepath.Ext(name)
    filename := strings.ReplaceAll(base, ext, "")
	// ファイル名として非推奨なドット除去
    filename = strings.ReplaceAll(filename, ".", "")

	path := filepath.Join(dir, filename+".json")

	// JSON化
	b, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, b, 0644)
}
