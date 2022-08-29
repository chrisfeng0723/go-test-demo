/**
 * @Author: fxl
 * @Description:
 * @File:  gock_test.go
 * @Version: 1.0.0
 * @Date: 2022/8/29 15:00
 */
package gock

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestGetResultByAPI(t *testing.T) {
	defer gock.Off()
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "x=1", args: args{x: 1, y: 100}, want: 101},
		{name: "x=2", args: args{x: 2, y: 200}, want: 202},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := GetResultByAPI(tt.args.x, tt.args.y); got != tt.want {
			//	t.Errorf("GetResultByAPI() = %v, want %v", got, tt.want)
			//}
			gock.New("http://your-api.com/").
				Post("/post").
				MatchType("json").
				JSON(map[string]int{"x": tt.args.x}).
				Reply(200).
				JSON(map[string]int{"value": tt.args.y})

			res := GetResultByAPI(tt.args.x, tt.args.x)
			assert.Equal(t, res, tt.want)
			assert.True(t, gock.IsDone())
		})

	}
}
