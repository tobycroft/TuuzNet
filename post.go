package Net

type Post struct {
	curl Curl
	ret  *response
	err  error
}

func (self Post) PostRpc(url string, postData interface{}, username, password string) Post {
	req := self.curl.NewRequest().request
	self.curl.SetHeaderJson()
	req.SetBasicAuth(username, password)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	self.ret, self.err = req.post(url, postData)
	return self
}

func (self Post) PostRaw(url string, postData interface{}) Post {
	req := self.curl.NewRequest().request
	self.curl.SetHeaderTextPlain()
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	self.ret, self.err = req.post(url, postData)
	return self
}

func (self Post) PostFormData(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) Post {
	req := self.curl.NewRequest().request
	self.curl.SetHeaderFormData()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	q := ""
	if queries != nil {
		q = "?" + self.curl.Http_build_query(queries)
	}
	self.ret, self.err = req.post(url+q, postData)
	return self
}

func (self Post) PostUrlXEncode(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) Post {
	req := self.curl.NewRequest().request
	self.curl.SetHeaderUrlEncode()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	q := ""
	if queries != nil {
		q = "?" + self.curl.Http_build_query(queries)
	}
	self.ret, self.err = req.post(url+q, postData)
	return self
}

func (self Post) PostJson(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) Post {
	req := self.curl.NewRequest().request
	self.curl.SetHeaderJson()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	q := ""
	if queries != nil {
		q = "?" + self.curl.Http_build_query(queries)
	}
	self.ret, self.err = req.post(url+q, postData)
	return self
}

func (self Post) RetCookie() (cookie map[string]interface{}) {
	return self.curl.cookieHandler(self.ret.Cookies())
}

func (self Post) RetString() (string, error) {
	return self.ret.bodystring()
}

func (self Post) RetBytes() ([]byte, error) {
	return self.ret.bodybytes()
}

func (self Post) RetJson(v any) error {
	return self.ret.bodyjson(v)
}
