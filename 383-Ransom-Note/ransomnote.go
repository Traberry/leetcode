package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	var r [26]int
	var m [26]int

	for i:= 0; i < len(magazine); i++ {
		v := magazine[i]
		m[v - 'a']++
	}

	for i:= 0; i < len(ransomNote); i++ {
		v := ransomNote[i]
		r[v - 'a']++
	}

	for i := 0; i < 26; {
		if r[i] <= m[i] {
			i++
		}else {
			return false
		}
	}

	return true
}

func main() {
	ran := "faassd"
	mag := "fdgrfaklass"

	fmt.Println(canConstruct(ran, mag))
}
