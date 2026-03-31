package Net

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

type Post struct {
	curl               Curl
	InsecureSkipVerify bool
	ret                *response
	err                error
	Timeout            time.Duration
}

func (self Post) New() *Post {
	return &self
}

func (self *Post) Proxy(proxyUrl string) *Post {
	purl, err := url.Parse(proxyUrl)
	if err != nil {
		self.err = err
		return self
	}
	self.curl.request.Proxy(http.ProxyURL(purl))
	return self
}

// proxy by socks5 is dont by golang proxy module
func (self *Post) ProxySocks5(tcpudp, addr string, proxyauth *proxy.Auth) *Post {
	self.curl.request.ProxySocks5(tcpudp, addr, proxyauth)
	return self
}

// this proxy method is done by http.request itself
func (self *Post) SetTimeOut(Timeout time.Duration) *Post {
	self.Timeout = Timeout
	return self
}

func (self *Post) AllowInsecure() *Post {
	self.InsecureSkipVerify = true
	return self
}

func (self *Post) PostRpc(url string, postData interface{}, username, password string) *Post {
	req := self.curl.newRequest().request
	self.curl.SetHeaderJson()
	req.SetBasicAuth(username, password)
	if self.Timeout != 0 {
		req.SetTimeout(self.Timeout)
	}
	req.DisableKeepAlives(true)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	self.ret, self.err = req.post(url, postData)
	return self
}

func (self *Post) PostRaw(url string, postData interface{}) *Post {
	req := self.curl.newRequest().request
	self.curl.SetHeaderTextPlain()
	if self.Timeout != 0 {
		req.SetTimeout(self.Timeout)
	}
	req.DisableKeepAlives(true)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	self.ret, self.err = req.post(url, postData)
	return self
}

func (self *Post) PostFormData(url string, queries map[string]interface{}, postData map[string]string, headers map[string]string, cookies map[string]string) *Post {
	req := self.curl.newRequest().request
	self.curl.SetHeaderFormData()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	if self.Timeout != 0 {
		req.SetTimeout(self.Timeout)
	}
	req.DisableKeepAlives(true)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	url, self.err = buildUrl(url, queries)
	if self.err != nil {
		return self
	}
	self.ret, self.err = req.postFD(url, postData)
	return self
}

func (self *Post) PostFormDataAny(url string, queries map[string]interface{}, postData map[string]any, headers map[string]string, cookies map[string]string) *Post {
	req := self.curl.newRequest().request
	self.curl.SetHeaderFormData()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	if self.Timeout != 0 {
		req.SetTimeout(self.Timeout)
	}
	req.DisableKeepAlives(true)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	url, self.err = buildUrl(url, queries)
	if self.err != nil {
		return self
	}
	self.ret, self.err = req.postFD(url, postData)
	return self
}

func (self *Post) PostUrlXEncode(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) *Post {
	req := self.curl.newRequest().request
	self.curl.SetHeaderUrlEncode()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	if self.Timeout != 0 {
		req.SetTimeout(self.Timeout)
	}
	req.DisableKeepAlives(true)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	url, self.err = buildUrl(url, queries)
	if self.err != nil {
		return self
	}
	self.ret, self.err = req.post(url, postData)
	return self
}

func (self *Post) PostJson(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) *Post {
	req := self.curl.newRequest().request
	self.curl.SetHeaderJson()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	if self.Timeout != 0 {
		req.SetTimeout(self.Timeout)
	}
	req.DisableKeepAlives(true)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	url, self.err = buildUrl(url, queries)
	if self.err != nil {
		return self
	}
	self.ret, self.err = req.post(url, postData)
	return self
}

func (self *Post) RetCookie() (cookie map[string]interface{}, err error) {
	if self.err != nil {
		return nil, self.err
	}
	return self.curl.cookieHandler(self.ret.Cookies()), nil
}

func (self *Post) RetString() (string, error) {
	if self.err != nil {
		return "", self.err
	}
	return self.ret.bodystring()
}

func (self *Post) RetBytes() ([]byte, error) {
	if self.err != nil {
		return nil, self.err
	}
	return self.ret.bodybytes()
}

func (self *Post) RetJson(v any) error {
	if self.err != nil {
		return self.err
	}
	return self.ret.bodyjson(v)
}
