package proxy

import (
	"io"
	"net/http"
	"net/url"

	"github.com/kuritka/gext/concurency"
	"github.com/kuritka/gext/httphead"
)

func processRequest(inreq *webRequest) {

	defer func() {
		if err := recover(); err != nil {
			logger.Error().Msgf("recover from <%s>: %s", inreq.r.URL, err)
			inreq.w.WriteHeader(http.StatusInternalServerError)
			inreq.doneCh <- struct{}{}
		}
	}()

	//build url for new host
	hostUrl, _ := url.Parse(inreq.r.URL.String())
	hostUrl.Scheme = "https"
	hostUrl.Host = inreq.host

	outreq, _ := http.NewRequest(inreq.r.Method, hostUrl.String(), inreq.r.Body)

	//because inreq headers in go is map of slice of strings we must translate into string of headers to new inreq
	inheaders := httphead.HeaderAsMap(inreq.r.Header)
	for k, v := range inheaders {
		outreq.Header.Add(k, v)
	}

	//promise helps to unblock threads when waiting for response and makes transparent callback doom
	call(outreq).
		Then(func(obj interface{}) error { return response(obj, inreq) }, func(err error) { err500(inreq) }).
		Then(func(i interface{}) error {
			inreq.doneCh <- struct{}{}
			return nil
		}, func(err error) { err500(inreq) })

}

func call(r *http.Request) *concurency.Promise {
	result := new(concurency.Promise)
	result.SuccessChannel = make(chan interface{}, 1)
	result.ErrorChannel = make(chan error, 1)

	go func(r *http.Request) {
		resp, err := client.Do(r)
		if err != nil {
			result.ErrorChannel <- err
			return
		}
		result.SuccessChannel <- resp

	}(r)
	return result
}

func err500(inreq *webRequest) {
	inreq.w.WriteHeader(http.StatusInternalServerError)
	inreq.doneCh <- struct{}{}
}

func response(obj interface{}, inreq *webRequest) error {
	resp := obj.(*http.Response)
	respheaders := httphead.HeaderAsMap(resp.Header)
	for key, headers := range respheaders {
		inreq.w.Header().Add(key, headers)
	}
	_, err := io.Copy(inreq.w, resp.Body)
	return err
}
