package store

import (
	"os"
	"path/filepath"
)

// 保存されたコマンドファイルのリストを取得
func List() []string {
	
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	dir := filepath.Join(home, ".apictl", "requests")

	// ~/.apictl/requestからファイル取得
	// 今後、子階層にも保存できるようにしたら`filepath.Walk`で再起的に取得できるようにする
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}

	var filenames []string
	for _, file := range files {
		// ディレクトリはないはずだけど念の為判定
		if file.IsDir() {
			continue
		}
		filenames = append(filenames, file.Name())
	}
	
	return filenames
}