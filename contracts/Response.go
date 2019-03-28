package contracts

import "net/http"

type Response struct {
	http.ResponseWriter
	status       int
	wroteHeader  bool
	ErrorMessage string
}

func (*Response) Header() http.Header {
	return http.Header{}
}
func (res *Response) Write(data []byte) (int, error) {
	if !res.wroteHeader {
		res.WriteHeader(http.StatusOK)
	}
	res.ErrorMessage = string(data)
	return 0, nil
}
func (res *Response) WriteHeader(statusCode int) {
	if res.wroteHeader {
		return
	}
	res.status = statusCode
	res.wroteHeader = true
}
func (res *Response) Status() int {
	return res.status
}
