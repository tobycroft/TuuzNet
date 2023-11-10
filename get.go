package Net

import (
	"fmt"
)

func (self *Get) Get(url string, queries map[string]interface{}, headers map[string]string, cookies map[string]string) (string, error) {
	req := self.request
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	req.Transport(transport)
	ret, err := req.Get(url, queries)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	body, err := ret.Content()

	if err != nil {
		return "", err
	} else {
		return body, err
	}
}

func (self *Get) GetCookie(url string, queries map[string]interface{}, headers map[string]string, cookies map[string]string) (string, map[string]interface{}, error) {
	req := self.request
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	req.Transport(transport)
	ret, err := req.Get(url, queries)
	body, err := ret.Content()
	cookie_arr := CookieHandler(ret.Cookies())
	//fmt.Println(cookie_arr)
	if err != nil {
		return "", cookie_arr, err
	} else {
		return body, cookie_arr, err
	}
}
