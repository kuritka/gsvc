package httpsproxy

import (
	"context"
	"fmt"
	"net/http"
)

type HttpsProxy struct {
	settings Settings
	ctx context.Context
}

func New(settings Settings, ctx context.Context) *HttpsProxy {
	proxy := new(HttpsProxy)
	proxy.settings = settings
	proxy.ctx = ctx
	return proxy
}


func (p *HttpsProxy) Run() error {


	err := http.ListenAndServeTLS(fmt.Sprintf(":%d",p.settings.Port), p.settings.CertPath, p.settings.KeyPath, nil)
	return err
}

func (p *HttpsProxy) Name() string {
	return "Https Reverse Proxy"
}
