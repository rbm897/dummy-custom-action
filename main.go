package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("INPUT_NAME")
	fmt.Printf("Name = %s", name)
}
