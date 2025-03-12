package api

import (
	"github.com/sleep-go/lakala-pay/model"
)

const (
	payUrl   = "/api/v3/ipsdf/paid/pay"
	queryUrl = "/api/v3/ipsdf/paid/query"
)

// 实时付款接口
func (c *Client) Pay(req *model.PayReq) (*model.PayRet, error) {
	return doRequest[model.PayReq, model.PayRet](c, payUrl, req)
}

// 实时付款查询接口
func (c *Client) PayQuery(req *model.PayQueryReq) (*model.PayQueryRet, error) {
	return doRequest[model.PayQueryReq, model.PayQueryRet](c, queryUrl, req)
}
