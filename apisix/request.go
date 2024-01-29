package apisix

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

const (
	RequestTimeout = 5000
)

// get GET请求
func (c *Client) get(url string) ([]byte, error) {
	fullUrl := c.Address + url
	r := resty.New().
		SetTimeout(time.Duration(RequestTimeout)*time.Millisecond).
		R().
		SetHeader("Content-Type", "application/json").
		SetHeaderVerbatim("X-API-KEY", c.Token)

	resp, err := r.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode(), resp.Body())
	}
	return resp.Body(), nil
}

// post POST请求
func (c *Client) post(url string, bytes []byte) ([]byte, error) {
	fullUrl := c.Address + url
	r := resty.New().
		SetTimeout(time.Duration(RequestTimeout)*time.Millisecond).
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-API-KEY", c.Token)

	r.SetBody(bytes)

	resp, err := r.Post(fullUrl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode(), resp.Body())
	}
	return resp.Body(), nil
}

// delete DELETE请求
func (c *Client) delete(url string) ([]byte, error) {
	fullUrl := c.Address + url
	r := resty.New().
		SetTimeout(time.Duration(RequestTimeout)*time.Millisecond).
		R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-API-KEY", c.Token)

	resp, err := r.Delete(fullUrl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode(), resp.Body())
	}
	return resp.Body(), nil
}
