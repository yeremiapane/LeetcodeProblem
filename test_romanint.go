package main

import "fmt"

func main() {
	src := "1, 2, 3"
	var a, b, c int

	parsed, err := fmt.Sscanf(src, "%d-%d-%d", &a, &b, &c)
	fmt.Println(parsed, err, a, b, c)
}
