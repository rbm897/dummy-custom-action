package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1, _ := strconv.Atoi(os.Getenv("INPUT_NUM1"))
	num2, _ := strconv.Atoi(os.Getenv("INPUT_NUM2"))
	sum := num1 + num2
	fmt.Printf("echo {sum}={%d} >> $GITHUB_OUTPUT", sum)
}
