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
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// ファイルパス生成
	if strings.Contains(name, "..") || strings.Contains(name, "/") {
		// 不正なディレクトリアクセスを防ぐ
		return fmt.Errorf("invalid name")
	}
	path := filepath.Join(dir, name+".json")

	// JSON化
	b, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, b, 0644)
}
