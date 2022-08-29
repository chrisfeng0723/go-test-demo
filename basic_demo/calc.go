/**
 * @Author: fxl
 * @Description:
 * @File:  calc.go
 * @Version: 1.0.0
 * @Date: 2022/6/22 10:22
 */
package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func main() {
	addResult := Add(1, 3)
	mulResult := Mul(3, 2)
	fmt.Println(addResult, mulResult)

}
