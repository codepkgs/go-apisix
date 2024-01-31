package apisix

import (
	"encoding/json"
	"fmt"
)

// Global Rule 可以设置全局运行的插件，设置为全局规则的插件将在所有路由级别的插件之前优先运行。

// GetGlobalRules 获取全局插件列表
func (c *Client) GetGlobalRules(options ...QueryParamsOption) ([]*GlobalRule, error) {
	var (
		gris globalruleItems
		grs  []*GlobalRule
		resp []byte
		err  error
	)
	if len(options) != 0 {
		resp, err = c.do(getMethod, "/global_rules", nil, options...)
	} else {
		resp, err = c.do(getMethod, "/global_rules", nil)
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &gris)
	if err != nil {
		return nil, err
	}

	for _, gr := range gris.List {
		grs = append(grs, gr.Value)
	}

	return grs, nil
}

// GetGlobalRule 查看指定的全局规则
func (c *Client) GetGlobalRule(id string) (*GlobalRule, error) {
	resp, err := c.do(getMethod, fmt.Sprintf("/global_rules/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var gri globalruleItem
	if err = json.Unmarshal(resp, &gri); err != nil {
		return nil, err
	} else {
		return gri.Value, nil
	}

}

// DeleteGlobalRule 删除指定的全局规则
func (c *Client) DeleteGlobalRule(id string) (*DeleteItemResp, error) {
	resp, err := c.do(deleteMethod, fmt.Sprintf("/global_rules/%s", id), nil)
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
