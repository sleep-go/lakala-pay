package api

import (
	"encoding/base64"
	"fmt"
	"github.com/sleep-go/lakala-pay/model"
	"log"
	"os"
	"testing"
	"time"
)

func TestUpload(t *testing.T) {
	orderId := model.CreateOrderStr()
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	file := "../data/name.pdf"
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// 转换为Base64编码的字符串
	base64Content := base64.StdEncoding.EncodeToString(content)
	log.Println(len(base64Content))
	req := model.UploadReqData{
		Version:    "1.0",
		OrderNo:    orderId,
		OrgCode:    "1",
		AttType:    "NETWORK_XY",
		AttExtName: "pdf",
		AttContext: base64Content,
	}
	upReg := model.UploadReq{
		Ver:     "1.0.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}
	ret, err := client.upload(&upReg)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestCardBin(t *testing.T) {
	orderId := model.CreateOrderStr()
	CardNo := "6217681406150014"
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.CardBinReqData{
		Version: "1.0", OrderNo: orderId, OrgCode: "1", CardNo: CardNo,
	}
	upReg := model.CardBinReq{
		Ver:     "1.0.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}
	ret, err := client.cardBin(&upReg)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestApply(t *testing.T) {
	orderId := model.CreateOrderStr()
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.ApplyReqData{
		Version:              "1.0",
		OrderNo:              orderId,
		OrgCode:              "1",
		MerInnerNo:           "822290059430BFE",
		ContactMobile:        "13263116556",
		SplitLowestRatio:     1.01,
		SplitEntrustFileName: "授权委托书.pdf",
		SplitEntrustFilePath: "MMS/20250311/151221-e802075955e24dc6be1ba9c4109ef8f2.pdf",
		SplitRange:           "MARK",
		SepFundSource:        "BA",
		SplitLaunchMode:      "MANUAL",
		SettleType:           "01",
		RetUrl:               "http://sss.ss.s",
	}
	applyReg := model.ApplyReq{
		Ver:     "2.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}

	ret, err := client.ledgerApply(&applyReg)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestModifyLedgerMer(t *testing.T) {
	orderId := model.CreateOrderStr()
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.ModifyLedgerMerData{
		Version:              "1.0",
		OrderNo:              orderId,
		OrgCode:              "1",
		MerInnerNo:           "822290059430BFE",
		ContactMobile:        "13263116556",
		SplitLowestRatio:     "1.01",
		SplitEntrustFileName: "授权委托书.pdf",
		SplitEntrustFilePath: "MMS/20250311/151221-e802075955e24dc6be1ba9c4109ef8f2.pdf",
		SplitRange:           "MARK",
		RetUrl:               "http://sss.ss.s",
	}
	applyReg := model.ModifyLedgerMerReq{
		Ver:     "2.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}

	ret, err := client.modifyLedgerMer(&applyReg)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestLedgerQuery(t *testing.T) {
	orderId := model.CreateOrderStr()
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.LedgerQueryData{
		Version:    "1.0",
		OrderNo:    orderId,
		OrgCode:    "1",
		MerInnerNo: "822290059430BFE",
	}
	applyReg := model.LedgerQueryReq{
		Ver:     "2.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}

	ret, err := client.ledgerQuery(&applyReg)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestApplyLedgerReceiver(t *testing.T) {
	orderId := model.CreateOrderStr()
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")

	attach := model.Attach{
		AttachType:      "BANK_CARD",
		AttachName:      "银行卡",
		AttachStorePath: "MMS/20250311/151221-e802075955e24dc6be1ba9c4109ef8f2.pdf",
	}
	attachList := []model.Attach{attach}
	req := model.ApplyLedgerReceiverData{
		Version:             "1.0",
		OrderNo:             orderId,
		OrgCode:             "1",
		ReceiverName:        "高峰公棚",
		ContactMobile:       "13263116556",
		AcctNo:              "22",
		AcctName:            "授权",
		AcctTypeCode:        "58",
		AcctCertificateType: "17",
		AcctCertificateNo:   "23432",
		AcctOpenBankCode:    "12312",
		AcctOpenBankName:    "天津银行",
		AcctClearBankCode:   "12312",
		SettleType:          "01",
		AttachList:          attachList,
	}

	applyReg := model.ApplyLedgerReceiverReq{
		Ver:     "2.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}

	ret, err := client.applyLedgerReceiver(&applyReg)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestModifyLedgerReceiver(t *testing.T) {
	orderId := model.CreateOrderStr()
	ReceiverNo := "123312" //applyLedgerReceiver 接口返回的
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")

	attach := model.Attach{
		AttachType:      "BANK_CARD",
		AttachName:      "银行卡",
		AttachStorePath: "MMS/20250311/151221-e802075955e24dc6be1ba9c4109ef8f2.pdf",
	}
	attachList := []model.Attach{attach}
	req := model.ModifyLedgerReceiverData{
		Version:           "1.0",
		OrderNo:           orderId,
		OrgCode:           "1",
		ReceiverName:      "高峰公棚",
		ContactMobile:     "13263116556",
		AcctNo:            "22",
		AcctTypeCode:      "58",
		AcctOpenBankCode:  "12312",
		AcctOpenBankName:  "天津银行",
		AcctClearBankCode: "12312",
		AttachList:        attachList,
		ReceiverNo:        ReceiverNo,
	}

	applyReg := model.ModifyLedgerReceiverReq{
		Ver:     "2.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}

	ret, err := client.modifyLedgerReceiver(&applyReg)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestQueryReceiverDetail(t *testing.T) {
	orderId := model.CreateOrderStr()
	ReceiverNo := "123312" //applyLedgerReceiver 接口返回的
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.QueryReceiverDetailReqData{
		Version:    "1.0",
		OrderNo:    orderId,
		OrgCode:    "1",
		ReceiverNo: ReceiverNo,
	}

	applyReg := model.QueryReceiverDetailReq{
		Ver:     "2.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}

	ret, err := client.queryReceiverDetail(&applyReg)
	fmt.Println(ret)
	fmt.Println(err)
	if err == nil {
		fmt.Println(ret.RespData.ReceiverNo) //绑定申请要用到
	}
}

func TestApplyBind(t *testing.T) {
	orderId := model.CreateOrderStr()
	ReceiverNo := "123312" //applyLedgerReceiver 接口返回的
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.ApplyBindData{
		Version:    "1.0",
		OrderNo:    orderId, // 示例订单编号，需要符合格式要求
		OrgCode:    "1",
		MerInnerNo: model.MERCHANT_NO_TEST, // 或者填写MerCupNo字段
		// MerCupNo:    "exampleMerCupNo",       // 与MerInnerNo选传其一
		ReceiverNo:      ReceiverNo,
		EntrustFileName: "cooperation_agreement.pdf",
		EntrustFilePath: "/path/to/uploaded/file", // 调用进件附件上传接口获取到的路径
		RetUrl:          "http://example.com/callback",
	}

	applyReg := model.ApplyBindReq{
		Ver:     "2.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}

	ret, err := client.applyBind(&applyReg)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestApplyUnBind(t *testing.T) {
	orderId := model.CreateOrderStr()
	ReceiverNo := "123312" //applyLedgerReceiver 接口返回的
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.ApplyUnBindData{
		Version:    "1.0",
		OrderNo:    orderId, // 示例订单编号，需要符合格式要求
		OrgCode:    "1",
		MerInnerNo: model.MERCHANT_NO_TEST, // 或者填写MerCupNo字段
		// MerCupNo:    "exampleMerCupNo",       // 与MerInnerNo选传其一
		ReceiverNo:      ReceiverNo,
		EntrustFileName: "cooperation_agreement.pdf",
		EntrustFilePath: "/path/to/uploaded/file", // 调用进件附件上传接口获取到的路径
		RetUrl:          "http://example.com/callback",
	}

	applyReg := model.ApplyUnBindReq{
		Ver:     "2.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}

	ret, err := client.applyUnBind(&applyReg)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestBalanceQuery(t *testing.T) {
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.BalanceQueryData{
		MerchantNo: model.MERCHANT_NO_TEST,
		OrgNo:      "1",
	}

	applyReg := model.BalanceQueryReq{
		Ver:     "1.0.0",
		ReqData: req,
		ReqTime: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:   fmt.Sprintf("%d", time.Now().UnixMicro()),
	}

	ret, err := client.balanceQuery(&applyReg)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestBalanceSeparate(t *testing.T) {
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	orderId := model.CreateOrderStr()
	req := model.BalanceSeparateReq{
		MerchantNo:    model.MERCHANT_NO_TEST,
		OutSeparateNo: orderId,
		TotalAmt:      "100",
	}
	applyReg := model.RecvData{
		RecvMerchantNo: model.MERCHANT_NO_TEST,
		RecvNo:         "1",
		SeparateValue:  "10",
	}
	req.RecvDatas = append(req.RecvDatas, applyReg)

	ret, err := client.balanceSeparate(&req)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestBalanceCancel(t *testing.T) {
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	orderId := model.CreateOrderStr()
	req := model.BalanceCancelReq{
		MerchantNo:          model.MERCHANT_NO_TEST,
		OutSeparateNo:       orderId,
		TotalAmt:            "100",
		OriginOutSeparateNo: "131312",
	}

	ret, err := client.balanceCancel(&req)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestBalanceFallback(t *testing.T) {
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	orderId := model.CreateOrderStr()
	req := model.BalanceFallbackReq{
		MerchantNo:          model.MERCHANT_NO_TEST,
		OutSeparateNo:       orderId,
		TotalAmt:            "10",
		OriginOutSeparateNo: "131312",
	}
	applyReg := model.OriginRecvData{
		RecvNo: "1",
		Amt:    "10",
	}
	req.OriginRecvDatas = append(req.OriginRecvDatas, applyReg)

	ret, err := client.balanceFallback(&req)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestBalanceSeparateQuery(t *testing.T) {
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	orderId := model.CreateOrderStr()
	req := model.BalanceSeparateQueryReq{
		MerchantNo:    model.MERCHANT_NO_TEST,
		OutSeparateNo: orderId, //应该是分账时的orderid
	}

	ret, err := client.balanceSeparateQuery(&req)
	fmt.Println(ret)
	fmt.Println(err)
}
