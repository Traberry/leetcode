package main

import "fmt"

func reverseString(s string) string {
	var b []byte
	for i := len(s) - 1; i >= 0 ; i-- {
		b = append(b, s[i])
	}
	return string(b)
}

func main() {
	s := "absd"
	fmt.Println(reverseString(s))
}

//I got aware of one Golang grammar different with others
func reverseString2(s string) string {
	i := 0
	j := len(s) - 1
	for i < j {
		temp := s[i]
		s[i] = s[j] // Error: cannot assign to s[i]
		s[j] = temp // Error: cannot assign to s[j]

		i++
		j--
	}
}
