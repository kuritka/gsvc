//Package provides L7 HTTPS reverse Proxy. Default Host is used in case that host is not given, otherwise it takes host from url
//Each request is goes in standalone thread
//i.e. https://<address of reverse proxy>/ uses default url (i.e. https://ulozto.cz)
//i.e. https://<address of reverse proxy>/>/?host=forbiddenpage.com uses custom url (i.e. https://forbiddenpage.com)
package proxy

import (
	"crypto/tls"
	"github.com/kuritka/gext/log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type httpsproxy struct {
	host string
}

type webRequest struct {
	r      *http.Request
	w      http.ResponseWriter
	doneCh chan struct{}
	host   string
}

var (
	client http.Client

	logger = log.Log

	requestCh = make(chan *webRequest)

	//TLS error occurs when false. Requires real certificate
	transport = http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
)

func init() {
	http.DefaultClient = &http.Client{Transport: &transport}
	client = http.Client{Transport: &transport}
}

// Creates instance of HTTPS proxy
func NewHttpsProxy(defaultHost string) (p *httpsproxy) {
	return &httpsproxy{host: defaultHost}
}

//Proxy starts listening https requests
func (p *httpsproxy) Listen() {

	for request := range requestCh {

		//proxy listens requests from handler
		go processRequest(request)
	}
}

func (p *httpsproxy) Handle(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	host := request.URL.Query().Get("host")
	if host == "" {
		host = p.host
	}
	doneCh := make(chan struct{})
	//processRequest() will fill writer by response
	requestCh <- &webRequest{r: request, w: writer, doneCh: doneCh, host: host}
	//each request is done in standalone thread. This thread is waiting until done channel is filled
	//waits until proxy resend request to chosen app server
	//and resend response back or error occurs
	<-doneCh
}

