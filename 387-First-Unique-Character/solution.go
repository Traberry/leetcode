package main

import "fmt"

func firstUniqChar(s string) int {
	var freq [26]int

	//第一次遍历，得到频率统计数组
	for i := 0; i < len(s); i++ {
		freq[s[i] - 'a']++
	}

	//第二次遍历，利用freq数组找到第一个不重复的字符
	for j := 0; j < len(s); j++ {
		if freq[s[j] - 'a'] == 1 {
			return j
		}
	}
	return -1
}

func main() {
	test := "zabcdefgabcdef"
	fmt.Println(firstUniqChar(test)) //answer is 0
}
