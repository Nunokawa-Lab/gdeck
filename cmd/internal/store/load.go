package store

import (
	"apictl/cmd/internal/model"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 保存されたコマンド情報を構造体に書き込む処理
func Load(name string) (*model.Request, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// ファイルパス生成
	if strings.Contains(name, "..") || strings.Contains(name, "/") {
		// 不正なディレクトリアクセスを防ぐ
		return nil, fmt.Errorf("invalid name")
	}
	path := filepath.Join(home, ".apictl", "requests", name+".json")

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("request not found: %s", name)
	}

	var req model.Request
	if err := json.Unmarshal(b, &req); err != nil {
		return nil, err
	}

	return &req, nil
}
