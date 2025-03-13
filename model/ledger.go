package model

type UploadReq struct {
	Ver     string        `json:"ver"`
	ReqTime string        `json:"timestamp"`
	ReqId   string        `json:"reqId"`
	ReqData UploadReqData `json:"reqData"`
}

type UploadReqData struct {
	Version    string `json:"version" validate:"required,max=8"`     // 接口版本号，固定8位字符（示例："1.0"）
	OrderNo    string `json:"orderNo" validate:"required,max=32"`    // 订单跟踪编号，格式：14位年月日时分秒(YYYYMMDDHHMMSS)+8位随机数（示例：2024031109300012345678）
	OrgCode    string `json:"orgCode" validate:"required,max=32"`    // 合作机构唯一标识码，需向平台申请获取（示例：LKL_ORG_2024）
	AttType    string `json:"attType" validate:"required,max=32"`    // 证件附件类型枚举值：ID_CARD_FRONT-身份证正面，ID_CARD_BACK-身份证反面
	AttExtName string `json:"attExtName" validate:"required,max=32"` // 允许上传的扩展名类型：jpg/png/pdf，单文件大小不超过5MB
	AttContext string `json:"attContext" validate:"required"`        // BASE64编码文件内容，前端需用Base64Utils.encodeToString()编码，后端使用Spring Base64Utils.decodeFromString()解码
}

type UploadRet struct {
	OrgCode   string `json:"orgCode" validate:"required"`   // 合作机构唯一标识码，需向平台申请获取（示例：ORG_2024）
	OrderNo   string `json:"orderNo" validate:"required"`   // 业务订单编号，格式：14位年月日时分秒+8位随机数（示例：202403110930001234）
	AttFileId string `json:"attFileId" validate:"required"` // 附件唯一标识符，由文件服务生成（示例：ATT_20240311_123456）‌
	AttType   string `json:"attType" validate:"required"`   // 附件类型枚举值：ID_CARD_FRONT-身份证正面，INVOICE-发票，CONTRACT-合同‌
}

type ApplyReq struct {
	Ver     string       `json:"version"`
	ReqTime string       `json:"reqTime"`
	ReqId   string       `json:"reqId"`
	ReqData ApplyReqData `json:"reqData"`
}

type ApplyReqData struct {
	Version              string  `json:"version" validate:"required"`              // 接口版本号1.0（固定8字符长度）
	OrderNo              string  `json:"orderNo" validate:"required"`              // 订单编号（32字符，格式：14位年月日时分秒+8位随机数）
	OrgCode              string  `json:"orgCode" validate:"required"`              // 机构代码（32字符，唯一标识合作机构）
	MerInnerNo           string  `json:"merInnerNo,omitempty"`                     // 拉卡拉内部商户号（与银联商户号二选一，优先使用本字段）
	MerCupNo             string  `json:"merCupNo,omitempty"`                       // 银联商户号（与拉卡拉内部商户号二选一）
	ContactMobile        string  `json:"contactMobile" validate:"required"`        // 联系手机号（32字符，用于业务通知）
	SplitLowestRatio     float64 `json:"splitLowestRatio" validate:"required"`     // 最低分账比例（百分比，支持两位小数，如70.50）
	SplitEntrustFileName string  `json:"splitEntrustFileName" validate:"required"` // 分账委托书文件名（64字符，示例：分账委托书_V1.0.pdf）
	SplitEntrustFilePath string  `json:"splitEntrustFilePath" validate:"required"` // 分账委托书文件路径（64字符，通过文件上传接口获取）
	SplitRange           string  `json:"splitRange,omitempty"`                     // 分账范围（ALL-全部分账，MARK-标记分账，默认MARK）
	SepFundSource        string  `json:"sepFundSource,omitempty"`                  // 分账依据（TR-交易分账，BA-余额分账，默认TR）
	EleContractNo        string  `json:"eleContractNo,omitempty"`                  // 电子合同编号（32字符，已签署电子合同需填写）
	SplitLaunchMode      string  `json:"splitLaunchMode,omitempty"`                // 分账发起方式（AUTO-自动，POINTRULE-指定规则，MANUAL-手动，默认MANUAL）
	SettleType           string  `json:"settleType,omitempty"`                     // 提款类型（01-主动提款，03-自动结算，默认01）
	SplitRuleSource      string  `json:"splitRuleSource,omitempty"`                // 分账规则来源（MER-商户，PLATFORM-平台，自动/指定分账时必填）
	RetUrl               string  `json:"retUrl" validate:"required"`               // 回调通知地址（128字符，需支持HTTPS协议）
}

