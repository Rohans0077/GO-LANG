package main

import "fmt"

// returns string
func greet(name string) string {

	fmt.Println("The name of the person is ", name)
	var message string = "sos"
	return message

}

func main() {

	var message string = greet("Rohan")
	fmt.Println("The message returned from greet is :", message)

}
