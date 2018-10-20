package main

import (
	"strings"
	"fmt"
)

// 一般方法
func checkRecord(s string) bool {
	countA, countL := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			countA++
		}

		if s[i] == 'L' {
			countL++
		}else {
			countL = 0
		}

		if countA > 1 || countL > 2 {
			return false
		}

	}

	return true
}

func main() {
	test := "PLLPAPLLLP"
	fmt.Println(checkRecord(test))
}

// 走了一下捷径
func checkRecord2(s string) bool {
	countA := 0

	if strings.Contains(s, "LLL") {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			countA++
		}
		if countA > 1 {
			return false
		}
	}

	return true
}

// 以牺牲效率换取代码的简洁
func checkRecord3(s string) bool {
	//使用正则表达式，一行代码就能解决

	return true
}