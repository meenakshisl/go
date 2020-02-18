package main

import "fmt"

func main() {
	n:=5
//if
	if n==5 {
		fmt.Println("The number is ", n)
	}
//if-else
	if n%2 == 0 {
		fmt.Println("Even")
	} else	{
		fmt.Println("Odd")
	}
//if else if

	if num:=9;n>0 { 	//local variable available throughout the if-else structure
		fmt.Println("Is Positive")
	}else if num < 0 {
		fmt.Println("Is Negative")
	}else {
		fmt.Println("Is zero")
	}
}


//there is no ternary if in go
