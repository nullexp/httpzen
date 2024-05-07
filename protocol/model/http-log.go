package model

type HttpLog struct {
	Request         *Request
	Response        *Response
	RequestHeaders  []*Header
	ResponseHeaders []*Header
	RequestBody     *Body
	ResponseBody    *Body
}
