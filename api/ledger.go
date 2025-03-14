package api

import (
	"github.com/sleep-go/lakala-pay/model"
)

const (
	uploadUrl               = "/api/v2/mms/openApi/uploadFile"
	cardBinUrl              = "/api/v2/mms/openApi/cardBin"
	applyUrl                = "/api/v2/mms/openApi/ledger/applyLedgerMer"
	modifyLedgerMerUrl      = "/api/v2/mms/openApi/ledger/modifyLedgerMer"
	ledgerQueryUrl          = "/api/v2/mms/openApi/ledger/queryLedgerMer"
	applyLedgerReceiverUrl  = "/api/v2/mms/openApi/ledger/applyLedgerReceiver"
	modifyLedgerReceiverUrl = "/api/v2/mms/openApi/ledger/modifyLedgerReceiver"
	queryReceiverDetailUrl  = "/api/v2/mms/openApi/ledger/queryReceiverDetail"
	applyBindUrl            = "/api/v2/mms/openApi/ledger/applyBind"
	applyUnBindUrl          = "/api/v2/mms/openApi/ledger/applyUnBind"
)

// 文件上传
func (c *Client) upload(req *model.UploadReq) (*model.UploadRet, error) {
	return doRequest[model.UploadReq, model.UploadRet](c, uploadUrl, req)
}

// 卡BIN信息查询，仅支持对私结算账户
func (c *Client) cardBin(req *model.CardBinReq) (*model.CardBinRet, error) {
	return doRequest[model.CardBinReq, model.CardBinRet](c, cardBinUrl, req)
}

// 商户分账业务开通申请
func (c *Client) ledgerApply(req *model.ApplyReq) (*model.ApplyRet, error) {
	return doRequest[model.ApplyReq, model.ApplyRet](c, applyUrl, req)
}

// 商户分账信息变更申请
func (c *Client) modifyLedgerMer(req *model.ModifyLedgerMerReq) (*model.ModifyLedgerMerRet, error) {
	return doRequest[model.ModifyLedgerMerReq, model.ModifyLedgerMerRet](c, modifyLedgerMerUrl, req)
}

// 商户分账信息查询
func (c *Client) ledgerQuery(req *model.LedgerQueryReq) (*model.LedgerQueryRet, error) {
	return doRequest[model.LedgerQueryReq, model.LedgerQueryRet](c, ledgerQueryUrl, req)
}

// 分账接收方创建申请
func (c *Client) applyLedgerReceiver(req *model.ApplyLedgerReceiverReq) (*model.ApplyLedgerReceiverRet, error) {
	return doRequest[model.ApplyLedgerReceiverReq, model.ApplyLedgerReceiverRet](c, applyLedgerReceiverUrl, req)
}

// 分账接收方信息变更
func (c *Client) modifyLedgerReceiver(req *model.ModifyLedgerReceiverReq) (*model.ModifyLedgerReceiverRet, error) {
	return doRequest[model.ModifyLedgerReceiverReq, model.ModifyLedgerReceiverRet](c, modifyLedgerReceiverUrl, req)
}

// 分账接收方详情查询
func (c *Client) queryReceiverDetail(req *model.QueryReceiverDetailReq) (*model.QueryReceiverDetailRet, error) {
	return doRequest[model.QueryReceiverDetailReq, model.QueryReceiverDetailRet](c, queryReceiverDetailUrl, req)
}

// 分账关系绑定申请
func (c *Client) applyBind(req *model.ApplyBindReq) (*model.ApplyBindRet, error) {
	return doRequest[model.ApplyBindReq, model.ApplyBindRet](c, applyBindUrl, req)
}

// 分账关系解绑申请
func (c *Client) applyUnBind(req *model.ApplyUnBindReq) (*model.ApplyUnBindRet, error) {
	return doRequest[model.ApplyUnBindReq, model.ApplyUnBindRet](c, applyUnBindUrl, req)
}
