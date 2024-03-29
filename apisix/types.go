package apisix

// DeleteItemResp 删除条目
type DeleteItemResp struct {
	Key     string `json:"key,omitempty"`
	Deleted string `json:"deleted,omitempty"`
}

/*
SSL 相关的类型定义
*/

type sslItems struct {
	Total int       `json:"total,omitempty"`
	List  []sslItem `json:"list,omitempty"`
}

type sslItem struct {
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
	SSLEnable  SSLStatus = 1
)

type SSLProtocol string

const (
	TLSv11 SSLProtocol = "TLSv1.1"
	TLSv12 SSLProtocol = "TLSv1.2"
	TLSv13 SSLProtocol = "TLSv1.3"
)

type SSLType string

const (
	SSLClient SSLType = "client"
	SSLServer SSLType = "server"
)

type createSSL struct {
	Key          string            `json:"key"`
	Cert         string            `json:"cert"`
	Snis         []string          `json:"snis"`
	Status       *SSLStatus        `json:"status,omitempty"`
	SSLProtocols []SSLProtocol     `json:"ssl_protocols,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	Type         SSLType           `json:"type,omitempty"`
}

/*
Upstream 相关的类型定义
*/

type upstreamItems struct {
	Total int            `json:"total,omitempty"`
	List  []upstreamItem `json:"list"`
}

type upstreamItem struct {
	Key           string    `json:"key,omitempty"`
	Value         *Upstream `json:"value,omitempty"`
	ModifiedIndex int64     `json:"modifiedIndex,omitempty"`
	CreatedIndex  int64     `json:"createdIndex,omitempty"`
}

type Upstream struct {
	ID            string                `json:"id,omitempty"`
	Name          string                `json:"name,omitempty"`
	Desc          string                `json:"desc,omitempty"`
	Type          LoadBalancerType      `json:"type,omitempty"`
	Scheme        Scheme                `json:"scheme,omitempty"`
	ServiceName   string                `json:"service_name,omitempty"`
	DiscoveryType UpstreamDiscoveryType `json:"discovery_type,omitempty"`
	HashOn        UpstreamHashOn        `json:"hash_on,omitempty"`
	HashKey       UpstreamHashKey       `json:"hash_key,omitempty"`
	Labels        map[string]string     `json:"labels,omitempty"`
	Nodes         map[string]int64      `json:"nodes,omitempty"`
	PassHost      UpstreamPassHost      `json:"pass_host,omitempty"`
	Retries       int64                 `json:"retries,omitempty"`
	RetryTimeout  int64                 `json:"retry_timeout,omitempty"`
	Timeout       *Timeout              `json:"timeout,omitempty"`
	KeepalivePool *KeepalivePool        `json:"keepalive_pool,omitempty"`
	Checks        *Checks               `json:"checks,omitempty"`
	TLS           *UpstreamTLS          `json:"tls,omitempty"`
}

type LoadBalancerType string

const (
	RoundRobin LoadBalancerType = "roundrobin"
	CHash      LoadBalancerType = "chash"
	EWMA       LoadBalancerType = "ewma"
	LeastConn  LoadBalancerType = "least_conn"
)

type Scheme string

const (
	HTTP  Scheme = "http"
	HTTPS Scheme = "https"
	GRPC  Scheme = "grpc"
	GRPCS Scheme = "grpcs"
	TCP   Scheme = "tcp"
	UDP   Scheme = "udp"
	TLS   Scheme = "tls"
)

type UpstreamPassHost string

const (
	PASS UpstreamPassHost = "pass"
	NODE UpstreamPassHost = "node"
)

type UpstreamNode struct {
	Host   string `json:"host,omitempty"`
	Port   int64  `json:"port,omitempty"`
	Weight int64  `json:"weight,omitempty"`
}

type UpstreamTLS struct {
	ClientKey    string `json:"client_key,omitempty"`
	ClientCert   string `json:"client_cert,omitempty"`
	ClientCertId string `json:"client_cert_id,omitempty"`
}

type UpstreamHashOn string

const (
	VARS     UpstreamHashOn = "vars"
	HEADER   UpstreamHashOn = "header"
	COOKIE   UpstreamHashOn = "cookie"
	CONSUMER UpstreamHashOn = "consumer"
)

type UpstreamHashKey string

const (
	RemoteAddr  UpstreamHashKey = "remote_addr"
	RemotePort  UpstreamHashKey = "remote_port"
	URI         UpstreamHashKey = "uri"
	ServerName  UpstreamHashKey = "server_name"
	ServerAddr  UpstreamHashKey = "server_addr"
	RequestUri  UpstreamHashKey = "request_uri"
	QueryString UpstreamHashKey = "query_string"
	Host        UpstreamHashKey = "host"
	HostName    UpstreamHashKey = "hostname"
)

type UpstreamHashKeyArg string

const (
	Arg UpstreamHashKeyArg = "arg"
)

type UpstreamDiscoveryType string

const (
	DNS        UpstreamDiscoveryType = "dns"
	Consul     UpstreamDiscoveryType = "consul"
	ConsulKV   UpstreamDiscoveryType = "consul_kv"
	Nacos      UpstreamDiscoveryType = "nacos"
	Eureka     UpstreamDiscoveryType = "eureka"
	Kubernetes UpstreamDiscoveryType = "kubernetes"
)

type KeepalivePool struct {
	Requests    int64 `json:"requests,omitempty"`
	IdleTimeout int64 `json:"idle_timeout,omitempty"`
	Size        int64 `json:"size,omitempty"`
}

type Timeout struct {
	Connect float64 `json:"connect,omitempty"`
	Read    float64 `json:"read,omitempty"`
	Send    float64 `json:"send,omitempty"`
}

type Checks struct {
	Active  *ActiveCheck  `json:"active,omitempty"`
	Passive *PassiveCheck `json:"passive,omitempty"`
}

type ActiveCheck struct {
	Type                   string     `json:"type,omitempty"`
	HTTPPath               string     `json:"http_path,omitempty"`
	Host                   string     `json:"host,omitempty"`
	Port                   int64      `json:"port,omitempty"`
	Timeout                int64      `json:"timeout,omitempty"`
	Concurrency            int64      `json:"concurrency,omitempty"`
	HTTPSVerifyCertificate bool       `json:"https_verify_certificate"`
	ReqHeaders             []string   `json:"req_headers,omitempty"`
	Healthy                *Healthy   `json:"healthy,omitempty"`
	Unhealthy              *Unhealthy `json:"unhealthy,omitempty"`
}

type PassiveCheck struct {
	Type      string     `json:"type,omitempty"`
	Healthy   *Healthy   `json:"healthy,omitempty"`
	Unhealthy *Unhealthy `json:"unhealthy,omitempty"`
}

type Healthy struct {
	Interval     int64   `json:"interval,omitempty"`
	HTTPStatuses []int64 `json:"http_statuses,omitempty"`
	Successes    int64   `json:"successes,omitempty"`
}

type Unhealthy struct {
	Timeouts     int64   `json:"timeouts,omitempty"`
	Interval     int64   `json:"interval,omitempty"`
	HTTPFailures int64   `json:"http_failures,omitempty"`
	HTTPStatuses []int64 `json:"http_statuses,omitempty"`
	TCPFailures  int64   `json:"tcp_failures,omitempty"`
}

/*
HTTP METHOD 相关的类型定义
*/

type HTTPMethod string

const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	DELETE  HTTPMethod = "DELETE"
	PATCH   HTTPMethod = "PATCH"
	HEAD    HTTPMethod = "HEAD"
	OPTIONS HTTPMethod = "OPTIONS"
	CONNECT HTTPMethod = "CONNECT"
	TRACE   HTTPMethod = "TRACE"
	PURGE   HTTPMethod = "PURGE"
)

/*
Route 相关的类型定义
*/

type routeItems struct {
	Total int64       `json:"total,omitempty"`
	List  []routeItem `json:"list,omitempty"`
}

type routeItem struct {
	Key           string `json:"key,omitempty"`
	Value         *Route `json:"value,omitempty"`
	ModifiedIndex int64  `json:"modifiedIndex,omitempty"`
	CreatedIndex  int64  `json:"createdIndex,omitempty"`
}

type Route struct {
	Id              string            `json:"id,omitempty"`
	Uri             string            `json:"uri,omitempty"`
	Uris            []string          `json:"uris,omitempty"`
	Name            string            `json:"name,omitempty"`
	Desc            string            `json:"desc,omitempty"`
	Priority        int64             `json:"priority,omitempty"`
	Methods         []HTTPMethod      `json:"methods,omitempty"`
	Host            string            `json:"host,omitempty"`
	Hosts           []string          `json:"hosts,omitempty"`
	RemoteAddrs     []string          `json:"remote_addrs,omitempty"`
	Timeout         *Timeout          `json:"timeout,omitempty"`
	Plugins         map[string]any    `json:"plugins,omitempty"`
	ServiceId       string            `json:"service_id,omitempty"`
	Upstream        *Upstream         `json:"upstream,omitempty"`
	UpstreamId      string            `json:"upstream_id,omitempty"`
	Status          RouteStatus       `json:"status,omitempty"`
	EnableWebsocket bool              `json:"enable_websocket,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"  `
	CreateTime      int64             `json:"create_time,omitempty"`
	UpdateTime      int64             `json:"update_time,omitempty"`
}

type RouteStatus int64

const (
	RouteDisable RouteStatus = 0
	RouteEnable  RouteStatus = 1
)

/*
Service 相关的类型定义
*/

type serviceItems struct {
	Total int64         `json:"total,omitempty"`
	List  []serviceItem `json:"list,omitempty"`
}

type serviceItem struct {
	Key           string   `json:"key,omitempty"`
	Value         *Service `json:"value,omitempty"`
	ModifiedIndex int64    `json:"modifiedIndex,omitempty"`
	CreatedIndex  int64    `json:"createdIndex,omitempty"`
}

type Service struct {
	ID              string            `json:"id,omitempty"`
	Name            string            `json:"name,omitempty"`
	Desc            string            `json:"desc,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"`
	EnableWebsocket bool              `json:"enable_websocket,omitempty"`
	Upstream        *Upstream         `json:"upstream,omitempty"`
	UpstreamId      string            `json:"upstream_id,omitempty"`
	Hosts           []string          `json:"hosts,omitempty"`
	Plugins         map[string]any    `json:"plugins,omitempty"`
	CreateTime      int64             `json:"create_time,omitempty"`
	UpdateTime      int64             `json:"update_time,omitempty"`
}

/*
GlobalRule 类型定义
*/

type globalruleItems struct {
	Total int64            `json:"total,omitempty"`
	List  []globalruleItem `json:"list,omitempty"`
}

type globalruleItem struct {
	Key           string      `json:"key,omitempty"`
	Value         *GlobalRule `json:"value,omitempty"`
	ModifiedIndex int64       `json:"modifiedIndex,omitempty"`
	CreatedIndex  int64       `json:"createdIndex,omitempty"`
}

type GlobalRule struct {
	ID         string                          `json:"id,omitempty"`
	Plugins    map[string]GlobalRulePluginMeta `json:"plugins,omitempty"`
	CreateTime int64                           `json:"create_time,omitempty"`
	UpdateTime int64                           `json:"update_time,omitempty"`
}

type GlobalRulePluginMeta struct {
	Meta struct {
		Disable bool `json:"disable,omitempty"`
	} `json:"_meta,omitempty"`
}

/*
Plugin 类型定义
*/

type PluginSubsystem string

const (
	HTTPSubsystem   PluginSubsystem = "http"
	StreamSubsystem PluginSubsystem = "stream"
)

type PluginAttribute struct {
	Comment    string   `json:"$comment,omitempty"`
	Properties []string `json:"properties,omitempty"`
	Type       string   `json:"type,omitempty"`
	Required   []string `json:"required,omitempty"`
}

/*
StreamRoute 类型定义
*/
type streamRouteItems struct {
	Total int64             `json:"total,omitempty"`
	List  []streamRouteItem `json:"list,omitempty"`
}

type streamRouteItem struct {
	Key           string       `json:"key,omitempty"`
	Value         *StreamRoute `json:"value,omitempty"`
	ModifiedIndex int64        `json:"modifiedIndex,omitempty"`
	CreatedIndex  int64        `json:"createdIndex,omitempty"`
}

// StreamRoute TODO: 不支持xrpc协议
type StreamRoute struct {
	ID         string    `json:"id,omitempty"`
	ServerAddr string    `json:"server_addr,omitempty"`
	ServerPort int64     `json:"server_port,omitempty"`
	RemoteAddr string    `json:"remote_addr,omitempty"`
	Upstream   *Upstream `json:"upstream,omitempty"`
	UpstreamId string    `json:"upstream_id,omitempty"`
	ServiceId  string    `json:"service_id,omitempty"`
	Sni        string    `json:"sni,omitempty"`
	CreateTime int64     `json:"create_time,omitempty"`
	UpdateTime int64     `json:"update_time,omitempty"`
}
