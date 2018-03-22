package main

import (
  "fmt"
  "strings"
)

func main(){

  var (
      valueOne, valueTwo int
      userInput string
      finalValue int
  )

  //Takes in value 1 int
  fmt.Println("Enter your first value: ")
  fmt.Scanf("%d", &valueOne)

  //takes in value 2 int
  fmt.Println("Enter your second value: ")
  fmt.Scanf("%d", &valueTwo)

  //function to print options out
  printOptions();

  // takes in the string of math application and converts it to lowercase
  fmt.Println("Type the math function: ")
  fmt.Scanf("%s", &userInput)

  strings.ToLower(userInput)

  //switch to run the correct function needed
  switch  userInput{
  case "addition":
    finalValue = addition(valueOne, valueTwo)
    fmt.Println("The value is:", finalValue)

  case "subtraction":
    finalValue = subtraction(valueOne, valueTwo)
    fmt.Println("The value is:", finalValue)

  case "multiplication":
    finalValue = multiplication(valueOne, valueTwo)
    fmt.Println("The value is:", finalValue)

  case "division":
    finalValue = division(valueOne, valueTwo)
    fmt.Println("The value is:", finalValue)

  default:
    fmt.Println("Option Not avaible or misstyped")

  }



}

func printOptions (){
  options := `  addition
  Subtraction
  Multiplication
  Division`

  fmt.Println(options)
}

func addition (x int, y int) int {
  return x + y
}

func subtraction (x int, y int) int {
  return x - y
}

func multiplication (x int, y int) int {
  return x*y
}

func division (x int, y int) int  {
   if y != 0 {
     return x/y
   } else {
     fmt.Println("Can't divide by Zero!")
     return 0
   }

}
