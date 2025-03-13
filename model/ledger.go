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
