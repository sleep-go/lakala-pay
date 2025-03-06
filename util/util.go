package util

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func GetReqTime() string {
	return FormatTime(time.Now())
}

// FormatTime 格式化时间
func FormatTime(t time.Time) string {
	return t.Format("20060102150405")
}

func RandStr(n int) string {
	var letters = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func ParseResp[T any](resp *http.Response) (body *T, err error) {
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
