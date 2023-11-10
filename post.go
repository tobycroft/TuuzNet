package Net

func (self *Post) Rpc(url string, postData interface{}, username, password string) (string, error) {
	req := self.request
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
	req := self.request
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
	req := self.request
	self.Curl.SetHeaderFormData()
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	q := ""
	req.Transport(transport)
	if queries != nil {
		q = "?" + Http_build_query(queries)
	}
	ret, err := req.Post(url+q, postData)
	if err != nil {
		return "", err
	}
	return ret.Content()
}

func (self *Post) PostCookie(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookies map[string]string) (string, map[string]interface{}, error) {
	req := self.request
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	req.Transport(transport)
	ret, err := req.Post(url+"?"+Http_build_query(queries), postData)
	body, err := ret.Content()

	cookie_arr := CookieHandler(ret.Cookies())
	//fmt.Println(cookie_arr)
	if err != nil {
		return "", cookie_arr, err
	} else {
		return body, cookie_arr, err
	}
}

//func (self *Post) PostCookieAuto(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, ident string) (string, error) {
//	req := self.request
//	cookies, err := CookieSelector(ident)
//	cook := Array.MapAny2MapString(cookies)
//
//	req.SetHeaders(headers)
//	req.SetCookies(cook)
//	req.SetTimeout(5 )
//	req.DisableKeepAlives(true)
//	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
//	req.Transport(transport)
//	ret, err := req.Post(url+"?"+Http_build_query(queries), postData)
//	body, err := ret.Content()
//
//	cookie_arr := CookieHandler(ret.Cookies())
//	CookieUpdater(cookie_arr, ident)
//	if err != nil {
//		return "", err
//	} else {
//		return body, err
//	}
//}

//func (self *Post) PostCookieManual(url string, queries map[string]interface{}, postData map[string]interface{}, headers map[string]string, cookie map[string]interface{}, ident string) (string, error) {
//	req := self.request
//	CookieUpdater(cookie, ident)
//	cookies, err := CookieSelector(ident)
//	cook := Array.MapAny2MapString(cookies)
//
//	req.SetHeaders(headers)
//	req.SetCookies(cook)
//	req.SetTimeout(5 )
//	req.DisableKeepAlives(true)
//	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
//	req.Transport(transport)
//	ret, err := req.Post(url+"?"+Http_build_query(queries), postData)
//	body, err := ret.Content()
//
//	cookie_arr := CookieHandler(ret.Cookies())
//	CookieUpdater(cookie_arr, ident)
//	if err != nil {
//		return "", err
//	} else {
//		return body, err
//	}
//}

//func (self *Post) GetCookieAuto(url string, queries map[string]interface{}, headers map[string]string, ident string) (string, error) {
//	req := self.request
//	cookies, err := CookieSelector(ident)
//	cook := Array.MapAny2MapString(cookies)
//
//	req.SetHeaders(headers)
//	req.SetCookies(cook)
//	req.SetTimeout(5 )
//	req.DisableKeepAlives(true)
//	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
//	req.Transport(transport)
//	ret, err := req.Get(url, queries)
//	if err != nil {
//		fmt.Println(err)
//		return "", err
//	}
//	body, err := ret.Content()
//	if err != nil {
//		return "", err
//	}
//	cookie_arr := CookieHandler(ret.Cookies())
//	CookieUpdater(cookie_arr, ident)
//	if err != nil {
//		return "", err
//	} else {
//		return body, err
//	}
//}

//func (self *Post) GetCookieManual(url string, queries map[string]interface{}, headers map[string]string, cookie map[string]interface{}, ident string) (string, error) {
//	req := self.request
//	CookieUpdater(cookie, ident)
//	cookies, err := CookieSelector(ident)
//	cook := Array.MapAny2MapString(cookies)
//
//	req.SetHeaders(headers)
//	req.SetCookies(cook)
//	req.SetTimeout(5 )
//	req.DisableKeepAlives(true)
//	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
//	req.Transport(transport)
//	ret, err := req.Get(url, queries)
//	if err != nil {
//		fmt.Println(err)
//		return "", err
//	}
//	body, err := ret.Content()
//	cookie_arr := CookieHandler(ret.Cookies())
//	CookieUpdater(cookie_arr, ident)
//	if err != nil {
//		return "", err
//	} else {
//		return body, err
//	}
//}
