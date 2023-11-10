package Net

//
//// Debug model
//func Debug(v bool) *request {
//	r := NewRequest()
//	return r.Debug(v)
//}
//
//func Jar(v http.CookieJar) *request {
//	r := NewRequest()
//	return r.Jar(v)
//}
//
//func DisableKeepAlives(v bool) *request {
//	r := NewRequest()
//	return r.DisableKeepAlives(v)
//}
//
//func CheckRedirect(v func(req *http.Request, via []*http.Request) error) *request {
//	r := NewRequest()
//	return r.CheckRedirect(v)
//}
//
//func TLSClient(v *tls.Config) *request {
//	r := NewRequest()
//	return r.SetTLSClient(v)
//}
//
//func SetTLSClient(v *tls.Config) *request {
//	r := NewRequest()
//	return r.SetTLSClient(v)
//}
//
//func SetHeaders(headers map[string]string) *request {
//	r := NewRequest()
//	return r.SetHeaders(headers)
//}
//
//func SetCookies(cookies map[string]string) *request {
//	r := NewRequest()
//	return r.SetCookies(cookies)
//}
//
//func SetBasicAuth(username, password string) *request {
//	r := NewRequest()
//	return r.SetBasicAuth(username, password)
//}
//
//func Proxy(v func(*http.Request) (*url.URL, error)) *request {
//	r := NewRequest()
//	return r.Proxy(v)
//}
//
//func SetTimeout(d time.Duration) *request {
//	r := NewRequest()
//	return r.SetTimeout(d)
//}
//
//func Transport(v *http.Transport) *request {
//	r := NewRequest()
//	return r.Transport(v)
//}
//
//// Put is a put http request
//func Put(url string, data ...interface{}) (*response, error) {
//	r := NewRequest()
//	return r.Put(url, data...)
//}
//
//// Delete is a delete http request
//func Delete(url string, data ...interface{}) (*response, error) {
//	r := NewRequest()
//	return r.Delete(url, data...)
//}
//
//// Upload file
//func Upload(url, filename, fileinput string) (*response, error) {
//	r := NewRequest()
//	return r.Upload(url, filename, fileinput)
//}
