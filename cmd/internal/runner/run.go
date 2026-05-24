package runner

import (
	"github.com/nunokawa/gdeck/cmd/internal/env"
	"github.com/nunokawa/gdeck/cmd/internal/httpclient"
	"github.com/nunokawa/gdeck/cmd/internal/model"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

type RunOptions struct {
	Timeout int
	EnvName string
	Body    string
	Headers []string
}

func Run(
	name string,
	options RunOptions,
) ([]model.RunResult, error) {

	// request読み込み
	requests, err := store.Load(name)
	if err != nil {
		return nil, err
	}

	// env読み込み
	path, err := env.BuildEnvPath(options.EnvName)
	if err != nil {
		return nil, err
	}

	envs, err := env.LoadEnv(path)
	if err != nil {
		return nil, err
	}

	var results []model.RunResult

	for _, req := range requests {

		// request準備
		req, err = PrepareRequest(
			req,
			envs,
			options.Body,
			options.Headers,
		)
		if err != nil {
			return nil, err
		}

		// option設定
		httpOptions := store.DefaultOptions()

		if options.Timeout != 0 {
			httpOptions.Timeout = options.Timeout
		}

		// 実行
		res, err := httpclient.Do(
			req.Method,
			req.URL,
			req.Body,
			req.Headers,
			httpOptions,
		)

		// 結果保存
		results = append(results, model.RunResult{
			Request:  req,
			Response: res,
			Error:    err,
		})
	}

	return results, nil
}