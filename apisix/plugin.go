package apisix

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
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

// PluginReload 热加载插件
func (c *Client) PluginReload() (bool, error) {
	resp, err := c.do(putMethod, "/plugins/reload", nil)
	if err != nil {
		return false, err
	}

	if string(resp) != "done" {
		return false, fmt.Errorf("reload plugins err: %s", string(resp))
	} else {
		return true, nil
	}
}

// GetPluginAttribute 获取插件属性
func (c *Client) GetPluginAttribute(name string, subsystem PluginSubsystem) (*PluginAttribute, error) {
	resp, err := c.do(getMethod, fmt.Sprintf("/plugins/%s?subsystem=%s", name, subsystem), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(resp))
	var pa PluginAttribute
	pa.Type = gjson.Get(string(resp), "type").String()
	pa.Comment = gjson.Get(string(resp), "$comment").String()

	r := make([]string, 0)
	for _, result := range gjson.Get(string(resp), "required").Array() {
		r = append(r, result.String())
	}
	pa.Required = r

	p := make([]string, 0)
	for _, result := range gjson.Get(string(resp), "properties.@keys").Array() {
		p = append(p, result.String())
	}

	pa.Properties = p

	return &pa, nil
}
