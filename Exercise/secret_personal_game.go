package main

import (
	"fmt"
	"math/rand"
)

func randomStepUser() int {
	return rand.Intn(3)
}

func main() {
	// var personZero, personOne, personTwo string
	// personZero, personOne, personTwo = "Max", "Alex", "Tom"
	persons := [3]string{"Max", "Alex", "Tom"}

	fmt.Println("Number of secret person 0: Max, 1: Alex, 2: Tom")

	// fmt.Println(rand.Intn(3))
	secretPersonNumber := randomStepUser()
	var input int
	// scanner := bufio.NewScanner(os.Stdin)

	for { // Infinite Loop = While
		fmt.Printf("Please enter your answer with numer 0 - 2: ")
		fmt.Scan(&input)

		if input >= 0 && input <= 2 {
			break
		}
	}
	// scanner.Scan()
	// input := scanner.Text()
	// answer, err := strconv.Atoi(input)

	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println(out)
	if input == secretPersonNumber {
		fmt.Printf("Congratulations!! You answer is correct, %s is the secret person.\n", persons[input])
	} else {
		fmt.Printf("Sorry!! You answer is incorrect, %d is the secret person.\n", secretPersonNumber)
	}
}
