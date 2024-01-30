package apisix

import (
	"encoding/json"
	"fmt"
)

// GetRoutes 获取所有路由或符合筛选条件的路由
func (c *Client) GetRoutes(options ...QueryParamsOption) ([]*Route, error) {
	var (
		ris    routeItems
		routes []*Route
		resp   []byte
		err    error
	)
	if len(options) != 0 {
		resp, err = c.do(getMethod, "/routes", nil, options...)
	} else {
		resp, err = c.do(getMethod, "/routes", nil)
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &ris)
	if err != nil {
		return routes, err
	}

	for _, route := range ris.List {
		routes = append(routes, route.Value)
	}

	return routes, nil
}

// GetRoute 获取指定的路由
func (c *Client) GetRoute(id string) (*Route, error) {
	var (
		ri   routeItem
		resp []byte
		err  error
	)
	resp, err = c.do(getMethod, fmt.Sprintf("/routes/%s", id), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &ri)
	if err != nil {
		return nil, err
	}
	return ri.Value, nil
}

// DeleteRoute 删除指定的路由
func (c *Client) DeleteRoute(id string) (*DeleteItemResp, error) {
	resp, err := c.do(deleteMethod, fmt.Sprintf("/routes/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var di DeleteItemResp
	if err = json.Unmarshal(resp, &di); err != nil {
		return &di, err
	} else {
		return &di, nil
	}
}

// CreateOrModifyRouteOption 创建或修改路由的选项
type CreateOrModifyRouteOption func(*Route)

// CreateOrModifyRouteWithName 修改路由名称
func CreateOrModifyRouteWithName(name string) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Name = name
	}
}

// CreateOrModifyRouteWithLabels 创建或修改路由标签
func CreateOrModifyRouteWithLabels(labels map[string]string) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Labels = labels
	}
}

// CreateOrModifyRouteWithDesc 创建或修改路由描述
func CreateOrModifyRouteWithDesc(desc string) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Desc = desc
	}
}

// CreateOrModifyRouteWithUri 创建或修改路由匹配的URI
func CreateOrModifyRouteWithUri(uri string) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Uri = uri
	}
}

// CreateOrModifyRouteWithUris 创建或修改路由匹配的URI列表.可以匹配多个URI
func CreateOrModifyRouteWithUris(uris []string) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Uris = uris
	}
}

// CreateOrModifyRouteWithUpstream 创建或修改路由的Upstream
func CreateOrModifyRouteWithUpstream(upstream *Upstream) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Upstream = upstream
	}
}

// CreateOrModifyRouteWithUpstreamId 创建或修改路由的UpstreamId
func CreateOrModifyRouteWithUpstreamId(upstreamId string) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.UpstreamId = upstreamId
	}
}

// CreateOrModifyRouteWithServiceId 创建或修改路由的ServiceId
func CreateOrModifyRouteWithServiceId(serviceId string) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.ServiceId = serviceId
	}
}

// CreateOrModifyRouteWithHost 创建或修改路由当前请求的域名
func CreateOrModifyRouteWithHost(host string) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Host = host
	}
}

// CreateOrModifyRouteWithHosts 创建或修改路由当前请求的域名列表
func CreateOrModifyRouteWithHosts(hosts []string) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Hosts = hosts
	}
}

// CreateOrModifyRouteWithMethods 创建或修改路由允许的方法列表.如果为空或没有该选项，则表示没有任何 method 限制。
func CreateOrModifyRouteWithMethods(methods []HTTPMethod) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Methods = methods
	}
}

// CreateOrModifyRouteWithPriority 创建或修改路由的优先级.如果不同路由包含相同的 uri，则根据属性 priority 确定哪个 route 被优先匹配，值越大优先级越高，默认值为 0。
func CreateOrModifyRouteWithPriority(priority int64) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Priority = priority
	}
}

// CreateOrModifyRouteWithStatus 创建或修改路由的状态. 1启用,0禁用
func CreateOrModifyRouteWithStatus(status RouteStatus) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Status = status
	}
}

// CreateOrModifyRouteWithEnableWebSocket 创建或修改路由的web socket.当设置为 true 时，启用 websocket(boolean), 默认值为 false。
func CreateOrModifyRouteWithEnableWebSocket(enabled bool) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.EnableWebsocket = enabled
	}
}

// CreateOrModifyRouteWithTimeout 创建或修改路由的timeout。
// 为 Route 设置 Upstream 连接、发送消息和接收消息的超时时间（单位为秒）。该配置将会覆盖在 Upstream 中配置的 timeout 选项。
func CreateOrModifyRouteWithTimeout(timeout *Timeout) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Timeout = timeout
	}
}

// CreateOrModifyRouteWithPlugins 设置插件
func CreateOrModifyRouteWithPlugins(plugins map[string]any) CreateOrModifyRouteOption {
	return func(r *Route) {
		r.Plugins = plugins
	}
}

// CreateRoute 创建路由
func (c *Client) CreateRoute(name string, options ...CreateOrModifyRouteOption) (*Route, error) {
	r := &Route{
		Name: name,
	}
	for _, option := range options {
		option(r)
	}

	body, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(postMethod, "/routes", body)
	if err != nil {
		return nil, err
	}

	var ri routeItem
	err = json.Unmarshal(resp, &ri)
	if err != nil {
		return nil, err
	}

	return ri.Value, nil
}

// ModifyRoute 修改Route
func (c *Client) ModifyRoute(id string, options ...CreateOrModifyRouteOption) (*Route, error) {
	r, err := c.GetRoute(id)
	if err != nil {
		return r, err
	}

	for _, option := range options {
		option(r)
	}

	body, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(putMethod, fmt.Sprintf("/routes/%s", id), body)
	if err != nil {
		return nil, err
	}

	var ri routeItem
	err = json.Unmarshal(resp, &ri)
	if err != nil {
		return nil, err
	}

	return ri.Value, nil
}
