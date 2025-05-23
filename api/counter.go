package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/sleep-go/lakala-pay/model"
)

// 拉卡拉收银台相关接口
const (
	specialCreateUrl               = "/api/v3/ccss/counter/order/special_create"
	orderQueryUrl                  = "/api/v3/ccss/counter/order/query"
	orderCloseUrl                  = "/api/v3/ccss/counter/order/close"
	refundUrl                      = "/api/v3/lams/trade/trade_refund"
	refundQueryUrl                 = "/api/v3/lams/trade/trade_refund_query"
	rfdRefundUrl                   = "/api/v3/rfd/refund_front/refund"
	rfdRefundQueryUrl              = "/api/v3/rfd/refund_front/refund_query"
	queryRefundAvailableBalanceUrl = "/api/v2/mrssQuery/queryRefundAvailableBalance" //退货可用余额查询
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

type CallbackResp struct {
	Authorization string
	Body          string
	Notify        *model.OrderNotify
}

// OrderNotifyCallback 收银台订单回调通知
func (c *Client) OrderNotifyCallback(r *http.Request) (*CallbackResp, error) {
	auth := r.Header.Get("Authorization")
	body, err := io.ReadAll(r.Body)
	var resp = &CallbackResp{
		Authorization: auth,
		Body:          string(body),
		Notify:        new(model.OrderNotify),
	}
	if err != nil {
		return resp, err
	}
	if !c.SignatureVerification(auth, string(body)) {
		return resp, errors.New("签名验证失败")
	}
	err = json.Unmarshal(body, &resp.Notify)
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

// OrderRfdRefund 扫码银行卡退货
func (c *Client) OrderRfdRefund(req *model.RfdRefundRequest) (resp *model.RfdRefundResponse, err error) {
	return doRequest[model.RfdRefundRequest, model.RfdRefundResponse](c, rfdRefundUrl, req)
}

// OrderRfdRefundQuery 扫码银行卡退货 退货查询
func (c *Client) OrderRfdRefundQuery(req *model.RfdRefundQueryRequest) (resp *model.RfdRefundQueryResponse, err error) {
	return doRequest[model.RfdRefundQueryRequest, model.RfdRefundQueryResponse](c, rfdRefundQueryUrl, req)
}

// QueryRefundAvailableBalance 退货可用余额查询
func (c *Client) QueryRefundAvailableBalance(req *model.QueryRefundAvailableBalanceReq) (resp *model.QueryRefundAvailableBalanceRet, err error) {
	return doRequest[model.QueryRefundAvailableBalanceReq, model.QueryRefundAvailableBalanceRet](c, queryRefundAvailableBalanceUrl, req)
}
