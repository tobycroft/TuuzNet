package Net

import (
	"crypto/tls"
	"time"
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

func (self *GetBuilder) SetUrl(url string) *GetBuilder {
	self.url = url
	return self
}

func (self *GetBuilder) SetQuery(query map[string]any) *GetBuilder {
	self.query = query
	return self
}

func (self *GetBuilder) SetHeader(header map[string]string) *GetBuilder {
	self.header = header
	return self
}

func (self *GetBuilder) SetCookies(cookies map[string]string) *GetBuilder {
	self.cookies = cookies
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

func (self *GetBuilder) SendGet() *Get {
	req := self.Get.curl.newRequest().request
	req.SetHeaders(self.header)
	req.SetCookies(self.cookies)
	req.SetTimeout(self.setTimeOut)
	req.DisableKeepAlives(self.Get.DisableKeepAlives)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.Get.InsecureSkipVerify})
	self.Get.ret, self.Get.err = req.Get(self.url, self.query)
	return self.Get
}
