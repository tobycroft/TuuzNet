package Net

type Net struct {
	Curl
}

type Curl struct {
	request *request
}

type Post struct {
	Curl
}

type Get struct {
	Curl
}
