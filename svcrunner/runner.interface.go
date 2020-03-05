package svcrunner

type Service interface {
	Run() error
	Name() string
}
