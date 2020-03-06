package httpsproxy

import (
	"context"
	"fmt"
	"net/http"
	"runtime"

	"github.com/kuritka/gext/log"
	"github.com/kuritka/gsvc/services/httpsproxy/internal/controller"
)

type HttpsProxy struct {
	settings Settings
	ctx      context.Context
}

var logger = log.Log

func New(settings Settings, ctx context.Context) *HttpsProxy {
	proxy := new(HttpsProxy)
	proxy.settings = settings
	proxy.ctx = ctx
	return proxy
}

func (p *HttpsProxy) Run() error {
	runtime.GOMAXPROCS(4)
	controller.Startup(p.settings.DefaultHost.Host)
	logger.Info().Msgf("listening on :%d", p.settings.Port)
	err := http.ListenAndServeTLS(fmt.Sprintf("%s", p.settings.Port), p.settings.CertPath, p.settings.KeyPath, controller.Router())
	return err
}

func (p *HttpsProxy) Name() string {
	return "HTTPS Reverse Proxy"
}
