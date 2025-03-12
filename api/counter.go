package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sleep-go/lakala-pay/model"
	"io"
	"net/http"
)

const (
	specialCreateUrl = "/api/v3/ccss/counter/order/special_create"
	orderQueryUrl    = "/api/v3/ccss/counter/order/query"
	orderCloseUrl    = "/api/v3/ccss/counter/order/close"
	refundUrl        = "/api/v3/lams/trade/trade_refund"
	refundQueryUrl   = "/api/v3/lams/trade/trade_refund_query"
)

// OrderSpecialCreate 收银台订单创建
func (c *Client) OrderSpecialCreate(req *model.SpecialCreateReq) (*model.SpecialCreateRes, error) {
	return doRequest[model.SpecialCreateReq, model.SpecialCreateRes](c, specialCreateUrl, req)
}

// OrderQuery 收银台订单查询
func (c *Client) OrderQuery(req *model.OrderQueryReq) (*model.OrderQueryRes, error) {
	fmt.Println(req)
	return doRequest[model.OrderQueryReq, model.OrderQueryRes](c, orderQueryUrl, req)
}

// OrderClose 收银台订单关单
func (c *Client) OrderClose(req *model.OrderCloseReq) (resp *model.OrderCloseRes, err error) {
	return doRequest[model.OrderCloseReq, model.OrderCloseRes](c, orderCloseUrl, req)
}

// OrderNotifyCallback 收银台订单回调通知
func (c *Client) OrderNotifyCallback(r *http.Request) (resp *model.OrderNotify, err error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if !c.SignatureVerification(r.Header.Get("Authorization"), string(body)) {
		return nil, errors.New("签名验证失败")
	}
	err = json.Unmarshal(body, &resp)
	return resp, err
}

// OrderRefund 统一退货
func (c *Client) OrderRefund(req *model.RefundRequest) (resp *model.RefundResponse, err error) {
	return doRequest[model.RefundRequest, model.RefundResponse](c, refundUrl, req)
}

// RefundQuery 退货查询
func (c *Client) RefundQuery(req *model.RefundQueryRequest) (resp *model.RefundQueryResponse, err error) {
	return doRequest[model.RefundQueryRequest, model.RefundQueryResponse](c, refundQueryUrl, req)
}
