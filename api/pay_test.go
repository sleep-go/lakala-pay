package api

import (
	"fmt"
	"github.com/sleep-go/lakala-pay/model"
	"testing"
)

func TestPay(t *testing.T) {
	orderId := model.CreateOrderStr()
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, model.KEY_TEST)
	payinfo := model.PayerInfo{
		IdentityType: "OutBankCard",
		AcctNo:       "xxx",
		AcctName:     "xxx",
		BankName:     "中信银行xx支行",
		BankNo:       "xxx",
		AcctType:     "1",
	}
	req := model.PayReq{
		OutOrderNo: orderId,
		MerchantNo: model.MERCHANT_NO_TEST,
		TermNo:     model.TERM_NO_TEST,
		TxnAmt:     1,
		PeriodType: "D0",
		Remark:     "保证金退款",
		Summary:    "保证金退款",
		PayerInfo:  payinfo,
	}
	ret, err := client.Pay(&req)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestPayQuery(t *testing.T) {
	orderId := "1740981618428919732"
	PayOrderNo := "25030311012001101011001738446"
	PayOrderNo = ""
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.PayQueryReq{
		OutOrderNo: orderId,
		MerchantNo: model.MERCHANT_NO_TEST,
		TradeNo:    PayOrderNo,
		TermNo:     model.TERM_NO_TEST,
	}
	ret, err := client.PayQuery(&req)
	fmt.Println(ret)
	fmt.Println(err)
}