type ApplyRet struct {
	Code     string `json:"retCode"`
	Msg      string `json:"retMsg"`
	RespData struct {
		Version string `json:"version"` // 接口版本号（示例值：547110502170558464）‌
		OrderNo string `json:"orderNo"` // 订单编号（示例值：2021020112000012345678）‌
		OrgCode string `json:"orgCode"` // 机构代码（示例值：200669）‌
		ApplyId int64  `json:"applyId"` // 受理编号（长整型数字，示例：548099616395186176）‌
	} `json:"respData"`
}

type LedgerQueryData struct {
	Version    string `json:"version"`              // 接口版本号 1.0
	OrderNo    string `json:"orderNo"`              // 订单编号，14位年月日时（24小时制）分秒+8位的随机数（不重复）如：2021020112000012345678
	OrgCode    string `json:"orgCode"`              // 机构代码
	MerInnerNo string `json:"merInnerNo,omitempty"` // 拉卡拉内部商户号，拉卡拉内部商户号和银联商户号必须传一个，都送以内部商户号为准
	MerCupNo   string `json:"merCupNo,omitempty"`   // 银联商户号，
}

type LedgerQueryReq struct {
	Ver     string          `json:"version"`
	ReqTime string          `json:"reqTime"`
	ReqId   string          `json:"reqId"`
	ReqData LedgerQueryData `json:"reqData"`
}

type LedgerQueryRet struct {
	Code     string             `json:"retCode"`
	Msg      string             `json:"retMsg"`
	RespData LedgerQueryRetData `json:"respData"`
}

type LedgerQueryRetData struct {
	// 分账商户机构号
	OrgId string `json:"orgId"`
	// 分账商户机构名称
	OrgName string `json:"orgName"`
	// 拉卡拉内部商户号
	MerInnerNo string `json:"merInnerNo"`
	// 银联商户号
	MerCupNo string `json:"merCupNo"`
	// 最低分账比例（百分比，支持2位精度），例如 "70" 或 "70.50"
	SplitLowestRatio string `json:"splitLowestRatio"`
	// 商户分账状态，VALID 表示启用，INVALID 表示禁用
	SplitStatus string `json:"splitStatus"`
	// 分账范围，ALL 表示全部交易分账（商户所有交易默认待分账），MARK 表示标记交易分账（只有带分账标识交易待分账，其余交易正常结算），默认值为 MARK
	SplitRange string `json:"splitRange"`
	// 分账依据，TR 或空表示交易分账，BA 表示余额分账，默认值为 TR（交易分账）
	SepFundSource string `json:"sepFundSource"`
	// 平台ID，如果商户和绑定平台分账，则返回平台ID
	PlatformId string `json:"platformId"`
	// 分账发起方式，AUTO 表示自动规则分账，POINTRULE 表示指定规则分账，MANUAL 表示手动规则分账
	SplitLaunchMode string `json:"splitLaunchMode"`
	// 分账规则来源，MER 表示商户分账规则，PLATFORM 表示平台分账规则
	SplitRuleSource string `json:"splitRuleSource"`
	// 已绑定接收方列表，是一个集合类型，具体类型需要根据实际情况定义（例如：[]string, []Receiver 等）
	// 这里为了简化，使用 interface{} 类型表示，实际使用时需要替换为具体类型
	BindRelations []BindRelation `json:"bindRelations"`
}

type BindRelation struct {
	// 拉卡拉内部商户号
	MerInnerNo string `json:"merInnerNo"`
	// 银联商户号，
	MerCupNo string `json:"merCupNo"`
	// 接收方编号，用于唯一标识一个接收方
	ReceiverNo string `json:"receiverNo"`
	// 接收方名称，如店铺名称等
	ReceiverName string `json:"receiverName"`
}

type ApplyLedgerReceiverReq struct {
	Ver     string                  `json:"version"`
	ReqTime string                  `json:"reqTime"`
	ReqId   string                  `json:"reqId"`
	ReqData ApplyLedgerReceiverData `json:"reqData"`
}

