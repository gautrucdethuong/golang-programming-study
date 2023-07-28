package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func randomStepUser(user string) int {
// 	step := rand.Intn(10)
// 	fmt.Printf("User %s go %d step\n", user, step)
// 	return step
// }

// func main() {
// 	// var numberPlayer int
// 	// var nameSingleUser string
// 	// var nameOfUsers = make([]string, 2, 6)
// 	// var test = []string{"a", "i", "u", "e", "o"}
// 	// fmt.Printf("Enter number players:")
// 	// fmt.Scan(&numberPlayer)
// 	// fmt.Println(nameOfUsers, len(nameOfUsers), cap(nameOfUsers))
// 	// for i := 1; i <= numberPlayer; i++ {
// 	// 	fmt.Printf("Enter name for user %d: ", i)
// 	// 	fmt.Scan(&nameSingleUser)
// 	// 	nameOfUsers = append(nameOfUsers, nameSingleUser)
// 	// }

// 	// fmt.Println(nameOfUsers, len(nameOfUsers), cap(nameOfUsers))
// 	// fmt.Println(test, len(test), cap(test))

// 	var numberOfStepUserOne, numberOfStepUserTwo int
// 	var UserOneName, UserTwoName string
// 	fmt.Println("===== WELCOME GAME RUNNER WITH TWO USER =====")
// 	fmt.Print("Enter name for user one: ")
// 	fmt.Scan(&UserOneName)
// 	fmt.Print("Enter name for user two: ")
// 	fmt.Scan(&UserTwoName)

// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("========")
// 		fmt.Printf("Ready for round %d ...\n", i)
// 		numberOfStepUserOne += randomStepUser(UserOneName)
// 		numberOfStepUserTwo += randomStepUser(UserTwoName)

// 		time.Sleep(1000 * time.Millisecond)
// 	}
// 	fmt.Println("==========RESULT============")
// 	if numberOfStepUserOne > numberOfStepUserTwo {
// 		fmt.Printf("User %s winner with of step is %d\n", UserOneName, numberOfStepUserOne)
// 	} else {
// 		fmt.Printf("User %s winner with of step is %d\n", UserTwoName, numberOfStepUserTwo)
// 	}
// }
