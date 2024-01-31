package apisix

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
	"time"
)

const (
	RequestTimeout = 5000
)

// QueryParamsOption Get请求查询参数
type QueryParamsOption func(map[string]string)

// WithPageNumber 页数，默认展示第一页。
func WithPageNumber(pageNumber int) QueryParamsOption {
	return func(m map[string]string) {
		m["page"] = strconv.Itoa(pageNumber)
	}
}

// WithPageSize 每页资源数量。如果不配置该参数，则展示所有查询到的资源。
func WithPageSize(pageSize int) QueryParamsOption {
	return func(m map[string]string) {
		m["page_size"] = strconv.Itoa(pageSize)
	}
}

// WithName 根据资源的 name 属性进行查询，如果资源本身没有 name 属性则不会出现在查询结果中。
func WithName(name string) QueryParamsOption {
	return func(m map[string]string) {
		m["name"] = name
	}
}

// WithLabelKey 根据资源的 label 属性进行查询，如果资源本身没有 label 属性则不会出现在查询结果中。
func WithLabelKey(key string) QueryParamsOption {
	return func(m map[string]string) {
		m["label"] = key
	}
}

// WithRouteUri 该参数仅在 Route 资源上支持。如果 Route 的 uri 等于查询的 uri 或 uris 包含查询的 uri，则该 Route 资源出现在查询结果中。
func WithRouteUri(uri string) QueryParamsOption {
	return func(m map[string]string) {
		m["uri"] = uri
	}
}

type requestMethod string

const (
	getMethod    requestMethod = "get"
	putMethod    requestMethod = "put"
	postMethod   requestMethod = "post"
	patchMethod  requestMethod = "patch"
	deleteMethod requestMethod = "delete"
)

// 发送HTTP请求
func (c *Client) do(method requestMethod, url string, bytes []byte, options ...QueryParamsOption) ([]byte, error) {
	var (
		resp *resty.Response
		err  error
	)

	fullUrl := c.Address + url
	r := resty.New().
		SetTimeout(time.Duration(RequestTimeout)*time.Millisecond).
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-API-KEY", c.Token)

	switch method {
	case getMethod:
		if len(options) != 0 {
			qs := make(map[string]string)
			for _, option := range options {
				option(qs)
			}

			r.SetQueryParams(qs)
		}
		resp, err = r.Get(fullUrl)
	case postMethod:
		if bytes != nil {
			r.SetBody(bytes)
		}
		resp, err = r.Post(fullUrl)
	case putMethod:
		if bytes != nil {
			r.SetBody(bytes)
		}
		resp, err = r.Put(fullUrl)
	case deleteMethod:
		resp, err = r.Delete(fullUrl)
	case patchMethod:
		if bytes != nil {
			r.SetBody(bytes)
		}
		resp, err = r.Patch(fullUrl)
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	return resp.Body(), nil
}
