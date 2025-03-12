package api

import (
	"encoding/json"
	"errors"
	"github.com/sleep-go/lakala-pay/model"
)

const (
	specialCreateUrl = "/api/v3/ccss/counter/order/special_create"
	orderQueryUrl    = "/api/v3/ccss/counter/order/query"
	orderCloseUrl    = "/api/v3/ccss/counter/order/close"
	refundUrl        = "/api/v3/labs/relation/refund"
)

// OrderSpecialCreate 收银台订单创建
func (c *Client) OrderSpecialCreate(req *model.SpecialCreateReq) (*model.SpecialCreateRes, error) {
	return doRequest[model.SpecialCreateReq, model.SpecialCreateRes](c, specialCreateUrl, req, false)
}

// OrderQuery 收银台订单查询
func (c *Client) OrderQuery(req *model.OrderQueryReq) (*model.OrderQueryRes, error) {
	return doRequest[model.OrderQueryReq, model.OrderQueryRes](c, orderQueryUrl, req, false)
}

// OrderClose 收银台订单关单
func (c *Client) OrderClose(req *model.OrderCloseReq) (resp *model.OrderCloseRes, err error) {
	return doRequest[model.OrderCloseReq, model.OrderCloseRes](c, orderCloseUrl, req, false)
}

// OrderNotifyCallback 收银台订单回调通知
func (c *Client) OrderNotifyCallback(authorization, body string) (resp *model.OrderNotify, err error) {
	if !c.SignatureVerification(authorization, body) {
		return nil, errors.New("签名验证失败")
	}
	err = json.Unmarshal([]byte(body), &resp)
	return resp, err
}
