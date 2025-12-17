package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to Go lang"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the username: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("The user logged in is", input)

	fmt.Println()

}
