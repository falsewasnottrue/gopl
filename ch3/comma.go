package main

import "fmt"

func main() {
	fmt.Println(comma("1234456791214234"))
}

func comma(s string) string {
	n := len(s)
	if (n<=3) {
		return s
	}

	return comma(s[:len(s)-3]) + "," + s[len(s)-3:]
}