package api

type Response struct {
	Status   int
	Response []byte
	Method   string
	URL      string
	Curl     string
	Err      error
}

func NewResponse(status int, response []byte, url, method, curl string, err error) Response {
	return Response{
		URL:      url,
		Method:   method,
		Curl:     curl,
		Status:   status,
		Response: response,
		Err:      err,
	}
}

func (r *Response) GetURL() string            { return r.URL }
func (r *Response) GetMethod() string         { return r.Method }
func (r *Response) GetCurl() string           { return r.Curl }
func (r *Response) GetStatus() int            { return r.Status }
func (r *Response) GetResponse() []byte       { return r.Response }
func (r *Response) GetResponseString() string { return string(r.Response) }
func (r *Response) GetError() error           { return r.Err }

func responseError(err error, url, method, curl string) Response {
	return Response{
		URL:    url,
		Method: method,
		Curl:   curl,
		Err:    err,
	}
}

// IsStatus2XX returns true if the status code is between 200 and 299.
func (r Response) IsStatus2XX() bool {
	return r.Status >= 200 && r.Status < 300
}
