package store

import "github.com/Nunokawa-Lab/gdeck/cmd/internal/model"

// デフォルトタイムアウト設定処理
func DefaultOptions() model.Options {
	return model.Options{
		Timeout: 10,
	}
}
