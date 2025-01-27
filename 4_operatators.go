// https://www.geeksforgeeks.org/data-types-in-go/?ref=next_article
// https://www.geeksforgeeks.org/go-operators/?ref=next_article

package main
import (
	// this package help us to use print func
    "fmt"
	// this package help us to get type of variable
    "reflect" 						
    )

// in Basic Data Types we have Numbers, string, bool

// In Numbers 3 categories: int, float, complex
	// float 
		//decimal (3.15)
		//exponential ( 12e18 or 3E10)
		//mixed (13.16e12)

// we can do below arthimetic operations for Int and all except reminder for float
// only % cannot be used with float

	// addition
	// substraction
	// reminder
	// division
	// multiply

// Result Type
	// int {operation} int = int
	// int {operation} float = float
	// float {operation} float = float

func main(){

  sum := 3.0+2  
  fmt.Println((reflect.TypeOf(sum)) // Float

  reminder := 3.0%2  
  fmt.Println((reflect.TypeOf(reminder)) // Error Invalid Opertation

  reminder2 := 3%2  
  fmt.Println((reflect.TypeOf(reminder2)) // int

  // we can concate strings using +
  string1 := "Hi "
  string2 := "Nada"
  fmt.Println(string1 + string2) 
  
}
