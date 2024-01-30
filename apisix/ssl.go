package apisix

import (
	"encoding/json"
	"fmt"
)

// GetSSLs 获取所有SSL证书
func (c *Client) GetSSLs(options ...QueryParamsOption) ([]*SSL, error) {
	var (
		gssls sslItems
		ssls  []*SSL
		err   error
		resp  []byte
	)

	if len(options) != 0 {
		resp, err = c.do(getMethod, "/ssls", nil, options...)
	} else {
		resp, err = c.do(getMethod, "/ssls", nil)
	}

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &gssls)
	if err != nil {
		return ssls, err
	}

	for _, ssl := range gssls.List {
		ssls = append(ssls, ssl.Value)
	}

	return ssls, nil
}

// GetSSL 获取指定SSL证书的信息
func (c *Client) GetSSL(id string) (*SSL, error) {
	var si sslItem

	resp, err := c.do(getMethod, fmt.Sprintf("/ssls/%s", id), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &si)
	if err != nil {
		return nil, err
	}
	return si.Value, nil
}

// DeleteSSL 删除指定的SSL证书
func (c *Client) DeleteSSL(id string) (*DeleteItemResp, error) {
	var di DeleteItemResp
	resp, err := c.do(deleteMethod, fmt.Sprintf("/ssls/%s", id), nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resp, &di); err != nil {
		return &di, err
	} else {
		return &di, nil
	}
}

type CreateSSLOption func(*createSSL)

func CreateSSLWithStatus(status SSLStatus) CreateSSLOption {
	return func(c *createSSL) {
		c.Status = &status
	}
}

func CreateSSLWithSSLProtocols(protocols []SSLProtocol) CreateSSLOption {
	return func(c *createSSL) {
		c.SSLProtocols = protocols
	}
}

func CreateSSLWithLabels(labels map[string]string) CreateSSLOption {
	return func(c *createSSL) {
		c.Labels = labels
	}
}

func CreateSSLWithSSLType(t SSLType) CreateSSLOption {
	return func(c *createSSL) {
		c.Type = t
	}
}

// CreateSSL 创建SSL证书
func (c *Client) CreateSSL(key, cert []byte, snis []string, options ...CreateSSLOption) (*SSL, error) {
	cs := &createSSL{
		Key:  string(key),
		Cert: string(cert),
		Snis: snis,
	}

	for _, option := range options {
		option(cs)
	}

	body, err := json.Marshal(cs)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(postMethod, "/ssls", body)
	if err != nil {
		return nil, err
	}

	var si sslItem
	err = json.Unmarshal(resp, &si)
	if err != nil {
		return nil, err
	}

	return si.Value, nil
}
