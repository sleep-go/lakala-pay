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

	ret, err := client.apply(&applyReg)
	fmt.Println(ret)
	fmt.Println(err)
}
