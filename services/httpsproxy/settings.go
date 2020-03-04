package httpsproxy

import "net/url"

type Settings struct{
	Port int
	DefaultHost url.URL
}
