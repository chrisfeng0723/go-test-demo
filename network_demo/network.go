/**
 * @Author: fxl
 * @Description:
 * @File:  network.go
 * @Version: 1.0.0
 * @Date: 2022/6/23 14:30
 */
package network_demo

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
