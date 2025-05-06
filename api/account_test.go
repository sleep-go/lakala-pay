package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/sleep-go/lakala-pay/model"
)

func TestClient_FundAcctBalanceQuery(t *testing.T) {
	client := NewClient("OP10000788", "01963c4c9aa4", model.KEY_PATH_TEST, model.CERT_PATH_TEST, true, "")
	req := model.BalanceQueryReq{
		Ver:       "1.0.0",
		Timestamp: time.Now().Unix(),
		ReqId:     model.CreateOrderStr(),
		ReqData: model.BalanceQueryData{
			MerchantNo: "8221100076300BJ",
			PayNo:      "03",
		},
	}
	ret, err := client.BalanceQuery(&req)
	fmt.Println(ret)
	fmt.Println(err)
}
func TestQueryRefundAvailableBalance(t *testing.T) {
	client := NewClient("OP10000788", "01963c4c9aa4", model.KEY_PATH_TEST, model.CERT_PATH_TEST, true, "")
	req := model.QueryRefundAvailableBalanceReq{
		Ver:       "1.0.0",
		Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
		ReqId:     fmt.Sprintf("%d", time.Now().UnixMicro()),
		ReqData: struct {
			MerId       string `json:"merId"`
			TermId      string `json:"termId"`
			QryNo       string `json:"qryNo"`
			BalanceFlag string `json:"balanceFlag"`
			OlogNo      string `json:"ologNo"`
			SourceFlag  string `json:"sourceFlag"`
		}{
			MerId: "8221100076300BJ",
		},
	}
	ret, err := client.QueryRefundAvailableBalance(&req)
	fmt.Println(ret)
	fmt.Println(err)
}
