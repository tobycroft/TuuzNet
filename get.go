package Net

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

type Get struct {
	curl               Curl
	InsecureSkipVerify bool
	ret                *response
	err                error
	DisableKeepAlives  bool
}

func (self Get) New(debug_leftAsFalse ...bool) *Get {
	if debug_leftAsFalse != nil {
		self.curl.request.debug = true
	}
	return &self
}

func (self *Get) Get(url string, queries map[string]any, headers map[string]string, cookies map[string]string) *Ret {
	req := self.curl.newRequest().request
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(self.DisableKeepAlives)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	self.ret, self.err = req.Get(url, queries)
	return &Ret{&self.curl, self.ret, self.err}
}

// proxy by socks5 is dont by golang proxy module
func (self *Get) ProxySocks5(tcpudp, addr string, proxyauth *proxy.Auth) *Get {
	self.curl.request.ProxySocks5(tcpudp, addr, proxyauth)
	return self
}

// this proxy method is done by http.request itself
func (self *Get) ProxyHttp(proxyUrl string) *Get {
	purl, err := url.Parse(proxyUrl)
	if err != nil {
		self.err = err
		return self
	}
	self.curl.request.Proxy(http.ProxyURL(purl))
	return self
}
