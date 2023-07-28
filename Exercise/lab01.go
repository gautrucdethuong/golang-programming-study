package main

// import (
// 	"flag"
// 	"fmt"
// )

// // execute binary file : go build -o app

// var koreanGreeting string

// func main() {
// 	// fmt.Println("Konnichiwa")

// 	// declare variable
// 	/*1. var <name> <type>
// 	2. var <name> = <value>
// 	3. <name> := <value>
// 	*/
// 	language := flag.String("lang", "en", "greeting language")
// 	flag.Parse()

// 	koreanGreeting = "안녕, 친구"
// 	var enGreeting = "Hello"
// 	chinaGreeting := "你好朋友"

// 	if *language == "en" {
// 		fmt.Println(enGreeting)
// 	} else if *language == "korean" {
// 		fmt.Println(koreanGreeting)
// 	} else if *language == "china" {
// 		fmt.Println(chinaGreeting)
// 	} else {
// 		fmt.Println("Language is not support")
// 	}
// }
