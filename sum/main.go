package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")
    var number1 float64
    fmt.Scanf("%f", &number1)

	fmt.Print("Enter a second number: ")
    var number2 float64
    fmt.Scanf("%f", &number2)

    output := number1 + number2

    fmt.Println(number1,"+", number2, "=", output)    
}