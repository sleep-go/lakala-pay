package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sleep-go/lakala-pay/model"
	"github.com/sleep-go/lakala-pay/util"
	"io"
	"net/http"
	"reflect"
)

const (
	specialCreateUrl = "/api/v3/ccss/counter/order/special_create"
	orderQueryUrl    = "/api/v3/ccss/counter/order/query"
	orderCloseUrl    = "/api/v3/ccss/counter/order/close"
	refundUrl        = "/api/v3/labs/relation/refund"
)

func hasField(i interface{}, fieldName string) bool {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	return v.FieldByName(fieldName).IsValid()
}

func newBuffer[T any](req *T) *bytes.Buffer {
	m := model.BaseReq[T]{
		ReqTime: util.GetReqTime(),
		Version: "3.0",
		ReqData: req,
	}
	data, err := json.Marshal(m)
	if hasField(req, "Ver") {
		data, err = json.Marshal(req)
	}
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(data)
}

func newBufferEncrypt[T any](req *T) *bytes.Buffer {
	m := model.BaseReq[T]{
		ReqTime: util.GetReqTime(),
		Version: "3.0",
		ReqData: req,
	}
	data, err := json.Marshal(m)
	if hasField(req, "Ver") {
		data, err = json.Marshal(req)
	}
	if err != nil {
		return nil
	}
	//fmt.Println("----------------")
	//fmt.Println("param:", string(data))

	key := model.KEY_TEST
	src := data
	endata, _ := EncryptECB([]byte(key), []byte(src))
	data = []byte(endata)

	return bytes.NewBuffer(data)
}

// doRequest 统一请求方法
func doRequest[T any, D any](c *Client, url string, req *T, needEncrypt bool) (*D, error) {
	var reqStr *bytes.Buffer
	if needEncrypt {
		reqStr = newBufferEncrypt[T](req)
	} else {
		reqStr = newBuffer[T](req)
	}
	fmt.Println("----------------")
	fmt.Println("param:", reqStr.String())
	auth, err := c.GetAuthorization(reqStr.Bytes())
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, c.Host+url, reqStr)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", auth)
	resp, err := c.Http.Do(request)
	//fmt.Println("-----------")
	//fmt.Println(auth)
	fmt.Println("-----------")
	p, _ := io.ReadAll(resp.Body)
	fmt.Println(string(p))
	fmt.Println("-----------")
	if err != nil {
		return nil, err
	}
	return util.ParseResp[D](resp)
}

// OrderSpecialCreate 收银台订单创建
func (c *Client) OrderSpecialCreate(req *model.SpecialCreateReq) (*model.SpecialCreateRes, error) {
	return doRequest[model.SpecialCreateReq, model.SpecialCreateRes](c, specialCreateUrl, req, false)
}

// OrderQuery 收银台订单查询
func (c *Client) OrderQuery(req *model.OrderQueryReq) (*model.OrderQueryRes, error) {
	return doRequest[model.OrderQueryReq, model.OrderQueryRes](c, orderQueryUrl, req, false)
}

// OrderClose 收银台订单关单
func (c *Client) OrderClose(req *model.OrderCloseReq) (resp *model.OrderCloseRes, err error) {
	return doRequest[model.OrderCloseReq, model.OrderCloseRes](c, orderCloseUrl, req, false)
}
