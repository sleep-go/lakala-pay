package model

type RefundRequest struct {
	// 必填字段
	MerchantNo      string `json:"merchant_no" validate:"required,max=32"`            // 商户号(32)
	TermNo          string `json:"term_no" validate:"required,max=32"`                // 终端号(32)
	OutTradeNo      string `json:"out_trade_no" validate:"required,max=32"`           // 商户流水号(32)
	RefundAmount    string `json:"refund_amount" validate:"required,numeric"`         // 退款金额(单位分，整数)
	OriginBizType   string `json:"origin_biz_type" validate:"required,oneof=1 2 3 4"` // 原业务类型(1-4)原交易类型:1 银行卡，2 外卡，3 扫码，4 线上
	OriginTradeDate string `json:"origin_trade_date" validate:"required,len=8"`       // 原交易日期(yyyyMMdd)

	// 条件选填字段
	OriginLogNo   string `json:"origin_log_no,omitempty" validate:"omitempty,len=14"`   // 拉卡拉流水号(14)
	OriginTradeNo string `json:"origin_trade_no,omitempty" validate:"omitempty,max=32"` // 原交易订单号(32)
	OriginCardNo  string `json:"origin_card_no,omitempty" validate:"omitempty,max=32"`  // 原银行卡号(32)

	// 可选字段
	LocationInfo Location `json:"location_info,omitempty"`                                   // 地理位置信息
	RefundType   string   `json:"refund_type,omitempty" validate:"omitempty,oneof=00 05 06"` // 退款模式
}

// 地理位置子结构
type Location struct {
	Latitude  string `json:"latitude,omitempty"`  // 纬度
	Longitude string `json:"longitude,omitempty"` // 经度
	Address   string `json:"address,omitempty"`   // 地址信息
}

type RefundResponse struct {
	Msg        string     `json:"msg"`  // 返回信息描述（示例："执行成功"）
	Code       string     `json:"code"` // 返回码（示例："000000"）
	RefundData RefundData `json:"resp_data"`
}
type RefundData struct {
	// 必填字段
	LogNo     string `json:"log_no"`     // 拉卡拉退款单号（示例："66202310111234"）
	TradeTime string `json:"trade_time"` // 交易时间 yyyyMMddHHmmss（示例："20231001143025"）
	// 条件必填字段
	OriginTradeNo string `json:"origin_trade_no,omitempty"` // 原交易订单号（当请求中携带时返回）

	// 可选字段
	OutTradeNo       string `json:"out_trade_no,omitempty"`        // 商户请求流水号
	TotalAmount      string `json:"total_amount,omitempty"`        // 交易金额（单位分）
	PayerAmount      string `json:"payer_amount,omitempty"`        // 实际退款金额（单位分）
	RefundAmount     string `json:"refund_amount,omitempty"`       // 申请退款金额（单位分）
	AccTradeNo       string `json:"acc_trade_no,omitempty"`        // 账户端订单号
	OriginLogNo      string `json:"origin_log_no,omitempty"`       // 原统一交易单号
	OriginOutTradeNo string `json:"origin_out_trade_no,omitempty"` // 原商户请求流水号
}

// 退款请求数据结构
type RefundQueryRequest struct {
	// [必填] 商户编号，由拉卡拉分配
	// 示例值: "M100001234567890"
	// 最大长度: 32字符
	MerchantNo string `json:"merchant_no"`

	// [必填] 终端编号，由拉卡拉分配
	// 示例值: "T000012345678901234567890"
	// 最大长度: 32字符
	TermNo string `json:"term_no"`

	// [必填] 商户系统唯一流水号
	// 示例值: "REF202403261234567890ABCDEF"
	// 注意: 需保证商户系统内唯一性
	// 最大长度: 64字符
	OutTradeNo string `json:"out_trade_no"`

	// [必填] 原交易日期（POSP接口必填）
	// 格式: yyyyMMdd
	// 示例值: "20240326" 表示2024年3月26日
	OriginTradeDate string `json:"origin_trade_date"`

	// [必填] 原交易业务类型
	// 枚举值:
	// "1"-银行卡交易   "2"-外卡交易
	// "3"-扫码交易    "4"-线上交易
	// 示例值: "3"
	OriginBizType string `json:"origin_biz_type"`

	// [条件必填] 原交易订单标识（以下两组字段必须二选一）
	// --------------------------------------------------
	// 选项1: 拉卡拉系统订单号（小票订单号/三方订单号）
	// 示例值: "20240326123456789012345678901234"
	// 最大长度: 32字符
	OriginTradeNo string `json:"origin_trade_no,omitempty"`

	// 选项2: 系统参考号（POS小票参考号）
	// 示例值: "123456789012"
	// 严格长度: 12位数字
	OriginTradeRefNo string `json:"origin_trade_ref_no,omitempty"`
	// --------------------------------------------------

	// [可选] 原退款请求的商户流水号
	// 适用场景: 对历史退款交易进行再退款操作时使用
	// 示例值: "REF20230328123456"
	// 最大长度: 32字符
	OriginOutTradeNo string `json:"origin_out_trade_no,omitempty"`

	// [可选] 拉卡拉订单系统订单号
	// 适用场景: 订单系统生成的唯一标识
	// 示例值: "PO2024032612345678901234567890"
	// 最大长度: 32字符
	PayOrderNo string `json:"pay_order_no,omitempty"`
}

// 退款响应主体结构
type RefundQueryResponse struct {
	Msg        string            `json:"msg"`         // 返回信息描述（示例："执行成功"）
	Code       string            `json:"code"`        // 返回码（示例："000000"）
	RefundList []RefundQueryItem `json:"refund_list"` // [必填] 退款明细列表
}

// 单笔退款明细结构
type RefundQueryItem struct {
	// 基础信息
	OutTradeNo string `json:"out_trade_no,omitempty"` // 商户请求流水号（原退款交易）
	TradeTime  string `json:"trade_time,omitempty"`   // 交易时间 yyyyMMddHHmmss
	TradeState string `json:"trade_state"`            // [必填] 交易状态

	// 系统标识
	TradeNo string `json:"trade_no,omitempty"` // 拉卡拉订单号
	LogNo   string `json:"log_no"`             // [必填] 对账单流水号

	// 金额信息
	RefundAmount string `json:"refund_amount,omitempty"` // 交易金额（单位分）

	// 银行卡交易专有字段
	PayMode string `json:"pay_mode,omitempty"` // 支付方式(00:借记卡,01:贷记卡...)
	CrdNo   string `json:"crd_no,omitempty"`   // 脱敏卡号（格式：622888&zwnj;******&zwnj;1234）

	// 扫码交易专有字段
	AccountType string `json:"account_type,omitempty"` // 钱包类型(WECHAT/ALIPAY等)
	OpenId      string `json:"open_id,omitempty"`      // 用户主标识（商户体系）
	SubOpenId   string `json:"sub_open_id,omitempty"`  // 用户子标识（商户体系）

	// 金额明细（扫码交易）
	PayerAmount        string `json:"payer_amount,omitempty"`         // 实付金额（分）
	AccSettleAmount    string `json:"acc_settle_amount,omitempty"`    // 账户端应结金额（分）
	AccMdiscountAmount string `json:"acc_mdiscount_amount,omitempty"` // 商户优惠金额（分）
	AccDiscountAmount  string `json:"acc_discount_amount,omitempty"`  // 平台优惠金额（分）

	// 银行信息
	BankType string `json:"bank_type,omitempty"` // 付款银行标识
}
