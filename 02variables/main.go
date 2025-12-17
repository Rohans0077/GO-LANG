package main

import "fmt"

func main() {

	var username string = "Rohan"
	fmt.Println("The name of the person is ", username)
	fmt.Printf("The data type of this username is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println("The value is ", isLoggedIn)

	fmt.Printf("The type is: %T \n", isLoggedIn)

	var score int = 98
	fmt.Println("The person scored", score, "marks")
	fmt.Printf("The type of score is %T \n", score)

	var percentage float32 = 87.45
	fmt.Println("The percentage obtained is ", percentage)
	fmt.Printf("The type of variable is %T ", percentage)
}
