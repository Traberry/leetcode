package main

import "fmt"

//simplified way, only use one count, no too much if...else...
func detectCapitalUse2(word string) bool {
	countU := 0
	for i := 0; i < len(word); i++ {
		if word[i] <97 {
			countU++
		}
	}
	return countU == 0 || (countU == 1 && word[0] < 97) || countU == len(word)
}

func main() {
	s := "FlaG"
	fmt.Println(detectCapitalUse2(s))
}
