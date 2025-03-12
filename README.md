# 拉卡拉sdk
https://o.lakala.com/#/home/document/detail?id=282

## 安装
```shell
go get -u github.com/sleep-go/lakala-pay
```
## 示例
```go
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
		OrderInfo:          "充值",
		ChannelID:          ChannelID,
	}
	ret, err := client.OrderSpecialCreate(&req)
	fmt.Println(ret)
	fmt.Println(err)
}

```
