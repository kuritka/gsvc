package svcrunner

import (
	"context"
	"github.com/kuritka/gext/guard"
	"github.com/kuritka/gext/log"
)

type ServiceRunner struct {
	service Service
	ctx context.Context
}

var logger = log.Log


func New(service Service, ctx context.Context) *ServiceRunner {
	return &ServiceRunner{
		service,
		ctx,
	}
}


//Run service once and panics if service is broken
func (r *ServiceRunner) MustRun() {
	logger.Info().Msgf("service %s started", r.service.Name())
	err := r.service.Run(r.ctx)
	guard.FailOnError(err, "service %s failed", r.service.Name())
}