package Net

type Net struct {
	Curl
}

type Curl struct {
	request *request
	Header
}

type Post struct {
	Curl
}

type Get struct {
	Curl
}

type Header struct {
}
