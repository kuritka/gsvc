//DepResolver provides configuration for particular services.
package depresolver

import (
	"github.com/kuritka/gext/env"
	"github.com/kuritka/gext/guard"
	"github.com/kuritka/gsvc/services/httpsproxy"
	"net/url"
	"sync"
)

type DepResolver struct {
	httpsproxy struct {
		initOnce sync.Once
		settings httpsproxy.Settings
	}
}

func New() *DepResolver {
	dr := new(DepResolver)
	return dr
}

func (dr *DepResolver) MustResolveHttpsProxy() httpsproxy.Settings {
	var err error
	dr.httpsproxy.initOnce.Do(func() {
		dr.httpsproxy.settings.Port = env.MustGetStringFlagFromEnv("HTTPS_PROXY_PORT")
		dr.httpsproxy.settings.DefaultHost, err =  url.Parse(env.MustGetStringFlagFromEnv("HTTPS_PROXY_DEFAULT_HOST"))
		guard.FailOnError(err, "parsing HTTPS_PROXY_DEFAULT_HOST")
		dr.httpsproxy.settings.CertPath = env.MustGetStringFlagFromEnv("HTTPS_PROXY_CERT_PATH")
		dr.httpsproxy.settings.KeyPath = env.MustGetStringFlagFromEnv("HTTPS_PROXY_KEY_PATH")
	})
	return dr.httpsproxy.settings
}
