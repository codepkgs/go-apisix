package apisix

import (
	"encoding/json"
)

// GetPlugins 获取插件列表
func (c *Client) GetPlugins() ([]string, error) {
	resp, err := c.do(getMethod, "/plugins/list", nil)
	if err != nil {
		return nil, err
	}

	var plugins []string
	if err = json.Unmarshal(resp, &plugins); err != nil {
		return nil, err
	} else {
		return plugins, nil
	}
}
