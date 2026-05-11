package store

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Delete(name string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	if strings.Contains(name, "..") {
		return fmt.Errorf("invalid name")
	}
	// 拡張子を除いたファイル名を取り出す
	base := filepath.Base(name)
	ext := filepath.Ext(name)
    filename := strings.ReplaceAll(base, ext, "")
	path := filepath.Join(home, ".apictl", "requests", filepath.Dir(name), filename+".json")

	err = os.Remove(path)
	if err != nil {
		return fmt.Errorf("invalid name")
	}

	return nil
}
