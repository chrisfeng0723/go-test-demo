/**
 * @Author: fxl
 * @Description:
 * @File:  network_test.go
 * @Version: 1.0.0
 * @Date: 2022/6/23 14:31
 */
package network_demo

import (
	"io"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost://example.com/hello", nil)
	w := httptest.NewRecorder()
	HelloHandler(w, req)
	bytes, _ := io.ReadAll(w.Result().Body)
	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}
