package store

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Delete(name string) error {
	if strings.Contains(name, "..") {
		// 不正なディレクトリアクセスを防ぐ
		return fmt.Errorf("invalid name")
	}

	// パス生成
	path, err := BuildRequestPath(name)
	if err != nil {
		return fmt.Errorf("invalid name")
	}

	if strings.Contains(name, "*") {
		// ディレクトリ内全件処理（ワイルドカードだったパスから各ファイルパスのスライスを取得）
		paths, err := filepath.Glob(path)
		if err != nil {
			return fmt.Errorf("invalid path")
		}

		for _, p := range paths {
			err = os.Remove(p)
			if err != nil {
				return fmt.Errorf("request not found: %s", p)
			}
		}

	} else {
		err = os.Remove(path)
		if err != nil {
			return fmt.Errorf("invalid name")
		}
	}

	return nil
}
