package main

import("fmt"
	"time"
)

func main() {
	i:=2

	//normal switch
	switch i {
	case 1 :
		fmt.Println("One")
	case 2 :
		fmt.Println("Two")
	case 3 :
		fmt.Println("Three")
	}

	//swtich with multiple expressions

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday :
		fmt.Println("Weekends")
	default :
		fmt.Println("Week day")
	}

	t:=time.Now()
	switch  {
		case t.Hour() < 12 :
			fmt.Println("Before Noon")
		default :
			fmt.Println("After Noon")
		}

	whatAmI := func( i interface{}) {
		switch t:=i.(type) {
		case bool :
			fmt.Println("Its a bool")
		case int :
			fmt.Println("Its a int ")
		default :
			fmt.Println("Recognising.....",t)
		}
	}

	whatAmI(true)
	whatAmI(8)
	whatAmI(9.78)
	whatAmI("hello")
}
