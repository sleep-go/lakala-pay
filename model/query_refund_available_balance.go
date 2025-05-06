package model

type QueryRefundAvailableBalanceReq struct {
	Ver       string `json:"ver"`
	Timestamp string `json:"timestamp"`
	ReqId     string `json:"reqId"`
	ReqData   struct {
		MerId       string `json:"merId"`
		TermId      string `json:"termId"`
		QryNo       string `json:"qryNo"`
		BalanceFlag string `json:"balanceFlag"`
		OlogNo      string `json:"ologNo"`
		SourceFlag  string `json:"sourceFlag"`
	} `json:"reqData"`
}
type QueryRefundAvailableBalanceRet struct {
	RetCode  string `json:"retCode"`
	RetMsg   string `json:"retMsg"`
	RespData struct {
		QryNo       string `json:"qryNo"`
		MerId       string `json:"merId"`
		TermId      string `json:"termId"`
		Balance     string `json:"balance"`
		BalanceFlag string `json:"balanceFlag"`
		AccNo       string `json:"accNo"`
		AccName     string `json:"accName"`
	} `json:"respData"`
}
