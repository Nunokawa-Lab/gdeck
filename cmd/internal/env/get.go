package env

import "fmt"

func Get(key string, name string) (string, error) {

	path, err := BuildEnvPath(name)
	if err != nil {
		return "", err
	}

	envs, err := LoadEnv(path)
	if err != nil {
		return "", err
	}

	val, ok := envs[key]
	if !ok {
		return "", fmt.Errorf("env not found: %s", key)
	}

	return val, nil
}
