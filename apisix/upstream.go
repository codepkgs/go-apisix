package apisix

import (
	"encoding/json"
	"fmt"
)

// GetUpstreams 获取符合条件的Upstreams
func (c *Client) GetUpstreams(options ...QueryParamsOption) ([]*Upstream, error) {
	var (
		upis upstreamItems
		ups  []*Upstream
		err  error
		resp []byte
	)

	if len(options) != 0 {
		resp, err = c.do(getMethod, "/upstreams", nil, options...)
	} else {
		resp, err = c.do(getMethod, "/upstreams", nil)
	}

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &upis)
	if err != nil {
		return ups, err
	}

	for _, up := range upis.List {
		ups = append(ups, up.Value)
	}

	return ups, nil
}

// GetUpstream 获取指定的Upstream信息
func (c *Client) GetUpstream(id string) (*Upstream, error) {
	resp, err := c.do(getMethod, fmt.Sprintf("/upstreams/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var up upstreamItem
	err = json.Unmarshal(resp, &up)
	if err != nil {
		return nil, err
	}

	return up.Value, nil
}

// DeleteUpstream 删除指定的Upstream
func (c *Client) DeleteUpstream(id string) (*DeleteItemResp, error) {
	resp, err := c.do(deleteMethod, fmt.Sprintf("/upstreams/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var di DeleteItemResp
	err = json.Unmarshal(resp, &di)
	if err != nil {
		return nil, err
	} else {
		return &di, nil
	}
}

// CreateOrModifyUpstreamOption 创建或修改Upstream的选项
type CreateOrModifyUpstreamOption func(*Upstream)

// CreateOrModifyUpstreamWithName Upstream的名称。
func CreateOrModifyUpstreamWithName(name string) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.Name = name
	}
}

// CreateOrModifyUpstreamWithLoadBalancerType 负载均衡的类型。
func CreateOrModifyUpstreamWithLoadBalancerType(lbType LoadBalancerType) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.Type = lbType
	}
}

// CreateOrModifyUpstreamWithDesc 上游服务描述、使用场景等。
func CreateOrModifyUpstreamWithDesc(desc string) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.Desc = desc
	}
}

// CreateOrModifyUpstreamWithNodes 上游服务器信息
func CreateOrModifyUpstreamWithNodes(nodes []UpstreamNode) CreateOrModifyUpstreamOption {
	nodesMap := make(map[string]int64, len(nodes))
	for _, node := range nodes {
		nodesMap[fmt.Sprintf("%s:%d", node.Host, node.Port)] = node.Weight
	}

	return func(u *Upstream) {
		u.Nodes = nodesMap
	}
}

// CreateOrModifyUpstreamWithServiceName 服务发现时使用的服务名
func CreateOrModifyUpstreamWithServiceName(serviceName string) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.ServiceName = serviceName
	}
}

func CreateOrModifyUpstreamWithDiscoveryType(discoveryType UpstreamDiscoveryType) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.DiscoveryType = discoveryType
	}
}

// CreateOrModifyUpstreamWithScheme 跟上游通信时使用的 scheme。默认值为 http
func CreateOrModifyUpstreamWithScheme(scheme Scheme) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.Scheme = scheme
	}
}

// CreateOrModifyUpstreamWithRetries 使用 NGINX 重试机制将请求传递给下一个上游，默认启用重试机制且次数为后端可用的节点数量。
// 如果指定了具体重试次数，它将覆盖默认值。当设置为 0 时，表示不启用重试机制。
func CreateOrModifyUpstreamWithRetries(retries int64) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.Retries = retries
	}
}

// CreateOrModifyUpstreamWithRetryTimeout 限制是否继续重试的时间，若之前的请求和重试请求花费太多时间就不再继续重试。当设置为 0 时，表示不启用重试超时机制。
func CreateOrModifyUpstreamWithRetryTimeout(retryTimeout int64) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.RetryTimeout = retryTimeout
	}
}

// CreateOrModifyUpstreamWithTimeout 设置连接、发送消息、接收消息的超时时间，以秒为单位。
func CreateOrModifyUpstreamWithTimeout(timeout Timeout) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.Timeout = &timeout
	}
}

// CreateOrModifyUpstreamWithLabels 设置标签
func CreateOrModifyUpstreamWithLabels(labels map[string]string) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.Labels = labels
	}
}

// CreateOrModifyUpstreamWithPassHost 请求发给上游时的 host 设置选型。 [pass，node] 之一，默认是 pass。
// pass: 将客户端的 host 透传给上游； node: 使用 upstream node 中配置的 host；
func CreateOrModifyUpstreamWithPassHost(passHost UpstreamPassHost) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.PassHost = passHost
	}
}

// CreateOrModifyUpstreamWithKeepalivePool 设置KeepalivePool
func CreateOrModifyUpstreamWithKeepalivePool(keepalivePool KeepalivePool) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.KeepalivePool = &keepalivePool
	}
}

// CreateOrModifyUpstreamWithClientTLS 当 scheme 为 https或grpcs时，需要指定
func CreateOrModifyUpstreamWithClientTLS(tls UpstreamTLS) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.TLS = &tls
	}
}

// CreateOrModifyUpstreamWithHashOn 根据什么来计算hash。hash_on 支持的类型有 vars（NGINX 内置变量），header（自定义 header），cookie，consumer，默认值为 vars。
func CreateOrModifyUpstreamWithHashOn(hashOn UpstreamHashOn) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.HashOn = hashOn
	}
}

// CreateOrModifyUpstreamWithHashKey 当负载均衡策略为CHash时，指定hash key。根据 key 来查找对应的节点 id，相同的 key 在同一个对象中，则返回相同 id
func CreateOrModifyUpstreamWithHashKey(key UpstreamHashKey) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.HashKey = key
	}
}

// CreateOrModifyUpstreamWithHashKeyArg 当负载均衡策略为CHash时，指定hash key。 只有当 key 为 Arg 时，arg参数为url请求参数。
func CreateOrModifyUpstreamWithHashKeyArg(key UpstreamHashKeyArg, arg string) CreateOrModifyUpstreamOption {
	key += UpstreamHashKeyArg("_" + arg)
	return func(u *Upstream) {
		u.HashKey = UpstreamHashKey(key)
	}
}

// CreateOrModifyUpstreamWithChecks 配置健康检查的参数
func CreateOrModifyUpstreamWithChecks(checks *Checks) CreateOrModifyUpstreamOption {
	return func(u *Upstream) {
		u.Checks = checks
	}
}

// CreateUpstream 创建Upstream
func (c *Client) CreateUpstream(name string, options ...CreateOrModifyUpstreamOption) (*Upstream, error) {
	up := &Upstream{
		Name: name,
	}

	for _, option := range options {
		option(up)
	}

	body, err := json.Marshal(up)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(postMethod, "/upstreams", body)
	if err != nil {
		return nil, err
	}

	var upi upstreamItem
	err = json.Unmarshal(resp, &upi)
	if err != nil {
		return nil, err
	}

	return upi.Value, nil
}

// ModifyUpstream 修改Upstream
func (c *Client) ModifyUpstream(id string, options ...CreateOrModifyUpstreamOption) (*Upstream, error) {
	up, err := c.GetUpstream(id)
	if err != nil {
		return up, err
	}

	for _, option := range options {
		option(up)
	}

	body, err := json.Marshal(up)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(putMethod, fmt.Sprintf("/upstreams/%s", id), body)
	if err != nil {
		return nil, err
	}

	var upi upstreamItem
	err = json.Unmarshal(resp, &upi)
	if err != nil {
		return nil, err
	}

	return upi.Value, nil
}
