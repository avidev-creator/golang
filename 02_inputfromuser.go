package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	welcome := "Welcome User !! "
	fmt.Println(welcome)
	fmt.Println("Please Rate Us from 1 to 5")

	reader := bufio.NewReader(os.Stdin) // Taking input from user.
	input, _ := reader.ReadString('\n') // comma, ok syntax to get the input untill user press Enter.
	result, err := strconv.ParseFloat(strings.TrimSpace(input),64) // Saving the whole operation in result and err for errors.

	if err != nil{
		fmt.Println(err) // To display any error we get.
	}
	fmt.Println("Thank You For Giving Us -",result,"Stars") // Display Result.
}

// Now here what we did is trying to get an Input from user and read it till the user presses Enter.
// After that we are trying to catch any errors if there is any using "comma, ok" syntax.
// After that we are converting the input which is a String type to Number and also removing space by using "strings.TrimSpace()".
/// Packages used -- 
// 1. bufio and os -- to get read functionality
// 2. fmt -- for Println
// 3. strconv -- to convert string into Number
// 4. strings -- to use TrimSpace function to remove the \n that we are getting when user press enter with the input.