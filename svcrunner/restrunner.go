package svcrunner

import "net"

type RestRunner struct {
	listenerFactory func() (net.Listener, error)
}

