package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/sleep-go/lakala-pay/model"
)

// 分账通相关接口
const (
	uploadUrl               = "/api/v2/mms/openApi/uploadFile"
	cardBinUrl              = "/api/v2/mms/openApi/CardBin"
	applyUrl                = "/api/v2/mms/openApi/ledger/applyLedgerMer"
	modifyLedgerMerUrl      = "/api/v2/mms/openApi/ledger/ModifyLedgerMer"
	ledgerQueryUrl          = "/api/v2/mms/openApi/ledger/queryLedgerMer"
	applyLedgerReceiverUrl  = "/api/v2/mms/openApi/ledger/applyLedgerReceiver"
	modifyLedgerReceiverUrl = "/api/v2/mms/openApi/ledger/modifyLedgerReceiver"
	queryReceiverDetailUrl  = "/api/v2/mms/openApi/ledger/queryReceiverDetail"
	applyBindUrl            = "/api/v2/mms/openApi/ledger/applyBind"
	applyUnBindUrl          = "/api/v2/mms/openApi/ledger/applyUnBind"
	balanceQueryUrl         = "/api/v2/laep/industry/ewalletBalanceQuery"
	balanceSeparateUrl      = "/api/v3/sacs/balanceSeparate"
	balanceCancelUrl        = "/api/v3/sacs/balanceCancel"
	balanceFallbackUrl      = "/api/v3/sacs/balanceFallback"
	balanceSeparateQueryUrl = "/api/v3/sacs/balance_separate_query_plus"
)

// Upload 文件上传
func (c *Client) Upload(req *model.UploadReq) (*model.UploadRet, error) {
	return doRequest[model.UploadReq, model.UploadRet](c, uploadUrl, req)
}

// CardBin 卡BIN信息查询，仅支持对私结算账户
func (c *Client) CardBin(req *model.CardBinReq) (*model.CardBinRet, error) {
	return doRequest[model.CardBinReq, model.CardBinRet](c, cardBinUrl, req)
}

// LedgerApply 商户分账业务开通申请
func (c *Client) LedgerApply(req *model.ApplyReq) (*model.ApplyRet, error) {
	return doRequest[model.ApplyReq, model.ApplyRet](c, applyUrl, req)
}

// ModifyLedgerMer 商户分账信息变更申请
func (c *Client) ModifyLedgerMer(req *model.ModifyLedgerMerReq) (*model.ModifyLedgerMerRet, error) {
	return doRequest[model.ModifyLedgerMerReq, model.ModifyLedgerMerRet](c, modifyLedgerMerUrl, req)
}

// LedgerQuery 商户分账信息查询
func (c *Client) LedgerQuery(req *model.LedgerQueryReq) (*model.LedgerQueryRet, error) {
	return doRequest[model.LedgerQueryReq, model.LedgerQueryRet](c, ledgerQueryUrl, req)
}

// ApplyLedgerReceiver 分账接收方创建申请
func (c *Client) ApplyLedgerReceiver(req *model.ApplyLedgerReceiverReq) (*model.ApplyLedgerReceiverRet, error) {
	return doRequest[model.ApplyLedgerReceiverReq, model.ApplyLedgerReceiverRet](c, applyLedgerReceiverUrl, req)
}

// ModifyLedgerReceiver 分账接收方信息变更
func (c *Client) ModifyLedgerReceiver(req *model.ModifyLedgerReceiverReq) (*model.ModifyLedgerReceiverRet, error) {
	return doRequest[model.ModifyLedgerReceiverReq, model.ModifyLedgerReceiverRet](c, modifyLedgerReceiverUrl, req)
}

// QueryReceiverDetail 分账接收方详情查询
func (c *Client) QueryReceiverDetail(req *model.QueryReceiverDetailReq) (*model.QueryReceiverDetailRet, error) {
	return doRequest[model.QueryReceiverDetailReq, model.QueryReceiverDetailRet](c, queryReceiverDetailUrl, req)
}

// ApplyBind 分账关系绑定申请
func (c *Client) ApplyBind(req *model.ApplyBindReq) (*model.ApplyBindRet, error) {
	return doRequest[model.ApplyBindReq, model.ApplyBindRet](c, applyBindUrl, req)
}

// ApplyUnBind 分账关系解绑申请
func (c *Client) ApplyUnBind(req *model.ApplyUnBindReq) (*model.ApplyUnBindRet, error) {
	return doRequest[model.ApplyUnBindReq, model.ApplyUnBindRet](c, applyUnBindUrl, req)
}

// BalanceQuery 账户余额查询
func (c *Client) BalanceQuery(req *model.BalanceQueryReq) (*model.BalanceQueryRet, error) {
	return doRequest[model.BalanceQueryReq, model.BalanceQueryRet](c, balanceQueryUrl, req)
}

// BalanceSeparate 余额分账
func (c *Client) BalanceSeparate(req *model.BalanceSeparateReq) (*model.BalanceSeparateRet, error) {
	return doRequest[model.BalanceSeparateReq, model.BalanceSeparateRet](c, balanceSeparateUrl, req)
}

// BalanceCancel 余额分账撤销
func (c *Client) BalanceCancel(req *model.BalanceCancelReq) (*model.BalanceCancelRet, error) {
	return doRequest[model.BalanceCancelReq, model.BalanceCancelRet](c, balanceCancelUrl, req)
}

// BalanceFallback 余额分账回退
func (c *Client) BalanceFallback(req *model.BalanceFallbackReq) (*model.BalanceFallbackRet, error) {
	return doRequest[model.BalanceFallbackReq, model.BalanceFallbackRet](c, balanceFallbackUrl, req)
}

// BalanceSeparateQuery 分账结果查询
func (c *Client) BalanceSeparateQuery(req *model.BalanceSeparateQueryReq) (*model.BalanceSeparateQueryRet, error) {
	return doRequest[model.BalanceSeparateQueryReq, model.BalanceSeparateQueryRet](c, balanceSeparateQueryUrl, req)
}

type SeparateCallbackResp struct {
	Authorization string
	Body          string
	Notify        *model.SeparateNoticeReq
}

// NotifyCallback 分账结果通知
// https://o.lakala.com/#/home/document/detail?id=393
// SeparateNoticeReq,　SeparateNoticeRet 已定义
func (c *Client) NotifyCallback(r *http.Request) (*SeparateCallbackResp, error) {
	auth := r.Header.Get("Authorization")
	body, err := io.ReadAll(r.Body)
	var resp = &SeparateCallbackResp{
		Authorization: auth,
		Body:          string(body),
		Notify:        new(model.SeparateNoticeReq),
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
