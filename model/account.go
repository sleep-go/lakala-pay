package model

type FundAcctBalanceQueryReq struct {
	OwnerId string `json:"owner_id"` //接入机构号
	CustId  string `json:"cust_id"`  //客户ID
	AcctNbr string `json:"acct_nbr"` //实账户
}
type FundAcctBalanceQueryRet struct {
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	RespTime string `json:"resp_time"`
	RespData struct {
		DealAcctNo   string `json:"deal_acct_no"`  //交易资金账户
		MerchantName string `json:"merchant_name"` //商户名称
		PreAmt       string `json:"pre_amt"`       //上一日余额
		CurAmt       string `json:"cur_amt"`       //当前余额
	} `json:"resp_data"`
}
