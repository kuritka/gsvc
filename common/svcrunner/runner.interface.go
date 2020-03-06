package svcrunner

import "context"

type Service interface {
	Run(ctx context.Context) error
	Name() string
}
