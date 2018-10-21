package main

import "fmt"

func detectCapitalUse(word string) bool {
	countU, countL := 0, 0
	for i := 0; i < len(word); i++ {
		if word[i] >= 65 && word[i] <= 90 {
			countU++
		}else {
			countL++
		}
	}

	if countU == 0 {
		return true
	}else if countU == 1 {
		if word[0] >= 65 && word[0] <= 90 {
			return true
		}
	}else {
		if countL == 0 {
			return true
		}
	}

	return false
}

func main() {
	s := "AbcdE"
	fmt.Println(detectCapitalUse(s))
}
