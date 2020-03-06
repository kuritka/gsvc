package httpsproxy

import "net/url"

type Settings struct{
	//Port where https proxy is listening. Port is in format `:<portnum>`
	Port string
	//Certificate
	CertPath string
	//Key
	KeyPath string
	//Default host when you not explicitly specify host.
	DefaultHost *url.URL
}
