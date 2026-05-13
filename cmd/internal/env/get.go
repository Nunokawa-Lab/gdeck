package env

import "fmt"

func Get(key string) (string, error) {

	envs, err := LoadEnv()
	if err != nil {
		return "", err
	}

	val, ok := envs[key]
	if !ok {
		return "", fmt.Errorf("env not found: %s", key)
	}

	return val, nil
}
