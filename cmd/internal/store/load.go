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
func Load(name string) ([]*model.Request, error) {
	if strings.Contains(name, "..") {
		// 不正なディレクトリアクセスを防ぐ
		return nil, fmt.Errorf("invalid name")
	}

	var path string
	var results []*model.Request

	// 拡張子が付いていても実行できるように、拡張子を除いたファイル名を取り出す処理を行う
	path, err := BuildRequestPath(name)
	if err != nil {
		return nil, fmt.Errorf("invalid path")
	}

	if strings.Contains(name, "*") {
		// ディレクトリ内全件処理（ワイルドカードだったパスから各ファイルパスのスライスを取得）
		paths, err := filepath.Glob(path)
		if err != nil {
			return nil, fmt.Errorf("invalid path")
		}
		
		for _, p := range paths {
			results, err = appendResults(p, results)
			if err != nil {
				return nil, fmt.Errorf("request not found: %s", p)
			}
		}

	} else {
		// 単体処理
		results, err = appendResults(path, results)
		if err != nil {
			return nil, fmt.Errorf("request not found: %s", path)
		}
	}

	return results, nil
}

func appendResults(path string, results []*model.Request) ([]*model.Request, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var req model.Request
	if err := json.Unmarshal(b, &req); err != nil {
		return nil, err
	}

	results = append(results, &req)

	return results, nil
}
