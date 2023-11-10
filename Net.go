package Net

import (
	"github.com/tobycroft/Calc"
	"net"
	"net/http"
	"net/url"
	"time"
)

var dialer = &net.Dialer{
	Timeout:   5 * time.Second,
	KeepAlive: 0 * time.Second,
	//DualStack: true,
}
var transport = &http.Transport{
	DialContext:  dialer.DialContext,
	MaxIdleConns: 100,
}

func (req *Curl) NewRequest() *Curl {
	req.request = new(request)
	req.request.SetTimeout(30)
	req.request.SetHeaders(map[string]string{})
	req.request.SetCookies(map[string]string{})
	req.request.Transport(transport)
	return req
}

func (req *Curl) Http_build_query(querymap map[string]interface{}) string {
	query := make(url.Values)
	for k, v := range querymap {
		query.Add(k, Calc.Any2String(v))
	}
	//fmt.Println(query.Encode())
	return query.Encode()
}

type Net struct {
	Curl
}
