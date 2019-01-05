package zengzhihai_pipeline

import (
	"errors"

	pf "github.com/easierway/pipefiter_framework/pipefilter"
)

const LoginSso string = "loginSso"

var NotFoundPipelineError = errors.New("can't find the pipeline definition")

func CreatePipeline(pipelineName string) (*pf.StraightPipeline, error) {

	if pipelineName == LoginSso {
		filters := []pf.Filter{new(RequestFilter), new(LoginSsoFilter)}
		loginSso := pf.StraightPipeline{
			Name:    LoginSso,
			Filters: &filters,
		}
		return &loginSso, nil
	}

	return new(pf.StraightPipeline), NotFoundPipelineError
}
