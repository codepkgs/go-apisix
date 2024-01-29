package apisix

import (
	"encoding/json"
	"fmt"
)

type sslsResp struct {
	Total int       `json:"total,omitempty"`
	List  []sslResp `json:"list"`
}

type sslResp struct {
	Key           string `json:"key,omitempty"`
	Value         *SSL   `json:"value,omitempty"`
	ModifiedIndex int64  `json:"modifiedIndex,omitempty"`
	CreatedIndex  int64  `json:"createdIndex,omitempty"`
}

type SSL struct {
	ID            string            `json:"id,omitempty"`
	Key           string            `json:"key,omitempty"`
	Cert          string            `json:"cert,omitempty"`
	Status        SSLStatus         `json:"status,omitempty"`
	Snis          []string          `json:"snis,omitempty"`
	Type          string            `json:"type,omitempty"`
	ValidityEnd   int64             `json:"validity_end,omitempty"`
	ValidityStart int64             `json:"validity_start,omitempty"`
	CreateTime    int64             `json:"create_time,omitempty"`
	UpdateTime    int64             `json:"update_time,omitempty"`
	Labels        map[string]string `json:"labels,omitempty"  `
}

type SSLStatus int64

const (
	SSLDisable SSLStatus = 0
	SSLEnable            = 1
)

type SSLProtocol string

const (
	TLSv11 SSLProtocol = "TLSv1.1"
	TLSv12             = "TLSv1.2"
	TLSv13             = "TLSv1.3"
)

type SSLType string

const (
	SSLClient SSLType = "client"
	SSLServer         = "server"
)

type DeleteSSLResp struct {
	Key     string `json:"key,omitempty"`
	Deleted string `json:"deleted,omitempty"`
}

type createSSL struct {
	Key          string            `json:"key"`
	Cert         string            `json:"cert"`
	Snis         []string          `json:"snis"`
	Status       *SSLStatus        `json:"status,omitempty"`
	SSLProtocols []SSLProtocol     `json:"ssl_protocols,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	Type         SSLType           `json:"type,omitempty"`
}

// GetSSLs 获取所有SSL证书
func (c *Client) GetSSLs() ([]*SSL, error) {
	var (
		gssls sslsResp
		ssls  []*SSL
	)

	resp, err := c.get("/ssls")
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
	var sr sslResp

	resp, err := c.get(fmt.Sprintf("/ssls/%s", id))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &sr)
	if err != nil {
		return sr.Value, err
	}
	return sr.Value, nil
}

// DeleteSSL 删除指定的SSL证书
func (c *Client) DeleteSSL(id string) (*DeleteSSLResp, error) {
	var dsr DeleteSSLResp
	resp, err := c.delete(fmt.Sprintf("/ssls/%s", id))
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resp, &dsr); err != nil {
		return &dsr, err
	} else {
		return &dsr, nil
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

	resp, err := c.post("/ssls", body)
	if err != nil {
		return nil, err
	}

	var sr sslResp
	err = json.Unmarshal(resp, &sr)
	if err != nil {
		return nil, err
	}

	return sr.Value, nil
}
