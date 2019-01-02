package main

import (
	"bufio"
	"fmt"
	"os"
)

// _ is the blank identifier and it acts as a placeholder
func main() {
	fmt.Println("Hello, World.")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
