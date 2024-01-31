package apisix

import (
	"encoding/json"
	"fmt"
)

// GetStreamRoutes 查询所有的StreamRoute
func (c *Client) GetStreamRoutes(options ...QueryParamsOption) ([]*StreamRoute, error) {
	var (
		sris streamRouteItems
		srs  []*StreamRoute
		err  error
		resp []byte
	)

	if len(options) != 0 {
		resp, err = c.do(getMethod, "/stream_routes", nil, options...)
	} else {
		resp, err = c.do(getMethod, "/stream_routes", nil)
	}

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &sris)
	if err != nil {
		return nil, err
	}

	for _, s := range sris.List {
		srs = append(srs, s.Value)
	}

	return srs, nil
}

// GetStreamRoute 查询指定ID的StreamRoute信息
func (c *Client) GetStreamRoute(id string) (*StreamRoute, error) {
	var sri streamRouteItem

	resp, err := c.do(getMethod, fmt.Sprintf("/stream_routes/%s", id), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &sri)
	if err != nil {
		return nil, err
	}
	return sri.Value, nil
}

// DeleteStreamRoute 删除指定ID的StreamRoute
func (c *Client) DeleteStreamRoute(id string) (*DeleteItemResp, error) {
	resp, err := c.do(deleteMethod, fmt.Sprintf("/stream_routes/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var di DeleteItemResp
	if err = json.Unmarshal(resp, &di); err != nil {
		return nil, err
	} else {
		return &di, nil
	}
}

// CreateOrModifyStreamRouteOption 创建或修改StreamRoute时的参数
type CreateOrModifyStreamRouteOption func(*StreamRoute)

// CreateOrModifyStreamRouteWithRemoteAddr 指定RemoteAddr
func CreateOrModifyStreamRouteWithRemoteAddr(remoteAddr string) CreateOrModifyStreamRouteOption {
	return func(r *StreamRoute) {
		r.RemoteAddr = remoteAddr
	}
}

// CreateOrModifyStreamRouteWithServerAddr 指定ServerAddr
func CreateOrModifyStreamRouteWithServerAddr(serverAddr string) CreateOrModifyStreamRouteOption {
	return func(r *StreamRoute) {
		r.ServerAddr = serverAddr
	}
}

// CreateOrModifyStreamRouteWithServerPort 指定ServerPort
func CreateOrModifyStreamRouteWithServerPort(serverPort int64) CreateOrModifyStreamRouteOption {
	return func(r *StreamRoute) {
		r.ServerPort = serverPort
	}
}

// CreateOrModifyStreamRouteWithServiceId 指定ServiceId
func CreateOrModifyStreamRouteWithServiceId(serviceId string) CreateOrModifyStreamRouteOption {
	return func(r *StreamRoute) {
		r.ServiceId = serviceId
	}
}

// CreateOrModifyStreamRouteWithSni 指定SNI host
func CreateOrModifyStreamRouteWithSni(sni string) CreateOrModifyStreamRouteOption {
	return func(r *StreamRoute) {
		r.Sni = sni
	}
}

// CreateOrModifyStreamRouteWithUpstreamId 指定UpstreamId
func CreateOrModifyStreamRouteWithUpstreamId(upstreamId string) CreateOrModifyStreamRouteOption {
	return func(r *StreamRoute) {
		r.UpstreamId = upstreamId
	}
}

// CreateOrModifyStreamRouteWithUpstream 指定Upstream
func CreateOrModifyStreamRouteWithUpstream(upstream *Upstream) CreateOrModifyStreamRouteOption {
	return func(r *StreamRoute) {
		r.Upstream = upstream
	}
}

// CreateStreamRoute 创建StreamRoute
func (c *Client) CreateStreamRoute(options ...CreateOrModifyStreamRouteOption) (*StreamRoute, error) {
	sr := new(StreamRoute)
	for _, option := range options {
		option(sr)
	}

	body, err := json.Marshal(sr)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(postMethod, "/stream_routes", body)
	if err != nil {
		return nil, err
	}

	var sri streamRouteItem
	err = json.Unmarshal(resp, &sri)
	if err != nil {
		return nil, err
	}

	return sri.Value, nil
}

// ModifyStreamRoute 修改StreamRoute
func (c *Client) ModifyStreamRoute(id string, options ...CreateOrModifyStreamRouteOption) (*StreamRoute, error) {
	sr, err := c.GetStreamRoute(id)
	if err != nil {
		return sr, err
	}

	for _, option := range options {
		option(sr)
	}

	body, err := json.Marshal(sr)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(putMethod, fmt.Sprintf("/stream_routes/%s", id), body)
	if err != nil {
		return nil, err
	}

	var sri streamRouteItem
	err = json.Unmarshal(resp, &sri)
	if err != nil {
		return nil, err
	}

	return sri.Value, nil
}
