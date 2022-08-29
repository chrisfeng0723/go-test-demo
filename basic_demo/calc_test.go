/**
 * @Author: fxl
 * @Description:
 * @File:  calc_test.go
 * @Version: 1.0.0
 * @Date: 2022/6/22 10:23
 */
package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("begin")
	fmt.Println(m.Run())
	fmt.Println("end")
}

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1+2 expeced be 3,but %d got", ans)
	}

	if ans := Add(-11, 20); ans != 9 {
		t.Errorf("-11+20 expeced be 9,but %d got", ans)

	}
}

func TestMul(t *testing.T) {
	if ans := Mul(1, 2); ans != 2 {
		t.Errorf("1*2 expeced be 2,but %d got", ans)
	}

	if ans := Mul(-11, 20); ans != -220 {
		t.Errorf("-11*20 expeced be 220,but %d got", ans)

	}
}
