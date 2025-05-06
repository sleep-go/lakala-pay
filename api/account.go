package api

import "github.com/sleep-go/lakala-pay/model"

const (
	fundAcctBalanceQueryUrl = "/api/v3/svas_front/open/bank/other/fund_acct_balance_query"
)

// FundAcctBalanceQuery 账户管理 > 银行虚拟户分账 > 账户通知及查询 > 实账户余额查询
func (c *Client) FundAcctBalanceQuery(req *model.FundAcctBalanceQueryReq) (*model.FundAcctBalanceQueryRet, error) {
	return doRequest[model.FundAcctBalanceQueryReq, model.FundAcctBalanceQueryRet](c, fundAcctBalanceQueryUrl, req)
}
