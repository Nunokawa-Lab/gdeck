package env

func SetEnv(key string, value string) error {

	// 既存の.envを取得
	envs, err := LoadEnv()

	if err != nil {
		return err
	}

	// 追加
	envs[key] = value

	return SaveEnv(envs)
}
