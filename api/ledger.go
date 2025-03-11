package api

import (
	"github.com/sleep-go/lakala-pay/model"
)

const (
	uploadUrl = "/api/v2/mms/openApi/uploadFile"
	applyUrl  = "/api/v2/mms/openApi/ledger/applyLedgerMer"
	//queryUrl = "/api/v3/ipsdf/paid/query"
)

// 商户分账业务开通申请
func (c *Client) apply(req *model.ApplyReq) (*model.ApplyRet, error) {
	return doRequest[model.ApplyReq, model.ApplyRet](c, applyUrl, req, false)
}

func (c *Client) upload(req *model.UploadReq) (*model.UploadRet, error) {
	return doRequest[model.UploadReq, model.UploadRet](c, uploadUrl, req, false)
}
