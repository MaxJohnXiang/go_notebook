package main

import "fmt"

func main() {
	// a := [...]int{1, 2, 4, 5, 6, 7}
	// b := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// reverse(a[:])
	// fmt.Println(a)
	// reverseV2(&b)
	// fmt.Println(b)
	// var runes []rune
	// for _, r := range "hello 时间" {
	// 	runes = append(runes, r)
	// }
	// fmt.Printf("%q\n", runes)
	st := [...]string{"a", "b", "c", "c", "d"}
	fmt.Println(noRepeatSite(st[:]))
}
func reverse(s []int) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseV2(s *[10]int) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func noRepeatSite(s []string) []string {
	count := len(s)
	for i := range s {
		if i == len(s)-1 {
			return s
		}
		if s[i] == s[i+1] {
			s = append(s[:i], s[i+1:])
			count--
		}
	}
	return s
}
