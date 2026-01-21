package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func executeRequest[T1 any, T2 any](c *RestClient, request *Request[T1, T2]) (*Response[T2], error) {
	var err error

	url, err := url.JoinPath(c.baseUrl, request.Path)
	if err != nil {
		return nil, err
	}

	reader, err := getBodyReader(request)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest(request.Method, url, reader)
	if err != nil {
		return nil, err
	}
	if request.ContentType != "" {
		httpReq.Header.Add("Content-Type", string(request.ContentType))
	}
	if request.AcceptType != "" {
		httpReq.Header.Add("Accept", string(request.AcceptType))
	}
	if c.bearerToken != "" {
		httpReq.Header.Add("Authorization", "Bearer "+c.bearerToken)
	}
	if c.basicAuth != nil {
		httpReq.SetBasicAuth(c.basicAuth.username, c.basicAuth.password)
	}
	for k, v := range request.AdditionalHeaders {
		httpReq.Header.Add(k, v)
	}

	shortRequestString := fmt.Sprintf("%s %s", request.Method, url)
	c.logger.Info(fmt.Sprintf("execute HTTP request: %s", shortRequestString))
	resp := &Response[T2]{}
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	resp.Status = httpResp.Status
	resp.StatusCode = httpResp.StatusCode

	if httpResp.StatusCode < 200 || httpResp.StatusCode > 399 {
		return resp, fmt.Errorf("request %s returned %d - %s", shortRequestString, resp.StatusCode, resp.Status)
	}

	responseData, err := decodeResponseData(c, request, httpResp.Body)
	if err != nil {
		return resp, err
	}
	resp.Data = responseData

	return resp, nil
}

func getBodyReader[T1 any, T2 any](c *RestClient, request *Request[T1, T2]) (io.Reader, error) {
	switch request.ContentType {
	case ContentTypeJson:
		bodyBytes, err := json.Marshal(request.RequestData)
		if err != nil {
			return nil, err
		}

		c.logger.V(4).Info("json body is:\n" + string(bodyBytes))

		return bytes.NewBuffer(bodyBytes), nil
	case ContentTypeFormUrlEncoded:
		values, err := EncodeUrlValues(request.RequestData)
		if err != nil {
			return nil, err
		}
		c.logger.V(4).Info("form encoded data is: " + values.Encode())
		return strings.NewReader(values.Encode()), nil
	}

	return nil, nil
}

func decodeResponseData[T1 any, T2 any](c *RestClient, request *Request[T1, T2], reader io.Reader) (*T2, error) {
	if request.ResponseData != nil {
		switch request.AcceptType {
		case ContentTypeJson:
			b, err := io.ReadAll(reader)
			if err != nil {
				return nil, err
			}

			c.logger.V(4).Info("response body:\n" + string(b))

			err = json.Unmarshal(b, request.ResponseData)
			if err != nil {
				return nil, err
			}

			return request.ResponseData, nil
		default:
			c.logger.V(4).Info("accept type unkown", "acceptType", request.AcceptType)
		}
	} else {
		c.logger.V(4).Info("response data object not set, skip body deserialization")
	}

	return nil, nil
}
