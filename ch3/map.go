package main

import "fmt"

func main() {
	// ages1 := make(map[string]int)
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	fmt.Println(ages["xi"])
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
