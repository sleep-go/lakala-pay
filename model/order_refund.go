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
	Msg     string `json:"msg"`  // 返回信息描述（示例："执行成功"）
	Code    string `json:"code"` // 返回码（示例："000000"）
	ResData struct {
		RefundList []RefundQueryItem `json:"refund_list"`
	} `json:"resp_data"` // [必填] 退款明细列表
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

type RfdRefundRequest struct {
	// 商户号（必填）
	// 拉卡拉分配的商户唯一标识，固定长度15位
	MerchantNo string `json:"merchant_no" validate:"required,len=15" comment:"拉卡拉分配的商户号"`

	// 终端号（必填）
	// 拉卡拉分配的终端设备编号，固定长度8位
	TermNo string `json:"term_no" validate:"required,len=8" comment:"拉卡拉分配的终端号"`

	// 商户请求流水号（必填）
	// 商户系统内部的唯一交易标识，最大长度32位
	OutTradeNo string `json:"out_trade_no" validate:"required,max=32" comment:"商户系统唯一请求流水号"`

	// 退款金额（必填）
	// 单位为分，必须为整数数字型字符串，最大长度12位
	RefundAmount string `json:"refund_amount" validate:"required,numeric,max=12" comment:"退款金额（单位：分）"`

	// 退货原因（可选）
	// 记录退货的简要说明，最大长度32位
	RefundReason string `json:"refund_reason,omitempty" validate:"omitempty,max=32" comment:"退货原因说明"`

	// 拉卡拉对账单流水号（可选）
	// 交易返回的拉卡拉统一交易单号，扫码交易以66开头，POSP交易以年份后两位开头，固定长度14位
	OriginLogNo string `json:"origin_log_no,omitempty" validate:"omitempty,len=14" comment:"拉卡拉对账单流水号"`

	// 原商户交易流水号（可选）
	// 商户系统原始交易流水号，最大长度32位
	// 至少需要填写 origin_out_trade_no、origin_log_no、origin_trade_no 中的一个
	OriginOutTradeNo string `json:"origin_out_trade_no,omitempty" validate:"omitempty,max=32" comment:"原商户交易流水号"`

	// 原交易拉卡拉交易订单号（可选）
	// 拉卡拉系统原始交易订单号，最大长度32位
	// 优先级：origin_trade_no > origin_log_no > origin_out_trade_no
	OriginTradeNo string `json:"origin_trade_no,omitempty" validate:"omitempty,max=32" comment:"原交易拉卡拉订单号"`

	// 地址位置信息（必填）
	// 包含退货涉及的地址位置相关信息
	LocationInfo LocationInfo `json:"location_info" validate:"required" comment:"地址位置信息"` // 实际使用时需要替换为具体的结构体类型

	// 退货账户模式（可选）
	// 00: 退货账户余额
	// 05: 商户余额
	// 06: 终端余额
	RefundAccMode string `json:"refund_acc_mode,omitempty" validate:"omitempty,oneof=00 05 06" comment:"退货账户模式"`

	// 后台通知地址（可选）
	// 接收交易结果通知的回调URL，最大长度128位
	NotifyURL string `json:"notify_url,omitempty" validate:"omitempty,url,max=128" comment:"交易结果通知地址"`

	// 退货资金状态（可选）
	// 00: 分账前
	// 01: 分账后
	// 用于标识分账交易部分退货时的资金状态
	RefundAmtSts string `json:"refund_amt_sts,omitempty" validate:"omitempty,oneof=00 01" comment:"退货资金状态标识"`
}

type LocationInfo struct {
	// 请求IP地址（必填）
	// 客户端发起请求时的IP地址，用于定位请求来源
	RequestIP string `json:"request_ip" validate:"required,ip" comment:"客户端请求IP地址"`

	// 地理位置信息（可选）
	// 通过IP解析得到的地理位置描述（如：北京市朝阳区）
	// 为空时表示未获取到具体位置信息
	Location string `json:"location,omitempty" validate:"omitempty,max=128" comment:"地理位置描述信息"`
}

