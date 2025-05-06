package api

import (
	"fmt"
	"testing"

	"github.com/sleep-go/lakala-pay/model"
)

func TestClient_FundAcctBalanceQuery(t *testing.T) {
	client := NewClient(model.APPID_TEST, model.SERIAL_NO_TEST, model.KEY_PATH_TEST, model.CERT_PATH_TEST, false, "")
	req := model.FundAcctBalanceQueryReq{
		OwnerId: "982511",
		CustId:  "8221100076300BJ",
		AcctNbr: "",
	}
	ret, err := client.FundAcctBalanceQuery(&req)
	fmt.Println(ret)
	fmt.Println(err)
}
