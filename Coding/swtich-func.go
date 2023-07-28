package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("hello word")

// 	i := 2
// 	fmt.Println("Write", i, "as")
// 	switch i {
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("two")
// 	case 3:
// 		fmt.Println("three")
// 	}

// 	switch time.Now().Weekday() {
// 	case time.Thursday, time.Saturday:
// 		fmt.Println("It's the weekend")
// 	default:
// 		fmt.Println("It's a weekday")
// 	}

// 	time := time.Now()
// 	switch {
// 	case time.Hour() < 12:
// 		fmt.Println(time)
// 		fmt.Println("It's before noon")
// 	default:
// 		fmt.Println("It's after noon")
// 	}
// 	// func
// 	fmt.Println("----------Func------------")
// 	whatAmI := func(i interface{}) {
// 		switch t := i.(type) {
// 		case bool:
// 			fmt.Println("I am bool")
// 		case int:
// 			fmt.Println("I am int")
// 		default:
// 			// fmt.Println("Don't know type %T\n", t)
// 			fmt.Printf("Don't know type %T\n", t)
// 		}
// 	}
// 	whatAmI(true)
// 	whatAmI(2)
// 	whatAmI("hey")
// }