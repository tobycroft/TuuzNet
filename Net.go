package Net

import (
	"net"
	"net/http"
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

type Net struct {
	Curl
	WsClient
}
