package Net

func (h *Curl) SetHeaderJson() {
	h.request.SetHeaders(map[string]string{"Content-Type": "application/json"})
}

func (h *Curl) SetHeaderUrlEncode() {
	h.request.SetHeaders(map[string]string{"Content-Type": "application/x-www-form-urlencoded"})
}

func (h *Curl) SetHeaderFormData() {
	h.request.SetHeaders(map[string]string{"Content-Type": "multipart/form-data"})
}
