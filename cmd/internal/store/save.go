package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 入力されたコマンドの保存処理。
// 同じ名前のファイルが既にある場合は新規作成せず上書き更新する。
// updated が true のとき更新、false のとき新規作成。
func Save(name string, req interface{}) (updated bool, err error) {
	if strings.Contains(name, "..") {
		// 不正なディレクトリアクセスを防ぐ
		return false, fmt.Errorf("invalid name")
	}

	// パス生成
	path, err := BuildRequestPath(name)
	if err != nil {
		return false, fmt.Errorf("invalid name")
	}

	if _, statErr := os.Stat(path); statErr == nil {
		// 保存しようとしているパスにファイルがあれば更新
		updated = true
	} else if !os.IsNotExist(statErr) {
		// 存在しない以外のエラーの場合はエラー返却
		return false, statErr
	}
	// 存在しないというエラーならエラー返却せず新規作成対象とする

	if strings.Contains(name, "/") {
		// 名称にパスが含まれる場合はそのパスもディレクトリ作成対象とする
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return false, err
		}
	}

	// JSON化
	b, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return false, err
	}

	if err := os.WriteFile(path, b, 0644); err != nil {
		return false, err
	}

	return updated, nil
}
