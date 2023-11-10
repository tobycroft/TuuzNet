package Net

type Post struct {
	Curl
	ret *response
	err error
}

func (self Post) Rpc(url string, postData interface{}, username, password string) Post {
	req := self.Curl.NewRequest().request
	self.Curl.SetHeaderJson()
	req.SetBasicAuth(username, password)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	self.ret, self.err = req.post(url, postData)
	return self
}

func (self Post) Raw(url string, postData interface{}) Post {
	req := self.Curl.NewRequest().request
	self.Curl.SetHeaderTextPlain()
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	self.ret, self.err = req.post(url, postData)
	return self
}

func (self Post) FormData(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) Post {
	req := self.Curl.NewRequest().request
	self.Curl.SetHeaderFormData()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	q := ""
	if queries != nil {
		q = "?" + self.Http_build_query(queries)
	}
	self.ret, self.err = req.post(url+q, postData)
	return self
}

func (self Post) UrlXEncode(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) Post {
	req := self.Curl.NewRequest().request
	self.Curl.SetHeaderUrlEncode()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	q := ""
	if queries != nil {
		q = "?" + self.Http_build_query(queries)
	}
	self.ret, self.err = req.post(url+q, postData)
	return self
}

func (self Post) Json(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) Post {
	req := self.Curl.NewRequest().request
	self.Curl.SetHeaderJson()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	q := ""
	if queries != nil {
		q = "?" + self.Http_build_query(queries)
	}
	self.ret, self.err = req.post(url+q, postData)
	return self
}

func (self Post) GetCookie() (cookie map[string]interface{}) {
	return self.CookieHandler(self.ret.Cookies())
}

func (self Post) GetString() (string, error) {
	return self.ret.bodystring()
}

func (self Post) Getbytes() ([]byte, error) {
	return self.ret.bodybytes()
}

func (self Post) GetJson(v any) error {
	return self.ret.bodyjson(v)
}
