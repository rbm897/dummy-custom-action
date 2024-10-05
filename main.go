package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("INPUT_name")
	fmt.Printf("Name = %s", name)
}
