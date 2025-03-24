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
	Code     string `json:"retCode"`
	Msg      string `json:"retMsg"`
	RespData struct {
		OrgCode   string `json:"orgCode" validate:"required"`   // 合作机构唯一标识码，需向平台申请获取（示例：ORG_2024）
		OrderNo   string `json:"orderNo" validate:"required"`   // 业务订单编号，格式：14位年月日时分秒+8位随机数（示例：202403110930001234）
		AttFileId string `json:"attFileId" validate:"required"` // 附件唯一标识符，由文件服务生成（示例：ATT_20240311_123456）‌
		AttType   string `json:"attType" validate:"required"`   // 附件类型枚举值：ID_CARD_FRONT-身份证正面，INVOICE-发票，CONTRACT-合同‌
	} `json:"respData"`
}

type CardBinReq struct {
	Ver     string         `json:"ver"`
	ReqTime string         `json:"timestamp"`
	ReqId   string         `json:"reqId"`
	ReqData CardBinReqData `json:"reqData"`
}

type CardBinReqData struct {
	// 接口版本号，必传字段，固定为"1.0"，长度为8
	Version string `json:"version" validate:"required,eq=1.0,max=8"`

	// 订单编号，必传字段，由14位时间（年月日时分秒）和8位随机数组成，总长度为32
	// 示例值："2021020112000012345678"
	// 注意：这里未直接在结构体中做格式校验，因为需要自定义逻辑来生成或验证该格式
	OrderNo string `json:"orderNo" validate:"required,max=32"`

	// 机构代码，必传字段，长度为32
	OrgCode string `json:"orgCode" validate:"required,max=32"`

	// 银行卡号，必传字段，长度为32（通常银行卡号长度为16-19位，但此处按给定长度32处理）
	// 注意：实际业务中，银行卡号可能需要进行加密或脱敏处理，这里仅作示例
	CardNo string `json:"cardNo" validate:"required,max=32"`
}

