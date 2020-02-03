package main

import (
	"fmt"
)

func main() {
	var a, b int
	var ans string
	fmt.Scan(&a, &b)
	if a <= 8 && b <= 8 {
		ans = "Yay!"
	} else {
		ans = ":("
	}

	fmt.Println(ans)

}
