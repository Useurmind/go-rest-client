package client

type Request[T any, R any] struct {
	Method            string
	Path              string
	ContentType       ContentType
	AcceptType        ContentType
	RequestData       *T
	ResponseData      *R
	AdditionalHeaders map[string]string
}

func (r *Request[T, R]) EnsureContentType(contentType ContentType) {
	if r.RequestData != nil {
		r.ContentType = contentType
	}
}

func (r *Request[T, R]) EnsureAcceptType(accpetType ContentType) {
	if r.ResponseData != nil {
		r.AcceptType = accpetType
	}
}