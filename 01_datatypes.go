package main

import "fmt"

const Loggedin bool = true // This is a public variable which is constant as it is starting with Uppercase L.

func main( ){
  
  var a string = "Avinash"
  fmt.Println("variable:",a)
  
  var isLoggedIn bool = true
  fmt.Println("Boolean Value is :", isLoggedIn)

  var aNum int = 2578 // Here you have multiple int type but for general purpose we use int.
  fmt.Println("Integer Value is :",aNum)

  var smallFloat float32 = 256.455 // float is of 2 types - float32 and float64. float64 will give more values after the decimal.
  fmt.Println("Float value is :", smallFloat)

  // When you dont want to declare a variable with var than use walrus operator ":=" and put the value.
  // This will only work inside the function but if try to do it in global scope it will not work.
  noVar := 55000
  fmt.Println(noVar)

}

// here this will give error-
// newName := "Avinash"
// fmt.Println(newName)
