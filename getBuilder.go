package TuuzNet

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

type GetBuilder struct {
	url        string
	query      map[string]any
	header     map[string]string
	cookies    map[string]string
	setTimeOut time.Duration
	Get        *Get
}

func (self GetBuilder) New() *GetBuilder {
	self.Get = new(Get)
	self.setTimeOut = 5
	return &self
}

// proxy by socks5 is dont by golang proxy module
func (self *GetBuilder) ProxySocks5(tcpudp, addr string, proxyauth *proxy.Auth) *GetBuilder {
	self.Get.curl.request.ProxySocks5(tcpudp, addr, proxyauth)
	return self
}

// this proxy method is done by http.request itself
func (self *GetBuilder) ProxyHttp(proxyUrl string) *GetBuilder {
	purl, err := url.Parse(proxyUrl)
	if err != nil {
		self.Get.err = err
		return self
	}
	self.Get.curl.request.Proxy(http.ProxyURL(purl))
	return self
}

func (self *GetBuilder) SetUrl(url string) *GetBuilder {
	self.url = url
	return self
}

func (self *GetBuilder) SetQuery(query map[string]any) *GetBuilder {
	for s, a := range query {
		self.query[s] = a
	}
	return self
}

func (self *GetBuilder) SetHeader(header map[string]string) *GetBuilder {
	for s, s2 := range header {
		self.header[s] = s2
	}
	return self
}

func (self *GetBuilder) SetCookies(cookies map[string]string) *GetBuilder {
	for s, s2 := range cookies {
		self.cookies[s] = s2
	}
	return self
}

func (self *GetBuilder) SetTimeOut(timeOut time.Duration) *GetBuilder {
	self.setTimeOut = timeOut
	return self
}

func (self *GetBuilder) AllowInsecure() *GetBuilder {
	self.Get.InsecureSkipVerify = true
	return self
}

func (self *GetBuilder) DisableKeepAlives() *GetBuilder {
	self.Get.DisableKeepAlives = true
	return self
}

func (self *GetBuilder) SendGet() *Ret {
	req := self.Get.curl.newRequest().request
	req.SetHeaders(self.header)
	req.SetCookies(self.cookies)
	req.SetTimeout(self.setTimeOut)
	req.DisableKeepAlives(self.Get.DisableKeepAlives)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.Get.InsecureSkipVerify})
	self.Get.ret, self.Get.err = req.Get(self.url, self.query)
	return &Ret{&self.Get.curl, self.Get.ret, self.Get.err}
}
