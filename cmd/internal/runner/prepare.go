package runner

import (
	"github.com/nunokawa/gdeck/cmd/internal/env"
	"github.com/nunokawa/gdeck/cmd/internal/model"
	"github.com/nunokawa/gdeck/cmd/internal/request"
)

func PrepareRequest(
	req *model.Request,
	envs map[string]string,
	body string,
	headers []string,
) (*model.Request, error) {

	// Body上書き
	if body != "" {
		req.Body = body
	}

	// Header上書き
	if len(headers) > 0 {
		req.Headers = request.MergeHeaders(
			req.Headers,
			headers,
		)
	}

	var err error

	// URL置換
	req.URL, err = env.ReplaceEnv(req.URL, envs)
	if err != nil {
		return nil, err
	}

	// Body置換
	req.Body, err = env.ReplaceEnv(req.Body, envs)
	if err != nil {
		return nil, err
	}

	// Header置換
	for i, h := range req.Headers {
		req.Headers[i], err = env.ReplaceEnv(h, envs)
		if err != nil {
			return nil, err
		}
	}

	return req, nil
}