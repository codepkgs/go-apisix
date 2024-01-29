package apisix

import (
	"fmt"
	"strings"
)

type Client struct {
	Address string
	Token   string
}

var ErrInvalidAddress = fmt.Errorf("apisix address must begin with http or https")

func NewClient(address, token string) (*Client, error) {
	var url string

	if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
		return nil, ErrInvalidAddress
	}

	if strings.HasSuffix(address, "/") {
		url = address + "apisix/admin"
	} else {
		url = address + "/apisix/admin"
	}

	return &Client{
		Address: url,
		Token:   token,
	}, nil
}
