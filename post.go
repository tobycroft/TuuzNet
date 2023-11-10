package Net

func (self *Post) Rpc(url string, postData interface{}, username, password string) (string, error) {
	req := self.Curl.NewRequest().request
	header := map[string]string{"Content-type": "application/json"}
	req.SetHeaders(header)
	req.SetBasicAuth(username, password)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	req.Transport(transport)
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
	header := map[string]string{"Content-type": "application/json"}
	req.SetHeaders(header)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	req.Transport(transport)
	ret, err := req.Post(url, postData)
	body, err := ret.Content()
	if err != nil {
		return "", err
	} else {
		return body, err
	}
}

func (self *Post) Post(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) (string, error) {
	// 链式操作

	req := self.Curl.NewRequest().request
	self.Curl.SetHeaderFormData()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	q := ""
	req.Transport(transport)
	if queries != nil {
		q = "?" + self.Http_build_query(queries)
	}
	ret, err := req.Post(url+q, postData)
	if err != nil {
		return "", err
	}
	return ret.Content()
}

func (self *Post) PostCookie(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) (string, map[string]interface{}, error) {
	req := self.Curl.NewRequest().request
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	req.Transport(transport)
	ret, err := req.Post(url+"?"+self.Http_build_query(queries), postData)
	body, err := ret.Content()

	cookie_arr := CookieHandler(ret.Cookies())
	//fmt.Println(cookie_arr)
	if err != nil {
		return "", cookie_arr, err
	} else {
		return body, cookie_arr, err
	}
}
