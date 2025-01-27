// refer to this link for Identifiers conditions
// https://www.geeksforgeeks.org/identifiers-in-go-language/
package main
import "fmt"

// Go Compilier is so Powerfull
	// we cannot declare a variable and not use it -> Error before runtime
	// we cannot redeclare a variable twice, but ofcourse we can assign a new value to it , if it is from same type

// Three Types of variables
// 1) Package Variable -> Outside any function or Block {if/else or for}
	// can be accessed from any file with in main package
	// can be declared using var, const keywords only

var a int							// only declare a package variable 
const x string = "Hello World"		// declare + assiging a value to package variable


func main(){

// Local variables,  2) functions variables or 3) blocks variables
// it is the best Practise to always use as local variables as possible

// 2) Functions variables
	var b string					// declare string
	var d int 						// delcare int
	var e bool 						// declare bool

	// for variable declared and not assigned a value they are "zero-typed" until the user assign a value for them 
	// there is no unintialized variables concept in go
		// b = ""
		// d = 0
		// e = false

// declare + assiging 
	var f int = 12						

// declare + assiging -> without explicitly setting type , go will infer type by itself from value assigned to variable
	var f = 12	
// the short hand format below, only available in local variables declaraton, cannot be used in const, cannot be used in package lvl variables
	// short hand for declare is to use ":" instead of "var" + assigin so Go can infer var type	
	g := "Hello"					

// we can declare and assign var for multiple var in same line
	var x, y,z int

	var x, y,z int = 1,2,3
	var x,y,z = 1, "hello", true

	x,y,z := 1,2,3
	x,y,z := 1,"hello", false

// we can declare variables and set their value from return of function
	var x,y = funct_name()
	x,y := funct_name()
// we cannot declare var twice -> Error other declaration of x or y
	y := 1
	y := 2

	var x int
	var x string

// but we can change value of variable after declaration -> condition new value of same type
	z :=1
	z =2

// we can change value of variable after declaration -> condition of same type
	z :=1
	z = 2

// But we cannot change value of variable with a different types
	// we cannot assign string to int
	z :=1
	z = 2
	z="hello" 	
	// we cannot assign uint8 to uint16
	var z uint16 =5
	var x uint8 =4
	z = x				
	// we cannot assign uint8 to uint
	var z uint =5
	var x uint8 =4
	z =x				

// 3) block variables
	//block variable , once block scope finished , it is destroyed, cannot be accessibile from outside block
	if x>9 {
		var sum = 12 // block var
	}

// Define constant
	const const_var = "Hi"

// if we don't know type of variable expected from user 

} 