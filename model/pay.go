package model

type PayReq struct {
	OutOrderNo string    `json:"out_trade_no" validate:"required,max=32"`       // 商户订单号，必填且长度不超过32
	MerchantNo string    `json:"merchant_no" validate:"required,max=32"`        // 银联商户号，必填且长度不超过32
	TermNo     string    `json:"term_no"`                                       //拉卡拉分配的业务终端号
	VposID     string    `json:"vpos_id,omitempty" validate:"omitempty,max=32"` // 交易设备标识，可选且长度不超过32
	PeriodType string    `json:"period_type"  validate:"required,max=2"`        //D0 - 实时 暂时只支持：D0
	TxnAmt     int       `json:"txn_amt"       validate:"required,max=12"`      //单位分，整数型字符
	PayerInfo  PayerInfo `json:"payer_info"`
	Remark     string    `json:"remark"       validate:"max=128"` //附言
	Summary    string    `json:"summary"       validate:"max=64"` //业务摘要
}

type PayerInfo struct {
	IdentityType   string `json:"identity_type"`    //参与方类型, OutBankCard:外部银行 InAccNo:内部账号(拉卡拉内部账号)
	BankNo         string `json:"bank_no"`          //开户行行号,OutBankCard:外部银行卡 必选
	BankName       string `json:"bank_name"`        //开户行名称,OutBankCard:外部银行卡 必选
	AcctNo         string `json:"acct_no"`          //账号,OutBankCard:外部银行卡 InAccNo:内部账号
	AcctName       string `json:"acct_name"`        //账户名称,OutBankCard:外部银行卡户InAccNo:内部账号对应商户名称
	AcctType       string `json:"acct_type"`        //收款账户类型。 0：对公（在金融机构开设的公司账户） 1：对私（在金融机构开设的个人账户）
	BankBranchNo   string `json:"bank_branch_no"`   //联行号,如果银行卡为对公，必传
	BankBranchName string `json:"bank_branch_name"` //支行名称, 如果银行卡为对公，必传
}

type PayRet struct {
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	RespTime string `json:"resp_time"`
	RespData struct {
		AccountTradeNo string `json:"account_trade_no"`
		EndTime        string `json:"end_time"`
		FeeAmount      string `json:"fee_amount"`
		MerchantNo     string `json:"merchant_no"`
		OutTradeNo     string `json:"out_trade_no"`
		PeriodType     string `json:"period_type"`
		TermNo         string `json:"term_no"`
		TradeNo        string `json:"trade_no"`
		TxnAmt         string `json:"txn_amt"`
	} `json:"resp_data"`
}

type PayQueryReq struct {
	OutOrderNo string `json:"out_trade_no,omitempty" validate:"omitempty,max=32"` // 商户订单号
	MerchantNo string `json:"merchant_no,omitempty" validate:"omitempty,max=32"`  // 银联商户号
	TradeNo    string `json:"trade_no,omitempty" validate:"omitempty,max=64"`     // 交易号
	TermNo     string `json:"term_no,omitempty" validate:"omitempty,max=32"`      // 终端号
}

type PayQueryRet struct {
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	RespTime string `json:"resp_time"`
	RespData struct {
		AccountTradeNo string `json:"account_trade_no"` //账户端交易流水号
		EndTime        string `json:"end_time"`         //完成时间
		FeeAmount      string `json:"fee_amount"`       //手续费,单位分，整数数字型字符
		MerchantNo     string `json:"merchant_no"`      //拉卡拉分配的商户号
		OutOrderNo     string `json:"out_trade_no"`     //请求报文中的商户请求流水号
		PeriodType     string `json:"period_type"`      //结算/出款周期
		ReturnDesc     string `json:"return_desc"`      //结果描述
		TermNo         string `json:"term_no"`          //拉卡拉分配的业务终端号
		TradeNo        string `json:"trade_no"`         //拉卡拉交易流水号
		TradeStatus    string `json:"trade_status"`     //INIT-初始化 CREATE-下单成功 SUCCESS-交易成功 FAIL-交易失败 DEAL-交易处理中 UNKNOWN-未知状态 PART_REFUND-部分退款 REFUND-全部退款
		TxnAmt         string `json:"txn_amt"`          //订单金额, 单位分，整数数字型字符
	} `json:"resp_data"`
}
