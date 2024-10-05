package main

import (
	"fmt"
	"os"
)

func main() {
	num1 := os.Getenv("INPUT_NUM1")
	num2 := os.Getenv("INPUT_NUM2")
	sum := num1 + num2
	fmt.Printf("::set-output name=sum::%s", sum)
}
