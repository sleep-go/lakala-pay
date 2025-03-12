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
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	expeirTime := time.Now().Add(24 * time.Hour).Format("20060102150405")
	req := model.SpecialCreateReq{
		OutOrderNo:         orderId,
		MerchantNo:         model.MERCHANT_NO_TEST,
		TotalAmount:        3,
		TermNo:             model.TERM_NO_TEST,
		OrderEfficientTime: expeirTime,
		OrderInfo:          "保证金充值",
		ChannelID:          ChannelID,
		SupportRefund:      1,
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
	orderId := "202503121452037062272"
	PayOrderNo := "25031211012001101011001747089"
	//PayOrderNo = ""
	////orderId = ""
	ChannelID := ""
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.OrderQueryReq{
		OutOrderNo: orderId,
		MerchantNo: model.MERCHANT_NO_TEST,
		PayOrderNo: PayOrderNo,
		ChannelID:  ChannelID,
	}
	ret, err := client.OrderQuery(&req)
	fmt.Println(ret)
	fmt.Println(err)
	if err == nil {
		fmt.Println(ret.ResData.TermNo)                        //退款用
		fmt.Println(ret.ResData.OrderTradeInfoList[0].TradeNo) //退款用
		fmt.Println(ret.ResData.OrderTradeInfoList[0].LogNo)   //退款用
	}
}

func TestClose(t *testing.T) {
	orderId := "1740981618428919732"
	PayOrderNo := "25030311012001101011001738446"
	PayOrderNo = ""
	//orderId = ""
	ChannelID := ""
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
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

func TestVerify(t *testing.T) {
	var body = `{"channel_id":"95","merchant_no":"82229007392000A","order_create_time":"20250311195529","order_efficient_time":"20250315152558","order_info":"这个第一个测试订单，不要删除","order_status":"2","order_trade_info":{"acc_discount_amount":"","acc_mdiscount_amount":"","acc_other_discount_amount":"","acc_settle_amount":"1","acc_trade_no":"2025031122001468521412228883","acc_type":"04","busi_type":"SCPAY","log_no":"66200819100678","pay_mode":"ALIPAY","payer_amount":1,"settle_merchant_no":"82229007392000A","settle_term_no":"D9296381","trade_amount":1,"trade_no":"2025031166200819100678","trade_remark":"","trade_status":"S","trade_time":"20250311195612","trade_type":"PAY","user_id1":"sky***@163.com","user_id2":"2088002264868521"},"out_order_no":"202020202028","pay_order_no":"25031111012001101011001746791","term_no":"D9296381","total_amount":1,"trans_merchant_no":"82229007392000A","trans_term_no":"D9296381"}`
	var au = `LKLAPI-SHA256withRSA timestamp="1741694218",nonce_str="dGl8VibGpN0w",signature="ejpsXcHGtJucdoFPE8jIdYeVgJv7wSRTURldTsTkqMhTpgXTFTrwPRcWKZVVOaePIJ404YUkZVRT1STvWzHr83Xjs/mwk/4uwsZBQsaJMD+zQG9CBsnrRt8UoRyTCtfO5wJ7QVZjIP+QAjNqn2lHFv32tf714IImyvGGBylETw/KziZa3rGvnvsBguH37UjBZ4bCTbaf3mGNCoCdcFQLQPxJIuIwlHH1BzfCqyLgKB46UPKumjnXOmlnQgPjt3uQBxXpgsFDnYWUVJpyvljx7Yf+IXBCiItzqELI0/ZjR6/GhuEJnRuLOB0vbzQg3OeuSbCvutll8/3J0N5qpPOkeg=="`
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	verification := client.SignatureVerification(au, body)
	fmt.Println(verification)
}

func TestRefund(t *testing.T) {
	OriginTradeNo := "2025031266200819110009"
	RefundOrderId := model.CreateOrderStr()
	fmt.Println(RefundOrderId)
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.RefundRequest{
		MerchantNo:      model.MERCHANT_NO_TEST,
		TermNo:          "D9296381",    //model.TERM_NO_TEST,
		OutTradeNo:      RefundOrderId, //必传，退款查询时有用
		RefundAmount:    "1",
		OriginBizType:   "3",
		OriginTradeDate: "20250312",
		OriginTradeNo:   OriginTradeNo,
		//OriginLogNo: "66200819101562",
	}
	ret, err := client.OrderRefund(&req)
	fmt.Println(ret)
	fmt.Println(err)
}

func TestRefundQuery(t *testing.T) {
	RefundQueryOrderId := model.CreateOrderStr()
	OriginTradeNo := "2025031266200819110009"
	fmt.Println(RefundQueryOrderId)
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.RefundQueryRequest{
		MerchantNo:      model.MERCHANT_NO_TEST,
		TermNo:          "D9296381", //model.TERM_NO_TEST,
		OutTradeNo:      RefundQueryOrderId,
		OriginBizType:   "3",
		OriginTradeDate: "20250312",
		//OriginTradeRefNo: "66200819101562",
		OriginTradeNo: OriginTradeNo,
	}
	ret, err := client.RefundQuery(&req)
	fmt.Println(ret)
	fmt.Println(err)
}
