package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.ParseInt("0001", 10, 64)
	fmt.Println(i, err)
	fmt.Printf("%4d%4d\n", 12, 345)
}
