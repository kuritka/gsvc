package svcrunner

import (
	"github.com/kuritka/gext/guard"
	"github.com/kuritka/gext/log"
)

type ServiceRunner struct {
	service Service
}

var logger = log.Log


func New(service Service) *ServiceRunner {
	return &ServiceRunner{
		service,
	}
}


//Run service once and panics if service is broken
func (r *ServiceRunner) MustRun() {
	logger.Info().Msgf("service %s started", r.service.Name())
	err := r.service.Run()
	guard.FailOnError(err, "service %s failed", r.service.Name())
}