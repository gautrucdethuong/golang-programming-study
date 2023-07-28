package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"

// 	// "go/scanner"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	// "strings"
// )

// func main() {
// 	var personZero, personOne, personTwo string
// 	personZero, personOne, personTwo = "Max", "Alex", "Tom"

// 	fmt.Println("Number of secret person 0: Max, 1: Alex, 2: Tom")

// 	// fmt.Println(rand.Intn(3))
// 	secretPersonNumber := rand.Intn(3)
// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Printf("Please enter your answer with numer 0 - 2: ")
// 	scanner.Scan()
// 	input := scanner.Text()
// 	answer, err := strconv.Atoi(input)

// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	// fmt.Println(out)
// 	if answer == secretPersonNumber {
// 		if secretPersonNumber == 0 {
// 			fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", personZero)
// 		} else if secretPersonNumber == 1 {
// 			fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", personOne)
// 		} else if secretPersonNumber == 2 {
// 			fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", personTwo)
// 		}
// 	} else {
// 		if secretPersonNumber == 0 {
// 			fmt.Printf("Sorry!! You answer is incorrect, 0: %s is the secret person.\n", personZero)
// 		} else if secretPersonNumber == 1 {
// 			fmt.Printf("Sorry!! You answer is incorrect, 1: %s is the secret person.\n", personOne)
// 		} else if secretPersonNumber == 2 {
// 			fmt.Printf("Sorry!! You answer is incorrect, 2: %s is the secret person.\n", personTwo)
// 		}
// 	}
// }
