package apisix

import (
	"encoding/json"
	"fmt"
)

// GetServices 获取符合条件的Services或所有的Services
func (c *Client) GetServices(options ...QueryParamsOption) ([]*Service, error) {
	var (
		sis  serviceItems
		ss   []*Service
		err  error
		resp []byte
	)

	if len(options) != 0 {
		resp, err = c.do(getMethod, "/services", nil, options...)
	} else {
		resp, err = c.do(getMethod, "/services", nil)
	}

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &sis)
	if err != nil {
		return nil, err
	}

	for _, s := range sis.List {
		ss = append(ss, s.Value)
	}

	return ss, nil
}

// GetService 获取指定Service的信息
func (c *Client) GetService(id string) (*Service, error) {
	var si serviceItem

	resp, err := c.do(getMethod, fmt.Sprintf("/services/%s", id), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &si)
	if err != nil {
		return nil, err
	}
	return si.Value, nil
}

// DeleteService 删除指定的Service。如果service被route引用，则会删除失败
func (c *Client) DeleteService(id string) (*DeleteItemResp, error) {
	var di DeleteItemResp
	resp, err := c.do(deleteMethod, fmt.Sprintf("/services/%s", id), nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resp, &di); err != nil {
		return &di, err
	} else {
		return &di, nil
	}
}

// CreateOrModifyServiceOption 创建或修改Service选项
type CreateOrModifyServiceOption func(*Service)

// CreateOrModifyServiceWithName Service的名称
func CreateOrModifyServiceWithName(name string) CreateOrModifyServiceOption {
	return func(s *Service) {
		s.Name = name
	}
}

// CreateOrModifyServiceWithDesc Service的描述
func CreateOrModifyServiceWithDesc(desc string) CreateOrModifyServiceOption {
	return func(s *Service) {
		s.Desc = desc
	}
}

// CreateOrModifyServiceWithUpstream Service关联的Upstream
func CreateOrModifyServiceWithUpstream(upstream *Upstream) CreateOrModifyServiceOption {
	return func(s *Service) {
		s.Upstream = upstream
	}
}

// CreateOrModifyServiceWithUpstreamId Service关联的UpstreamId
func CreateOrModifyServiceWithUpstreamId(upstreamId string) CreateOrModifyServiceOption {
	return func(s *Service) {
		s.UpstreamId = upstreamId
	}
}

// CreateOrModifyServiceWithLabels Service标签
func CreateOrModifyServiceWithLabels(labels map[string]string) CreateOrModifyServiceOption {
	return func(s *Service) {
		s.Labels = labels
	}
}

// CreateOrModifyServiceWithHosts 允许的host列表
func CreateOrModifyServiceWithHosts(hosts []string) CreateOrModifyServiceOption {
	return func(s *Service) {
		s.Hosts = hosts
	}
}

// CreateOrModifyServiceWithEnableWebsocket 是否开启websocket
func CreateOrModifyServiceWithEnableWebsocket(enableWebsocket bool) CreateOrModifyServiceOption {
	return func(s *Service) {
		s.EnableWebsocket = enableWebsocket
	}
}

// CreateService 创建Service
func (c *Client) CreateService(name string, options ...CreateOrModifyServiceOption) (*Service, error) {
	s := &Service{
		Name: name,
	}
	for _, option := range options {
		option(s)
	}

	body, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(postMethod, "/services", body)
	if err != nil {
		return nil, err
	}

	var si serviceItem
	err = json.Unmarshal(resp, &si)
	if err != nil {
		return nil, err
	}

	return si.Value, nil
}

// ModifyService 修改Service
func (c *Client) ModifyService(id string, options ...CreateOrModifyServiceOption) (*Service, error) {
	s, err := c.GetService(id)
	if err != nil {
		return s, err
	}

	for _, option := range options {
		option(s)
	}

	body, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(putMethod, fmt.Sprintf("/services/%s", id), body)
	if err != nil {
		return nil, err
	}

	var si serviceItem
	err = json.Unmarshal(resp, &si)
	if err != nil {
		return nil, err
	}

	return si.Value, nil
}
