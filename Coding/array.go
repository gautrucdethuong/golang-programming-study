package main

// import "fmt"

// /*
// By default an array is zero-valued, which for ints means 0s.

// In Go, an array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.
// */
// func main() {
// 	fmt.Println("a ...any")

// 	var numbers [5]int
// 	fmt.Println("Arrays number:", numbers)

// 	numbers[4] = 100
// 	fmt.Println("set value element 5: :", numbers)
// 	fmt.Println("get num last index:", numbers[4])

// 	fmt.Println("length:", len(numbers))
// 	// fmt.Println("get element first:")

// 	b := [5]int{1, 2, 3, 4, 5}
// 	fmt.Println("dcl:", b)

// 	var twoD [2][3]int
// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 3; j++ {
// 			twoD[i][j] = i + j
// 		}
// 	}
// 	fmt.Println("2d: ", twoD)
// }
