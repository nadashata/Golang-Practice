// https://www.geeksforgeeks.org/loops-in-go-language/?ref=next_article
package main

import "fmt"

// we only have for as loop in GO
// The opening brace should be in the same line in which post statement exists.

/* If the array, string, slice, or map is empty, 
then for loop does not give an error and continue its flow. 
Or in other words, if the array, string, slice, 
or map is nil then the number of iterations of the for loop is zero.
*/

func main(){

// Infinite Loops it simuilates -> for true{}
for {
	// do some logic
	break		// we can use break keyword to exit loop if certain condition occurred
	// do some logic
	continue	// we can use continue keyword to ignore whatever below and continue loop
	// this logic will be ignored
}

// Normal Loops
// define counter
var i int = 0
i := 0

for i<=3  { // loop exit condition
	// do some logic
	i = i + 1 // increase counter
}

// short hand normal Loops
for i:=0; i<=5;i ++ {
	// do some logic
}

// Loops over a range 
for i := range 6 { // start from 0 till 5
	fmt.Println(i)
}

// when we illustrate lists, slices , we can loop over them too
// assume countires is list

// note to loop over array/slice we need to have 2 variable : index, country
for index, country := countries{
	fmt.Println(country)
	fmt.Println(index)
}

// we can loop over strings
var str string ="Hello nada"
for index,char:= range str{
	fmt.Println(index, char)  // Index -> order in str, Char -> aschi code
} 

// we can loop over strings
var str string ="Hello nada"
for _,char:= range str{
	fmt.Println( char)  // Index -> order in str, Char -> aschi code
}

// we can Loop over array/slices 
// in Go every variable decalred must be used -> what if we don't need Index variable 
for _,country:= countries{
	fmt.Println(country)
}

// we can loop over maps
mmap := map[int]string{ 
	1:"Geeks", 
	23:"GFG", 
	2:"GeeksforGeeks", 
} 
for key, value:= range mmap { 
   fmt.Println(key, value)  
} 

// we can loop over channel 
 chnl := make(chan int) 
 go func(){ 
	 chnl <- 100 
	 chnl <- 1000 
	chnl <- 10000 
	chnl <- 100000 
	close(chnl) 
 }() 
 for i:= range chnl { 
	fmt.Println(i)  
 } 

// we can have nested Loops
for i:=0; i<=2; i++{
	for j:=0;j<=3;j++{
		// Do whatever
	}
}

// If we used break in nested Loop inside inner -> it only break inner Loop
for i:=0;i<=1;i++{
    for j:=0;j<=2;j++{
    	if j == 1{
    	    fmt.Printf("You reached 1, for will stops i=%v, j=%v\n",i,j)
    	    break
    	}
    	fmt.Printf("You are in inner loop i:%v,j:%v\n",i,j)
    }
	    fmt.Printf("You are in outer loop i:%v\n",i)
	}

 
}