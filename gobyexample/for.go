package main

import "fmt"

func main() {
	i:=1
	for i<=5 {			//for loop with one condition; similiar to while
		fmt.Println(i)
		i=i+3
	}
	for {				//endless loop == while(True)
		fmt.Println("Endless loop bout to break")
		break
	}
	for j:=7;j<=10;j++ { 		//normal while loop
		fmt.Println(j)
	}
	for n:=0;n<=9;n=n+2 {
		if(n==3) {
			continue
		}
		fmt.Println(n)
	}
}
