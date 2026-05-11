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
	if strings.Contains(name, "..") {
		// 不正なディレクトリアクセスを防ぐ
		return nil, fmt.Errorf("invalid name")
	}
	// 拡張子を除いたファイル名を取り出す
	base := filepath.Base(name)
	ext := filepath.Ext(name)
    filename := strings.ReplaceAll(base, ext, "")
	path := filepath.Join(home, ".apictl", "requests", filepath.Dir(name), filename+".json")

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
