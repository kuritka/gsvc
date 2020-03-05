package httpsproxy

import "net/url"

type Settings struct{
	Port int
	CertPath string
	KeyPath string
	DefaultHost url.URL
}
