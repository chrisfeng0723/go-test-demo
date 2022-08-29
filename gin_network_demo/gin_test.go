/**
 * @Author: fxl
 * @Description:
 * @File:  gin_test.go
 * @Version: 1.0.0
 * @Date: 2022/8/29 14:41
 */
package gin_network_demo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// 注意：如果测试函数是小写，则有中间下划线，否则是驼峰
func Test_helloHandler(t *testing.T) {

	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"right case", `{"name": "孙悟空"}`, "hello 孙悟空"},
		{"lack param", "", "we need a name"},
	}
	r := SetupRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodPost,
				"/hello",
				strings.NewReader(tt.param),
			)
			//mock 响应记录器
			w := httptest.NewRecorder()

			//server处理mock请求并记录返回的响应内容
			r.ServeHTTP(w, req)
			//查看状态码是否符合预期
			assert.Equal(t, http.StatusOK, w.Code)
			//解析响应结果，查看是否符合预期
			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["msg"])
		})
	}
}
