package client

type RequestContextFormEncoded[T1 any, T2 any] struct {
	client *RestClient
	acceptType ContentType
}

func NewRequestContextFormEncoded[T1 any, T2 any](client *RestClient, acceptType ContentType) RequestContextFormEncoded[T1, T2] {
	return RequestContextFormEncoded[T1, T2]{
		client: client,
		acceptType: acceptType,
	}
}

func (c RequestContextFormEncoded[T1, T2]) Execute(request *Request[T1, T2]) (response *Response[T2], err error) {
	request.EnsureContentType(ContentTypeFormUrlEncoded)
	request.EnsureAcceptType(c.acceptType)
	return executeRequest(c.client, request)
}

func (c RequestContextFormEncoded[T1, T2]) executeReturnStatus(request *Request[T1, T2]) (statusCode int, status string, err error) {
	resp, err := executeRequest(c.client, request)
	if resp != nil {
		statusCode = resp.StatusCode
		status = resp.Status
	}

	return statusCode, status, err
}

func (c RequestContextFormEncoded[T1, T2]) Get(path string, requestData *T1, responseData *T2) (statusCode int, status string, err error) {
	return c.executeReturnStatus(&Request[T1, T2]{
		Method:       "GET",
		Path:         path,
		RequestData:  requestData,
		ResponseData: responseData,
	})
}

func (c RequestContextFormEncoded[T1, T2]) Put(path string, requestData *T1, responseData *T2) (statusCode int, status string, err error) {
	return c.executeReturnStatus(&Request[T1, T2]{
		Method:       "PUT",
		Path:         path,
		RequestData:  requestData,
		ResponseData: responseData,
	})
}

func (c RequestContextFormEncoded[T1, T2]) Post(path string, requestData *T1, responseData *T2) (statusCode int, status string, err error) {
	return c.executeReturnStatus(&Request[T1, T2]{
		Method:       "POST",
		Path:         path,
		RequestData:  requestData,
		ResponseData: responseData,
	})
}

func (c RequestContextFormEncoded[T1, T2]) Patch(path string, requestData *T1, responseData *T2) (statusCode int, status string, err error) {
	return c.executeReturnStatus(&Request[T1, T2]{
		Method:       "PATCH",
		Path:         path,
		RequestData:  requestData,
		ResponseData: responseData,
	})
}

func (c RequestContextFormEncoded[T1, T2]) Delete(path string, requestData *T1, responseData *T2) (statusCode int, status string, err error) {
	return c.executeReturnStatus(&Request[T1, T2]{
		Method:       "DELETE",
		Path:         path,
		RequestData:  requestData,
		ResponseData: responseData,
	})
}
