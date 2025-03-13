package api

import (
	"github.com/sleep-go/lakala-pay/model"
)

const (
	uploadUrl              = "/api/v2/mms/openApi/uploadFile"
	applyUrl               = "/api/v2/mms/openApi/ledger/applyLedgerMer"
	ledgerQueryUrl         = "/api/v2/mms/openApi/ledger/queryLedgerMer"
	applyLedgerReceiverUrl = "/api/v2/mms/openApi/ledger/applyLedgerReceiver"
)

// 文件上传
func (c *Client) upload(req *model.UploadReq) (*model.UploadRet, error) {
	return doRequest[model.UploadReq, model.UploadRet](c, uploadUrl, req)
}

// 商户分账业务开通申请
func (c *Client) ledgerApply(req *model.ApplyReq) (*model.ApplyRet, error) {
	return doRequest[model.ApplyReq, model.ApplyRet](c, applyUrl, req)
}

// 商户分账信息查询
func (c *Client) ledgerQuery(req *model.LedgerQueryReq) (*model.LedgerQueryRet, error) {
	return doRequest[model.LedgerQueryReq, model.LedgerQueryRet](c, ledgerQueryUrl, req)
}

// 分账接收方创建申请
func (c *Client) applyLedgerReceiver(req *model.ApplyLedgerReceiverReq) (*model.ApplyLedgerReceiverRet, error) {
	return doRequest[model.ApplyLedgerReceiverReq, model.ApplyLedgerReceiverRet](c, applyLedgerReceiverUrl, req)
}
