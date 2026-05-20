package env

func Set(key string, value string, name string) error {

	path, err := BuildEnvPath(name)
	if err != nil {
		return err
	}

	// 既存の.envを取得
	envs, err := LoadEnv(path)

	if err != nil {
		return err
	}

	// 追加
	envs[key] = value

	return SaveEnv(path, envs)
}
