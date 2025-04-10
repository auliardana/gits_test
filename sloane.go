package main

import (
	"fmt"
	"strconv"
	"strings"
)

func a000124(n int) string {
	var result []string
	for i := 0; i < n; i++ {
		val := (i * (i + 1) / 2) + 1
		result = append(result, strconv.Itoa(val))
	}
	return strings.Join(result, "-")
}

func main() {
	fmt.Println(a000124(7))
	fmt.Println(a000124(5))
	fmt.Println(a000124(3)) 
}