package Net

type Ret struct {
	curl *Curl
	ret  *response
	err  error
}

func (self *Ret) RetCookie() (cookie map[string]interface{}, err error) {
	if self.err != nil {
		return nil, self.err
	}
	return self.curl.cookieHandler(self.ret.Cookies()), nil
}

func (self *Ret) RetString() (string, error) {
	if self.err != nil {
		return "", self.err
	}
	return self.ret.bodystring()
}

func (self *Ret) RetBytes() ([]byte, error) {
	if self.err != nil {
		return nil, self.err
	}
	return self.ret.bodybytes()
}

func (self *Ret) RetJson(v any) error {
	if self.err != nil {
		return self.err
	}
	return self.ret.bodyjson(v)
}