type RfdRefundResponse struct {
	Msg        string        `json:"msg"`  // 返回信息描述（示例："执行成功"）
	Code       string        `json:"code"` // 返回码（示例："000000"）
	RefundData RfdRefundData `json:"resp_data"`
}
type RfdRefundData struct {
	// 交易状态（必填）
	// 表示交易的当前状态，可能值：
	// INIT-初始化（查单确认最终结果）
	// SUCCESS-交易成功
	// FAIL-交易失败
	// DEAL-交易处理中/未知（查单确认最终结果）
	// PROCESSING-交易已受理（过段时间查单确认最终结果）
	// TIMEOUT-超时未知（查单确认最终结果）
	// EXCEPTION-异常（失败）
	TradeState string `json:"trade_state" validate:"required,oneof=INIT SUCCESS FAIL DEAL PROCESSING TIMEOUT EXCEPTION" comment:"交易状态标识"`

	// 退货模式（必填）
	// 标识退货的处理模式，固定长度2位
	RefundType string `json:"refund_type" validate:"required,len=2" comment:"退货处理模式"`

	// 商户号（必填）
	// 拉卡拉分配的商户唯一标识，最大长度20位
	MerchantNo string `json:"merchant_no" validate:"required,len=20" comment:"拉卡拉分配的商户号"`

	// 商户请求流水号（必填）
	// 请求报文中的商户系统唯一交易标识，最大长度32位
	OutTradeNo string `json:"out_trade_no" validate:"required,max=32" comment:"请求中的商户流水号"`

	// 拉卡拉交易流水号（必填）
	// 拉卡拉系统生成的交易唯一标识，最大长度32位
	TradeNo string `json:"trade_no" validate:"required,max=32" comment:"拉卡拉交易流水号"`

	// 拉卡拉对账单流水号（必填）
	// 对应交易流水号的后14位，固定长度14位
	LogNo string `json:"log_no" validate:"required,len=14" comment:"拉卡拉对账单流水号"`

	// 账户端交易订单号（可选）
	// 账户系统生成的交易流水号，最大长度32位
	AccTradeNo string `json:"acc_trade_no,omitempty" validate:"omitempty,max=32" comment:"账户端交易流水号"`

	// 钱包类型（可选）
	// 支付渠道标识：
	// 微信：WECHAT
	// 支付宝：ALIPAY
	// 银联：UQRCODEPAY
	// 翼支付：BESTPAY
	// 苏宁易付宝：SUNING
	AccountType string `json:"account_type,omitempty" validate:"omitempty,oneof=WECHAT ALIPAY UQRCODEPAY BESTPAY SUNING" comment:"支付渠道类型"`

	// 交易金额（必填）
	// 原始交易金额，单位为分，整数数字型字符串，最大长度12位
	TotalAmount string `json:"total_amount" validate:"required,numeric,max=12" comment:"原始交易金额（单位：分）"`

	// 申请退款金额（必填）
	// 请求报文中的退款金额，单位为分，整数数字型字符串，最大长度12位
	RefundAmount string `json:"refund_amount" validate:"required,numeric,max=12" comment:"申请退款金额（单位：分）"`

	// 实际退款金额（必填）
	// 实际处理完成的退款金额，单位为分，整数数字型字符串，最大长度12位
	PayerAmount string `json:"payer_amount" validate:"required,numeric,max=12" comment:"实际退款金额（单位：分）"`

	// 退款时间（可选）
	// 实际完成退款的时间，格式为yyyyMMddHHmmss，固定长度14位
	TradeTime string `json:"trade_time,omitempty" validate:"omitempty,len=14,datetime=yyyyMMddHHmmss" comment:"实际退款时间"`

	// 原拉卡拉对账单流水号（可选）
	// 如果请求中携带该字段，响应中原样返回，固定长度14位
	OriginLogNo string `json:"origin_log_no,omitempty" validate:"omitempty,len=14" comment:"原拉卡拉对账单流水号"`

	// 原拉卡拉交易流水号（可选）
	// 如果请求中携带该字段，响应中原样返回，最大长度32位
	OriginTradeNo string `json:"origin_trade_no,omitempty" validate:"omitempty,max=32" comment:"原拉卡拉交易流水号"`

	// 原商户请求流水号（可选）
	// 如果请求中携带该字段，响应中原样返回，最大长度32位
	OriginOutTradeNo string `json:"origin_out_trade_no,omitempty" validate:"omitempty,max=32" comment:"原商户请求流水号"`

	// 单品营销附加数据（可选）
	// 扫码交易参与单品营销优惠时返回，最大长度8000位
	UpIssAddData string `json:"up_iss_add_data,omitempty" validate:"omitempty,max=8000" comment:"单品营销附加数据"`

	// 银联优惠信息（可选）
	// 扫码交易参与银联优惠时返回，最大长度500位
	// 示例格式：[{"fundchannel":"BOC","amount":"18"}]
	UpCouponInfo string `json:"up_coupon_info,omitempty" validate:"omitempty,max=500" comment:"银联优惠信息"`

	// 出资方信息（可选）
	// 扫码交易返回的资金方信息，最大长度512位
	// 示例格式：[{"fundchannel":"BOC","amount":"18"}]
	TradeInfo string `json:"trade_info,omitempty" validate:"omitempty,max=512" comment:"资金方出资信息"`

	// 返回描述信息（必填）
	// 格式为code#msg的组合，例如：
	// RFD00000#成功
	// RFD11112#网络请求超时
	ChannelRetDesc string `json:"channel_ret_desc" validate:"required,regex=^[A-Z0-9]{7}#.*$" comment:"接口返回状态描述"`
}

