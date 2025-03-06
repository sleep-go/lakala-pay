package model

// BaseReq 基础请求
type BaseReq[T any] struct {
	ReqTime string `json:"req_time" validate:"required"` // 请求时间，格式yyyyMMddHHmmss
	Version string `json:"version" validate:"required"`
	ReqData *T     `json:"req_data" validate:"required"`
}

// BaseResp 基础响应
type BaseResp[T any] struct {
	Code     string `json:"code"`                // 返回业务代码(000000为成功，其余按照错误信息来定)
	Msg      string `json:"msg"`                 // 返回业务代码描述
	RespTime string `json:"resp_time"`           // 响应时间，格式yyyyMMddHHmmss
	RespData *T     `json:"resp_data,omitempty"` // 返回数据.下文定义的响应均为该属性中的内容
}

const KEY_PATH_TEST = "../data/OP00000003_private_key.pem"
const CERT_PATH_TEST = "../data/lkl-apigw-v2.cer"

//const MERCHANT_NO_TEST = "82229007392000A"
//const APPID_TEST = "OP00000003"
//const SERIAL_NO_TEST = "00dfba8194c41b84cf"

const MERCHANT_NO_TEST = "82229007392000A"
const APPID_TEST = "OP00000003"
const SERIAL_NO_TEST = "00dfba8194c41b84cf"
const TERM_NO_TEST = "D9296400"
const KEY_TEST = "uIj6CPg1GZAY10dXFfsEAQ=="
