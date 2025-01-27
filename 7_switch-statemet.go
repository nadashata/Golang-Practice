package main

// to Import more than 1 package we use () ans seprate each one of them in seperate line

import (
	"fmt"
	"time"
)

func main(){

	x := "London"
	// x is called switch expression
	switch x {   					
		case "Cairo":
			// do something
		case "London","Tornonto":	// we can have use "," to seperate multiple expressions 
			// do something
		default:
			// do something
	}

	// we can have switch without switch expression -> act as if/else

	switch {   					
		case time.Hour() > 12:
			// do something
		default:
			// do something
	}
	
}
