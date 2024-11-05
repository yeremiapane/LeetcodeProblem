package main

import "fmt"

func minChanges(s string) int {
	temp := 0

	for i := 0; i < len(s)-1; i += 2 {
		if s[i] != s[i+1] {
			temp += 1
		}
	}
	return temp
}

func main() {
	fmt.Println(minChanges("1001"))     //expected : 2
	fmt.Println(minChanges("10"))       //expected : 1
	fmt.Println(minChanges("0000"))     //expected : 0
	fmt.Println(minChanges("11000111")) //expected : 1

}
