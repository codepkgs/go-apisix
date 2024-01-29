package apisix

// Get请求获取多条记录
type items struct {
	Total int    `json:"total,omitempty"`
	List  []item `json:"list"`
}

// Get请求获取单条记录
type item struct {
	Key           string `json:"key,omitempty"`
	Value         *SSL   `json:"value,omitempty"`
	ModifiedIndex int64  `json:"modifiedIndex,omitempty"`
	CreatedIndex  int64  `json:"createdIndex,omitempty"`
}

// DeleteItemResp 删除条目
type DeleteItemResp struct {
	Key     string `json:"key,omitempty"`
	Deleted string `json:"deleted,omitempty"`
}

// SSL SSL相关的类型定义
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

type createSSL struct {
	Key          string            `json:"key"`
	Cert         string            `json:"cert"`
	Snis         []string          `json:"snis"`
	Status       *SSLStatus        `json:"status,omitempty"`
	SSLProtocols []SSLProtocol     `json:"ssl_protocols,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	Type         SSLType           `json:"type,omitempty"`
}