type ApplyLedgerReceiverData struct {
	// 接口版本号
	Version string `json:"version"`
	// 订单编号（14位年月日时分秒+8位随机数）
	OrderNo string `json:"orderNo"`
	// 机构代码
	OrgCode string `json:"orgCode"`
	// 分账接收方名称
	ReceiverName string `json:"receiverName"`
	// 联系手机号
	ContactMobile string `json:"contactMobile"`
	// 营业执照号码（对公账户必填）
	LicenseNo string `json:"licenseNo,omitempty"`
	// 营业执照名称（对公账户必填）
	LicenseName string `json:"licenseName,omitempty"`
	// 法人姓名（对公账户必填）
	LegalPersonName string `json:"legalPersonName,omitempty"`
	// 法人证件类型（对公账户必填，17:身份证, 18:护照, 19:港澳通行证, 20:台湾通行证）
	LegalPersonCertificateType string `json:"legalPersonCertificateType,omitempty"`
	// 法人证件号（对公账户必填）
	LegalPersonCertificateNo string `json:"legalPersonCertificateNo,omitempty"`
	// 收款账户卡号
	AcctNo string `json:"acctNo"`
	// 收款账户名称
	AcctName string `json:"acctName"`
	// 收款账户类型（57:对公, 58:对私）
	AcctTypeCode string `json:"acctTypeCode"`
	// 收款账户证件类型（17:身份证, 18:护照, 19:港澳通行证, 20:台湾通行证）
	AcctCertificateType string `json:"acctCertificateType"`
	// 收款账户证件号
	AcctCertificateNo string `json:"acctCertificateNo"`
	// 收款账户开户行号（仅支持对私结算账户）
	AcctOpenBankCode string `json:"acctOpenBankCode"`
	// 收款账户开户名称
	AcctOpenBankName string `json:"acctOpenBankName"`
	// 收款账户清算行行号（仅支持对私结算账户）
	AcctClearBankCode string `json:"acctClearBankCode"`
	// 接收方附件资料列表
	AttachList []Attach `json:"attachList,omitempty"`
	// 提款类型（01:主动提款, 03:交易自动结算, 默认01）
	SettleType string `json:"settleType,omitempty"`
}

// https://o.lakala.com/#/home/document/detail?id=382
type Attach struct {
	// 附件类型编码
	AttachType string `json:"attachType"`
	// 附件名称
	AttachName string `json:"attachName"`
	// 附件路径（调用进件附件上传接口获取）
	AttachStorePath string `json:"attachStorePath"`
}

type ApplyLedgerReceiverRet struct {
	Code     string `json:"retCode"`
	Msg      string `json:"retMsg"`
	RespData struct {
		// 接口版本号（回传）
		Version string `json:"version"`
		// 订单编号（回传）
		OrderNo string `json:"orderNo"`
		// 申请机构代码（回传）
		OrgCode string `json:"orgCode"`
		// 接收方所属机构ID
		OrgId string `json:"orgId"`
		// 接收方所属机构名称
		OrgName string `json:"orgName"`
		// 接收方编号
		ReceiverNo string `json:"receiverNo"`
	} `json:"respData"`
}

type QueryReceiverDetailReqData struct {
	// 接口版本号
	Version string `json:"version"`
	// 订单编号（14位年月日时分秒+8位随机数）
	OrderNo string `json:"orderNo"`
	// 机构代码
	OrgCode string `json:"orgCode"`
	// 接收方编号
	ReceiverNo string `json:"receiverNo"`
}

type QueryReceiverDetailReq struct {
	Ver     string                     `json:"version"`
	ReqTime string                     `json:"reqTime"`
	ReqId   string                     `json:"reqId"`
	ReqData QueryReceiverDetailReqData `json:"reqData"`
}

type QueryReceiverDetailRespData struct {
	// 接收方编号
	ReceiverNo string `json:"receiverNo"`
	// 分账接收方名称
	ReceiverName string `json:"receiverName"`
	// 联系手机号
	ContactMobile string `json:"contactMobile"`
	// 营业执照号码
	LicenseNo string `json:"licenseNo"`
	// 营业执照名称
	LicenseName string `json:"licenseName"`
	// 法人姓名
	LegalPersonName string `json:"legalPersonName"`
	// 法人证件类型
	// 17: 身份证, 18: 护照, 19: 港澳居民来往内地通行证, 20: 台湾居民来往内地通行证
	LegalPersonCertificateType string `json:"legalPersonCertificateType"`
	// 法人证件号
	LegalPersonCertificateNo string `json:"legalPersonCertificateNo"`
	// 收款账户卡号
	AcctNo string `json:"acctNo"`
	// 收款账户名称
	AcctName string `json:"acctName"`
	// 收款账户账户类型
	// 57: 对公, 58: 对私
	AcctTypeCode string `json:"acctTypeCode"`
	// 收款账户证件类型
	// 17: 身份证, 18: 护照, 19: 港澳居民来往内地通行证, 20: 台湾居民来往内地通行证
	AcctCertificateType string `json:"acctCertificateType"`
	// 收款账户证件号
	AcctCertificateNo string `json:"acctCertificateNo"`
	// 收款账户开户行号（仅支持对私结算账户）
	AcctOpenBankCode string `json:"acctOpenBankCode"`
	// 收款账户开户名称
	AcctOpenBankName string `json:"acctOpenBankName"`
	// 收款账户清算行行号（仅支持对私结算账户）
	AcctClearBankCode string `json:"acctClearBankCode"`
	// 创建方编号（开放平台创建传递，接收方拥有者）
	OwnerNo string `json:"ownerNo"`
}

type QueryReceiverDetailRet struct {
	Code     string                      `json:"retCode"`
	Msg      string                      `json:"retMsg"`
	RespData QueryReceiverDetailRespData `json:"respData"`
}
