package env

import (
	"fmt"
	"regexp"
)

// 環境変数に置換
func ReplaceEnv(str string, envs map[string]string) (string, error) {

	regx := regexp.MustCompile(`{{(\w+)}}`)

	var err error
	result := regx.ReplaceAllStringFunc(str, func(match string) string {
		// FindStringSubmatchの戻り値：[]string{"{{TOKEN}}", "TOKEN",}
		key := regx.FindStringSubmatch(match)[1]
		val := envs[key]
		if val == "" {
			err = fmt.Errorf("env var not set: %s", key)
			return match
		}

		return val
	})

	if err != nil {
		return "", err
	}

	return result, nil
}
