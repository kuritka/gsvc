//DepResolver provides configuration for particular services.
package depresolver

import (
	"github.com/kuritka/gext/env"
	"github.com/kuritka/gsvc/services/httpsproxy"
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
	dr.httpsproxy.initOnce.Do(func() {
		env.MustGetStringFlagFromEnv("HTTPS_PROXY_PORT")
		env.MustGetStringFlagFromEnv("HTTPS_PROXY_DEFAULT_HOST")
	})
	return dr.httpsproxy.settings
}
