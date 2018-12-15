package zengzhihai_pipeline

import (
	"errors"

	pf "github.com/easierway/pipefiter_framework/pipefilter"
)

var NotFoundPipelineError = errors.New("can't find the pipeline definition")

func CreatePipeline(pipelineName string) (*pf.StraightPipeline, error) {

	return new(pf.StraightPipeline), NotFoundPipelineError
}
