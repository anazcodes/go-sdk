package api

type Response struct {
	status   int
	response []byte
	method   string
	url      string
	curl     string
	err      error
}

func (r *Response) URL() string            { return r.url }
func (r *Response) Method() string         { return r.method }
func (r *Response) Curl() string           { return r.curl }
func (r *Response) Status() int            { return r.status }
func (r *Response) Response() []byte       { return r.response }
func (r *Response) ResponseString() string { return string(r.response) }
func (r *Response) Error() error           { return r.err }

// IsStatus2XX returns true if the status code is between 200 and 299.
func (r *Response) IsStatus2XX() bool {
	return r.status >= 200 && r.status < 300
}

func NewResponse(status int, response []byte, url, method, curl string, err error) *Response {
	return &Response{
		url:      url,
		method:   method,
		curl:     curl,
		status:   status,
		response: response,
		err:      err,
	}
}

func responseError(err error, url, method, curl string) *Response {
	return &Response{
		url:    url,
		method: method,
		curl:   curl,
		err:    err,
	}
}
