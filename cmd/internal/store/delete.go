package store

import (
	"fmt"
	"os"
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

	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("invalid name")
	}

	return nil
}