type CardBinRet struct {
	Code     string `json:"retCode"`
	Msg      string `json:"retMsg"`
	RespData struct {
		// 机构代码
		OrgCode string `json:"orgCode"`
		// 订单号
		OrderNo string `json:"orderNo"`
		// 银行卡号（在实际应用中，出于安全考虑，银行卡号可能会进行部分隐藏或加密处理）
		CardNo string `json:"cardNo"`
		// 卡bin（银行卡号的前几位，用于标识发卡银行和卡类型）
		CardBin string `json:"cardBin"`
		// 开户行号
		BankCode string `json:"bankCode"`
		// 清算行号
		ClearingBankCode string `json:"clearingBankCode"`
		// 开户行名称
		BankName string `json:"bankName"`
		// 银行卡类别（如：借记卡、信用卡等）
		CardType string `json:"cardType"`
		// 卡种名称（如：金卡、普卡、白金卡等）
		CardName string `json:"cardName"`
	} `json:"respData"`
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

type ModifyLedgerMerReq struct {
	Ver     string              `json:"version"`
	ReqTime string              `json:"reqTime"`
	ReqId   string              `json:"reqId"`
	ReqData ModifyLedgerMerData `json:"reqData"`
}

type ModifyLedgerMerData struct {
	// 接口版本号，必传字段，长度为8，固定取值"1.0"
	Version string `json:"version" validate:"required,eq=1.0,max=8"`

	// 订单编号，必传字段，长度为32，格式为14位年月日时（24小时制）分秒+8位的随机数（不重复）
	// 例如："2021020112000012345678"
	OrderNo string `json:"orderNo" validate:"required,max=32,regexp=\\d{14}\\d{8}$"`

	// 机构代码，必传字段，长度为32
	OrgCode string `json:"orgCode" validate:"required,max=32"`

	// 拉卡拉内部商户号，可选字段，长度为32
	// 与银联商户号（MerCupNo）互斥，如果两者都传，则以内部商户号为准
	MerInnerNo string `json:"merInnerNo,omitempty" validate:"max=32"`

	// 银联商户号，可选字段，长度为32
	// 与拉卡拉内部商户号（MerInnerNo）互斥，如果两者都传，则以内部商户号为准
	MerCupNo string `json:"merCupNo,omitempty" validate:"max=32"`

	// 联系手机号，可选字段，长度为32
	ContactMobile string `json:"contactMobile,omitempty" validate:"max=32"`

	// 最低分账比例（百分比），可选字段，长度为32，支持2位精度
	// 例如："70" 或 "70.50"
	SplitLowestRatio string `json:"splitLowestRatio,omitempty" validate:"max=32,regexp=\\d+(\\.\\d{1,2})?$"`

	// 分账结算委托书文件名称，可选字段，长度为64
	// 当需要变更比例时必须传，格式为"分账结算委托书文件名称.pdf"
	SplitEntrustFileName string `json:"splitEntrustFileName,omitempty" validate:"max=64"`

	// 分账结算委托书文件路径，可选字段，长度为64
	// 当需要变更比例时必须传，调用商户入网接口上传附件后反馈的文件路径
	SplitEntrustFilePath string `json:"splitEntrustFilePath,omitempty" validate:"max=64"`

	// 分账范围，非必传字段，长度为32
	// 取值："ALL"（商户全量交易自动分账处理）或 "MARK"（按交易请求分账标识进行分账处理）
	SplitRange string `json:"splitRange,omitempty" validate:"max=32,oneof=ALL MARK"`

	// 分账规则来源，非必传字段，长度为32
	// 取值："MER"（商户分账规则）或 "PLATFORM"（平台分账规则）
	SplitRuleSource string `json:"splitRuleSource,omitempty" validate:"max=32,oneof=MER PLATFORM"`

	// 回调通知地址，必传字段，长度为128
	// 审核通过后通知的地址
	RetUrl string `json:"retUrl" validate:"required,max=128"`

	// 电子合同编号，非必传字段，长度为32
	// 如果已经签署过电子合同，此处上送电子合同编号，供审核人员复核使用
	EleContractNo string `json:"eleContractNo,omitempty" validate:"max=32"`

	// 附加资料，可选字段，为附加资料文件信息的集合
	Attachments []Attach `json:"attachments,omitempty"`
}

type ModifyLedgerMerRet struct {
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

type ModifyLedgerReceiverReq struct {
	Ver     string                   `json:"version"`
	ReqTime string                   `json:"reqTime"`
	ReqId   string                   `json:"reqId"`
	ReqData ModifyLedgerReceiverData `json:"reqData"`
}

type ModifyLedgerReceiverData struct {
	// 接口版本号，必传字段，长度为8，固定取值"1.0"
	Version string `json:"version" validate:"required,eq=1.0,max=8"`
	// 订单编号，必传字段，长度为32，格式为14位年月日时（24小时制）分秒+8位的随机数（不重复）
	// 例如："2021020112000012345678"
	OrderNo string `json:"orderNo" validate:"required,max=32,regexp=\\d{14}\\d{8}$"`
	// 分账接收方所属机构代码，必传字段，长度为32
	OrgCode string `json:"orgCode" validate:"required,max=32"`
	// 分账接收方编号，必传字段，长度为32
	ReceiverNo string `json:"receiverNo" validate:"required,max=32"`
	// 分账接收方名称，可选字段，长度为64
	ReceiverName string `json:"receiverName,omitempty" validate:"max=64"`
	// 联系手机号，可选字段，长度为16
	ContactMobile string `json:"contactMobile,omitempty" validate:"max=16"`
	// 收款账户卡号，可选字段，长度为32
	AcctNo string `json:"acctNo,omitempty" validate:"max=32"`
	// 收款账户账户类型，可选字段，长度为32，57：对公 58：对私
	AcctTypeCode string `json:"acctTypeCode,omitempty" `
	// 收款账户开户行号，可选字段，长度为32
	AcctOpenBankCode string `json:"acctOpenBankCode,omitempty" validate:"max=32"`
	// 收款账户开户名称，可选字段，长度为64
	AcctOpenBankName string `json:"acctOpenBankName,omitempty" validate:"max=64"`
	// 收款账户清算行行号，可选字段，长度为32
	AcctClearBankCode string `json:"acctClearBankCode,omitempty" validate:"max=32"`
	// 附件资料集合，可传字段
	AttachList []Attach `json:"attachList,omitempty"`
	// 接收方状态，可传字段，长度为32，有效：VALID，无效：INVALID
	Status string `json:"status,omitempty" validate:"max=32"`
}

type ModifyLedgerReceiverRet struct {
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

type ApplyBindReq struct {
	Ver     string        `json:"version"`
	ReqTime string        `json:"reqTime"`
	ReqId   string        `json:"reqId"`
	ReqData ApplyBindData `json:"reqData"`
}

type ApplyBindData struct {
	// 接口版本号
	Version string `json:"version"`
	// 订单编号（14位年月日时分秒+8位随机数）
	OrderNo string `json:"orderNo"`
	// 分账接收方所属机构代码
	OrgCode string `json:"orgCode"`
	// 分账商户内部商户号（与MerCupNo选传其一，不能都为空）
	MerInnerNo string `json:"merInnerNo"`
	// 分账商户银联商户号（与MerInnerNo选传其一，不能都为空）
	MerCupNo string `json:"merCupNo"`
	// 分账接收方编号
	ReceiverNo string `json:"receiverNo"`
	// 合作协议附件名称
	EntrustFileName string `json:"entrustFileName"`
	// 合作协议附件路径（调用进件附件上传接口获取到附件路径）
	EntrustFilePath string `json:"entrustFilePath"`
	// 回调通知地址（审核通过后通知的地址）
	RetUrl string `json:"retUrl"`
}

type ApplyBindRet struct {
	Code     string `json:"retCode"`
	Msg      string `json:"retMsg"`
	RespData struct {
		// 接口版本号（注意：这里的类型应该是根据实际的接口文档来确定的，
		// 如果接口文档确实指定为String类型且给出了这样的取值，则保留为String，
		// 但通常版本号可能是int或float类型，这里的类型选择应基于实际接口规范）
		Version string `json:"version"`
		// 订单编号
		OrderNo string `json:"orderNo"`
		// 机构代码
		OrgCode string `json:"orgCode"`
		// 受理编号
		ApplyId int64 `json:"applyId"`
	} `json:"respData"`
}

type ApplyUnBindReq struct {
	Ver     string          `json:"version"`
	ReqTime string          `json:"reqTime"`
	ReqId   string          `json:"reqId"`
	ReqData ApplyUnBindData `json:"reqData"`
}

type ApplyUnBindData struct {
	// 接口版本号，固定为"1.0"
	Version string `json:"version" validate:"required,eq=1.0"`
	// 订单编号，14位时间+8位随机数
	OrderNo string `json:"orderNo" validate:"required"` // 验证逻辑在业务层
	// 分账接收方所属机构代码
	OrgCode string `json:"orgCode" validate:"required"`
	// 分账商户内部商户号或银联商户号（二选一）
	MerInnerNo string `json:"merInnerNo" validate:"omitempty,max=32"` // 验证逻辑在业务层处理二选一
	MerCupNo   string `json:"merCupNo" validate:"omitempty,max=32"`   // 同上
	// 分账接收方编号
	ReceiverNo string `json:"receiverNo" validate:"required"`
	// 解除分账说明附件名称
	EntrustFileName string `json:"entrustFileName" validate:"required"`
	// 解除分账说明附件路径
	EntrustFilePath string `json:"entrustFilePath" validate:"required"`
	// 备注说明
	Remark string `json:"remark" validate:"max=128"`
	// 回调通知地址
	RetUrl string `json:"retUrl" validate:"required"`
}

type ApplyUnBindRet struct {
	Code     string `json:"retCode"`
	Msg      string `json:"retMsg"`
	RespData struct {
		// 接口版本号（注意：这里的类型应该是根据实际的接口文档来确定的，
		// 如果接口文档确实指定为String类型且给出了这样的取值，则保留为String，
		// 但通常版本号可能是int或float类型，这里的类型选择应基于实际接口规范）
		Version string `json:"version"`
		// 订单编号
		OrderNo string `json:"orderNo"`
		// 机构代码
		OrgCode string `json:"orgCode"`
		// 受理编号
		ApplyId int64 `json:"applyId"`
	} `json:"respData"`
}

type BalanceQueryReq struct {
	Ver     string           `json:"ver"`
	ReqTime string           `json:"reqTime"`
	ReqId   string           `json:"reqId"`
	ReqData BalanceQueryData `json:"reqData"`
}

type BalanceQueryData struct {
	// bmcp机构号，必填，最大长度32
	OrgNo string `json:"orgNo"`

	// 商户号、接收方编号（二选一或同时填写），必填，最大长度32
	// 商户号 或 receiveNo 或 商户用户编号
	MerchantNo string `json:"merchantNo"`

	// 账号（若提供，则payType无效），非必填，最大长度32
	PayNo string `json:"payNo,omitempty"`

	// 账号类型（若payNo未提供，则可能使用），非必填，默认01，最大长度32
	//账号类型（01：收款账户，02：付款账户，03：分账商户账户，04：分账接收方账户，05：充值代付账户，06：结算代付账户）- 未上送则默认为01
	PayType string `json:"payType,omitempty"` // 默认值为"01"（收款账户）

	// 账户标志（未提供则默认为01），非必填，最大长度32（待上线功能）
	// 账户标志（01:一般户 03:子虚户）- 未上送则默认为01
	MgtFlag string `json:"mgtFlag,omitempty"`
}

type BalanceQueryRet struct {
	// 账号
	PayNo string `json:"payNo"`

	// 账户类型
	PayType string `json:"payType"`

	// 账户状态
	// CLOSE: 销户
	// NORMAL: 正常
	// FREEZE: 冻结
	// STOPPAY: 止付
	AcctSt string `json:"acctSt"`

	// 预付余额（单位：元）
	ForceBalance float64 `json:"forceBalance"`

	// 上日余额（单位：元）–该字段已废弃使用
	// 注意：虽然此字段已废弃，但仍需保留以兼容旧系统或数据记录
	HisBalance float64 `json:"hisBalance"`

	// 实时余额（单位：元）
	ReBalance float64 `json:"reBalance"`

	// 当前可用余额（单位：元）
	CurBalance float64 `json:"curBalance"`
}
