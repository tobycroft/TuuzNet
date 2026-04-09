package TuuzNet

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

type PostBuilder struct {
	url        string
	query      map[string]any
	postData   map[string]any
	header     map[string]string
	cookies    map[string]string
	setTimeOut time.Duration
	debug      bool
	Post       *Post
}

func (self PostBuilder) New() *PostBuilder {
	self.Post = new(Post)
	self.setTimeOut = 5
	self.query = map[string]any{}
	self.header = map[string]string{}
	self.cookies = map[string]string{}
	self.postData = map[string]any{}
	return &self
}

func (self *PostBuilder) Debug() {
	self.debug = true
}

func (self *PostBuilder) Proxy(proxyUrl string) *PostBuilder {
	purl, err := url.Parse(proxyUrl)
	if err != nil {
		self.Post.err = err
		return self
	}
	self.Post.curl.request.Proxy(http.ProxyURL(purl))
	return self
}

// proxy by socks5 is dont by golang proxy module
func (self *PostBuilder) ProxySocks5(tcpudp, addr string, proxyauth *proxy.Auth) *PostBuilder {
	self.Post.curl.request.ProxySocks5(tcpudp, addr, proxyauth)
	return self
}

func (self *PostBuilder) SetUrl(url string) *PostBuilder {
	self.url = url
	return self
}

func (self *PostBuilder) SetQuery(query map[string]any) *PostBuilder {
	for s, a := range query {
		self.query[s] = a
	}
	return self
}

func (self *PostBuilder) SetPostData(postData map[string]any) *PostBuilder {
	for s, a := range postData {
		self.postData[s] = a
	}
	return self
}

func (self *PostBuilder) SetHeader(header map[string]string) *PostBuilder {
	for s, s2 := range header {
		self.header[s] = s2
	}
	return self
}

func (self *PostBuilder) SetCookies(cookies map[string]string) *PostBuilder {
	for s, s2 := range cookies {
		self.cookies[s] = s2
	}
	return self
}

func (self *PostBuilder) SetTimeOut(timeOut time.Duration) *PostBuilder {
	self.setTimeOut = timeOut
	return self
}

func (self *PostBuilder) AllowInsecure() *PostBuilder {
	self.Post.InsecureSkipVerify = true
	return self
}

func (self *PostBuilder) DisableKeepAlives() *PostBuilder {
	self.Post.DisableKeepAlives = true
	return self
}

func (self *PostBuilder) SendRPC(username, password string) *Ret {
	req := self.Post.curl.newRequest().request
	req.Debug(self.debug)
	self.Post.curl.SetHeaderJson()
	req.SetBasicAuth(username, password)
	if self.setTimeOut != 0 {
		req.SetTimeout(self.setTimeOut)
	}
	req.DisableKeepAlives(self.Post.DisableKeepAlives)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.Post.InsecureSkipVerify})
	self.url, self.Post.err = buildUrl(self.url, self.query)
	self.Post.ret, self.Post.err = req.post(self.url, self.postData)
	return &Ret{&self.Post.curl, self.Post.ret, self.Post.err}
}

func (self *PostBuilder) SendRAW(postData interface{}) *Ret {
	req := self.Post.curl.newRequest().request
	req.Debug(self.debug)
	self.Post.curl.SetHeaderTextPlain()
	if self.setTimeOut != 0 {
		req.SetTimeout(self.setTimeOut)
	}
	req.DisableKeepAlives(self.Post.DisableKeepAlives)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.Post.InsecureSkipVerify})
	self.Post.ret, self.Post.err = req.post(self.url, postData)
	return &Ret{&self.Post.curl, self.Post.ret, self.Post.err}
}

func (self *PostBuilder) SendFormData() *Ret {
	req := self.Post.curl.newRequest().request
	req.Debug(self.debug)
	self.Post.curl.SetHeaderFormData()
	req.SetHeaders(self.header)
	req.SetCookies(self.cookies)
	if self.setTimeOut != 0 {
		req.SetTimeout(self.setTimeOut)
	}
	req.DisableKeepAlives(self.Post.DisableKeepAlives)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.Post.InsecureSkipVerify})
	self.url, self.Post.err = buildUrl(self.url, self.query)
	if self.Post.err != nil {
		return &Ret{&self.Post.curl, self.Post.ret, self.Post.err}
	}
	self.Post.ret, self.Post.err = req.postFD(self.url, self.postData)
	return &Ret{&self.Post.curl, self.Post.ret, self.Post.err}
}

func (self *PostBuilder) SendFormDataAny() *Ret {
	req := self.Post.curl.newRequest().request
	req.Debug(self.debug)
	self.Post.curl.SetHeaderFormData()
	req.SetHeaders(self.header)
	req.SetCookies(self.cookies)
	if self.setTimeOut != 0 {
		req.SetTimeout(self.setTimeOut)
	}
	req.DisableKeepAlives(self.Post.DisableKeepAlives)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.Post.InsecureSkipVerify})
	self.url, self.Post.err = buildUrl(self.url, self.query)
	if self.Post.err != nil {
		return &Ret{&self.Post.curl, self.Post.ret, self.Post.err}
	}
	self.Post.ret, self.Post.err = req.postFD(self.url, self.postData)
	return &Ret{&self.Post.curl, self.Post.ret, self.Post.err}
}

func (self *PostBuilder) PostFormDataAny() *Ret {
	return self.SendFormDataAny()
}

func (self *PostBuilder) SendUrlXEncode() *Ret {
	req := self.Post.curl.newRequest().request
	req.Debug(self.debug)
	self.Post.curl.SetHeaderUrlEncode()
	req.SetHeaders(self.header)
	req.SetCookies(self.cookies)
	if self.setTimeOut != 0 {
		req.SetTimeout(self.setTimeOut)
	}
	req.DisableKeepAlives(self.Post.DisableKeepAlives)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.Post.InsecureSkipVerify})
	self.url, self.Post.err = buildUrl(self.url, self.query)
	if self.Post.err != nil {
		return &Ret{&self.Post.curl, self.Post.ret, self.Post.err}
	}
	self.Post.ret, self.Post.err = req.post(self.url, self.postData)
	return &Ret{&self.Post.curl, self.Post.ret, self.Post.err}
}

func (self *PostBuilder) PostUrlXEncode() *Ret {
	return self.SendUrlXEncode()
}

func (self *PostBuilder) SendJson() *Ret {
	req := self.Post.curl.newRequest().request
	req.Debug(self.debug)
	self.Post.curl.SetHeaderJson()
	req.SetHeaders(self.header)
	req.SetCookies(self.cookies)
	if self.setTimeOut != 0 {
		req.SetTimeout(self.setTimeOut)
	}
	req.DisableKeepAlives(self.Post.DisableKeepAlives)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.Post.InsecureSkipVerify})
	self.url, self.Post.err = buildUrl(self.url, self.query)
	if self.Post.err != nil {
		return &Ret{&self.Post.curl, self.Post.ret, self.Post.err}
	}
	self.Post.ret, self.Post.err = req.post(self.url, self.postData)
	return &Ret{&self.Post.curl, self.Post.ret, self.Post.err}
}

func (self *PostBuilder) PostJson() *Ret {
	return self.SendJson()
}
