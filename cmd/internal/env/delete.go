package env

import "fmt"

func Delete(key string, name string) error {

	path, err := BuildEnvPath(name)
	if err != nil {
		return err
	}

	envs, err := LoadEnv(path)
	if err != nil {
		return err
	}

	if _, ok := envs[key]; !ok {
		return fmt.Errorf("env not found: %s", key)
	}

	// mapから消して再保存
	delete(envs, key)

	return SaveEnv(path, envs)
}
