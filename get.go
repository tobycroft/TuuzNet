package Net

import (
	"fmt"
)

type Get struct {
	Curl
	ret *response
	err error
}

func (self Get) Get(url string, queries map[string]interface{}, headers map[string]string, cookies map[string]string) (string, error) {
	req := self.Curl.NewRequest().request
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	ret, err := req.Get(url, queries)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	body, err := ret.bodystring()

	if err != nil {
		return "", err
	} else {
		return body, err
	}
}

func (self Get) GetCookie(url string, queries map[string]interface{}, headers map[string]string, cookies map[string]string) (string, map[string]interface{}, error) {
	req := self.Curl.NewRequest().request
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	//req.SetTLSClient(&tls.Config{InsecureSkipVerify: true})
	ret, err := req.Get(url, queries)
	body, err := ret.bodystring()
	cookie_arr := self.CookieHandler(ret.Cookies())
	//fmt.Println(cookie_arr)
	if err != nil {
		return "", cookie_arr, err
	} else {
		return body, cookie_arr, err
	}
}