type RfdRefundQueryRequest struct {
	// 商户号（必填）
	// 拉卡拉分配的商户唯一标识，固定长度15位
	MerchantNo string `json:"merchant_no" validate:"required,len=15" comment:"拉卡拉分配的商户号"`

	// 终端号（必填）
	// 拉卡拉分配的业务终端标识，固定长度8位
	TermNo string `json:"term_no" validate:"required,len=8" comment:"拉卡拉分配的业务终端号"`

	// 商户交易流水号（可选）
	// 商户系统生成的原始交易流水号，最大长度32位
	// 用于关联原始交易请求
	OutTradeNo string `json:"out_trade_no,omitempty" validate:"omitempty,max=32" comment:"下单时的商户请求流水号"`

	// 拉卡拉交易流水号（可选）
	// 拉卡拉系统生成的交易唯一标识，最大长度32位
	// 优先使用此字段进行交易查询
	TradeNo string `json:"trade_no,omitempty" validate:"omitempty,max=32" comment:"拉卡拉交易流水号"`
}

// 退款响应主体结构
type RfdRefundQueryResponse struct {
	Msg     string             `json:"msg"`       // 返回信息描述（示例："执行成功"）
	Code    string             `json:"code"`      // 返回码（示例："000000"）
	ResData RfdRefundQueryItem `json:"resp_data"` // [必填] 退款明细列表
}

type RfdRefundQueryItem struct {
	// 商户请求流水号（必填）
	// 对应原始退款请求的商户流水号（扫码交易必返）
	OutTradeNo string `json:"out_trade_no" validate:"required" comment:"原退款交易商户请求流水号（扫码交易返回）"`

	// 交易时间（必填）
	// 格式：yyyyMMddHHmmss
	TradeTime string `json:"trade_time" validate:"required,datetime=yyyyMMddHHmmss" comment:"交易时间"`

	// 交易状态（必填）
	// 交易处理结果标识：
	// INIT-初始化；SUCCESS-成功；FAIL-失败；DEAL-处理中；TIMEOUT-超时；EXCEPTION-异常
	TradeState string `json:"trade_state" validate:"required,oneof=INIT SUCCESS FAIL DEAL TIMEOUT EXCEPTION" comment:"交易状态"`

	// 拉卡拉商户订单号（必填）
	// 拉卡拉系统生成的交易唯一标识
	TradeNo string `json:"trade_no" validate:"required" comment:"拉卡拉生成的交易流水"`

	// 拉卡拉对账单流水号（必填）
	// 用于财务对账的唯一标识（新增字段）
	LogNo string `json:"log_no" validate:"required" comment:"拉卡拉生成的对账单流水号（新增）"`

	// 账户端交易订单号（可选）
	// 账户系统生成的交易标识（如微信支付订单号）
	AccTradeNo string `json:"acc_trade_no,omitempty" validate:"omitempty,max=32" comment:"账户端交易订单号"`

	// 交易金额（必填）
	// 单位：分（如100表示1元）
	RefundAmount string `json:"refund_amount" validate:"required,numeric" comment:"交易金额"`

	// 支付方式（可选）
	// 银行卡交易返回：
	// 00-借记卡；01-贷记卡；02-准贷记卡
	PayMode string `json:"pay_mode,omitempty" validate:"omitempty,oneof=00 01 02" comment:"支付方式"`

	// 脱敏卡号（可选）
	// 格式：前6位 + * + 后4位（如622202&zwnj;******&zwnj;1234）
	CrdNo string `json:"crd_no,omitempty" validate:"omitempty,max=16,regexp=^\\d{6}\\*{4,}\\d{4}$" comment:"脱敏卡号"`

	// 钱包类型（可选）
	// 扫码交易返回：
	// 微信：WECHAT；支付宝：ALIPAY；银联：UNION；翼支付：BESTPAY；苏宁易付宝：SUNING
	AccountType string `json:"account_type,omitempty" validate:"omitempty,oneof=WECHAT ALIPAY UNION BESTPAY SUNING" comment:"钱包类型"`

	// 付款人实付金额（可选）
	// 实际退款金额，单位：分
	PayerAmount string `json:"payer_amount,omitempty" validate:"omitempty,numeric" comment:"付款人实付金额"`

	// 账户端结算金额（可选）
	// 账户端应结订单金额，单位：分
	AccSettleAmount string `json:"acc_settle_amount,omitempty" validate:"omitempty,numeric" comment:"账户端结算金额"`

	// 商户侧优惠金额（可选）
	// 商户提供的优惠金额，单位：分
	AccMdiscountAmount string `json:"acc_mdiscount_amount,omitempty" validate:"omitempty,numeric" comment:"商户侧优惠金额（账户端）"`

	// 账户端优惠金额（可选）
	// 拉卡拉提供的优惠金额，单位：分
	AccDiscountAmount string `json:"acc_discount_amount,omitempty" validate:"omitempty,numeric" comment:"账户端优惠金额"`

	// 返回描述信息（必填）
	// 格式：code#msg（如RFD00000#成功）
	ChannelRetDesc string `json:"channel_ret_desc" validate:"required,regexp=^[A-Z0-9]{5}#[^#]{1,64}$" comment:"返回描述信息"`
}
