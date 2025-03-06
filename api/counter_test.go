package api

import (
	"fmt"
	"testing"
	"time"
)
import "github.com/sleep-go/lakala-pay/model"

func TestCreate(t *testing.T) {
	orderId := model.CreateOrderStr()
	ChannelID := ""
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false)
	expeirTime := time.Now().Add(24 * time.Hour).Format("20060102150405")
	req := model.SpecialCreateReq{
		OutOrderNo:         orderId,
		MerchantNo:         model.MERCHANT_NO_TEST,
		TotalAmount:        100,
		OrderEfficientTime: expeirTime,
		OrderInfo:          "保证金充值",
		ChannelID:          ChannelID,
	}
	ret, err := client.OrderSpecialCreate(&req)
	fmt.Println(ret)
	fmt.Println(err)
	if err == nil {
		fmt.Println(ret.ResData.OutOrderNo)
		fmt.Println(ret.ResData.PayOrderNo)
	}
}

func TestQuery(t *testing.T) {
	orderId := "1740981618428919732"
	PayOrderNo := "25030311012001101011001738446"
	PayOrderNo = ""
	//orderId = ""
	ChannelID := ""
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false)
	req := model.OrderQueryReq{
		OutOrderNo: orderId,
		MerchantNo: model.MERCHANT_NO_TEST,
		PayOrderNo: PayOrderNo,
		ChannelID:  ChannelID,
	}
	ret, err := client.OrderQuery(&req)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestClose(t *testing.T) {
	orderId := "1740981618428919732"
	PayOrderNo := "25030311012001101011001738446"
	PayOrderNo = ""
	//orderId = ""
	ChannelID := ""
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false)
	req := model.OrderCloseReq{
		OutOrderNo: orderId,
		MerchantNo: model.MERCHANT_NO_TEST,
		PayOrderNo: PayOrderNo,
		ChannelID:  ChannelID,
	}
	ret, err := client.OrderClose(&req)
	fmt.Println(ret)
	fmt.Println(err)
}
