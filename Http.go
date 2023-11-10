package Net

type Curl struct {
	request *request
}

type Post struct {
	Curl
}

type Get struct {
	Curl
}
