package Net

func (self *Post) Rpc(url string, postData interface{}, username, password string) (string, error) {
	req := self.Curl.NewRequest().request
	self.Curl.SetHeaderJson()
	req.SetBasicAuth(username, password)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	ret, err := req.Post(url, postData)
	body, err := ret.Content()
	if err != nil {
		return "", err
	} else {
		return body, err
	}
}

func (self *Post) PostRaw(url string, postData interface{}) (string, error) {
	req := self.Curl.NewRequest().request
	self.Curl.SetHeaderTextPlain()
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	ret, err := req.Post(url, postData)
	body, err := ret.Content()
	if err != nil {
		return "", err
	} else {
		return body, err
	}
}

func (self *Post) Post(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) (string, error) {
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
	ret, err := req.Post(url+q, postData)
	if err != nil {
		return "", err
	}
	return ret.Content()
}

func (self *Post) PostJson(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) (string, error) {
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
	ret, err := req.Post(url+q, postData)
	if err != nil {
		return "", err
	}
	return ret.Content()
}

func (self *Post) PostCookie(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) (body string, cookie map[string]interface{}, err error) {
	req := self.Curl.NewRequest().request
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	ret, err := req.Post(url+"?"+self.Http_build_query(queries), postData)
	if err != nil {
		return
	}
	body, err = ret.Content()
	if err != nil {
		return
	}
	cookie_arr := self.CookieHandler(ret.Cookies())
	//fmt.Println(cookie_arr)
	if err != nil {
		return "", cookie_arr, err
	} else {
		return body, cookie_arr, err
	}
}
