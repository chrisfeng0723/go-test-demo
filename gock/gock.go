/**
 * @Author: fxl
 * @Description:
 * @File:  gock.go
 * @Version: 1.0.0
 * @Date: 2022/8/29 14:54
 */
package gock

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ReqParam struct {
	X int `json:"x"`
}

type Response struct {
	Value int `json:"value"`
}

func GetResultByAPI(x, y int) int {
	p := &ReqParam{
		X: x,
	}
	b, _ := json.Marshal(p)

	//调用外部api
	resp, err := http.Post(
		"http://your-api.com/post",
		"application/json",
		bytes.NewBuffer(b),
	)

	if err != nil {
		return -1
	}

	body, _ := io.ReadAll(resp.Body)
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return -1
	}
	return result.Value + y

}
