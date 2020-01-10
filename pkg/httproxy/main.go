package httproxy

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/elazarl/goproxy"
)

const (
	defaultHost = "127.0.0.1"
	defaultPort = 8181

	defaultWait = time.Second
)

// HttProxy is ...
type HttProxy struct {
	srv   *http.Server
	proxy *goproxy.ProxyHttpServer

	certContents []byte
	keyContents  []byte
	cert         tls.Certificate

	listenHost string
	listenPort int

	ch chan []string

	e    error
	wait time.Duration
}

// NewHttProxy function...
func NewHttProxy() (HttProxy, error) {
	var e error
	p := HttProxy{}

	p.listenHost = defaultHost
	p.listenPort = defaultPort
	p.wait = defaultWait
	p.ch = make(chan []string, 10)

	e = p.setCA(caCert, caKey)
	if e != nil {
		log.Fatal(e)
	}

	p.proxy = goproxy.NewProxyHttpServer()
	p.proxy.Verbose = false
	p.proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	p.proxy.OnResponse().DoFunc(p.responseAction)

	return p, nil
}

// StartBackground function...
func (h *HttProxy) StartBackground(addr string) error {
	h.srv = &http.Server{
		Addr:    addr,
		Handler: h.proxy,
	}

	go func() {
		h.e = nil
		h.e = h.srv.ListenAndServe()
	}()

	time.Sleep(h.wait)
	return h.e
}

// StopBackground function...
func (h *HttProxy) StopBackground() error {
	h.ch <- []string{"stop"}

	e := h.srv.Shutdown(nil)
	if e != nil {
		return e
	}

	time.Sleep(h.wait)
	return h.e
}

// Read function...
func (h *HttProxy) Read() (string, string, string, string, error) {
	v := <-h.ch
	if len(v) != 4 {
		return "", "", "", "", errors.New("HttProxy closed")
	}

	method := v[0]
	url := v[1]
	request := v[2]
	response := v[3]
	return method, url, request, response, nil
}

func (h *HttProxy) responseAction(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	// No action if Nil response
	if r == nil {
		return r
	}

	// Extract URL and Method
	urlString := r.Request.URL.Scheme + "://" + r.Request.URL.Host + r.Request.URL.Path
	methodString := r.Request.Method

	// Extract request and base64 encode
	requestData, e1 := httputil.DumpRequest(r.Request, true)
	if e1 != nil {
		return r
	}
	requestString := base64.StdEncoding.EncodeToString(requestData)

	// Extract response and base64 encode
	responseData, e2 := httputil.DumpResponse(r, true)
	if e2 != nil {
		return r
	}
	responseString := base64.StdEncoding.EncodeToString(responseData)

	// Send request/response to output channel
	h.ch <- []string{methodString, urlString, requestString, responseString}

	return r
}
