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
	if strings.Contains(name, "..") {
		// 不正なディレクトリアクセスを防ぐ
		return fmt.Errorf("invalid name")
	}

	// パス生成
	path, err := BuildRequestPath(name)
	if err != nil {
		return fmt.Errorf("invalid name")
	}

	if strings.Contains(name, "/") {
		// 名称にパスが含まれる場合はそのパスもディレクトリ作成対象とする
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
	}

	// JSON化
	b, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, b, 0644)
}
