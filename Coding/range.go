package main

// import (
// 	"fmt"
// )

// func main() {

// 	// nums := make([]int, 5)
// 	// nums[1] = 10
// 	// nums[2] = 20
// 	nums := []int{2, 3, 4}
// 	sum := 0
// 	// var sum int = 0

// 	/*range on arrays and slices provides both the index and value for each entry.
// 	Above we didn’t need the index, so we ignored it with the blank identifier _.
// 	Sometimes we actually want the indexes though.*/
// 	for _, num := range nums { // not use index, use _ define null
// 		sum += num
// 	}
// 	fmt.Println("sum of number:", sum)

// 	for i, num := range nums { // get by index i
// 		if num == 4 {
// 			fmt.Println("index:", i)
// 		}
// 	}
// 	fmt.Println("-----------------------")
// 	// using map
// 	animals_map := map[string]string{"coop1": "fish", "coop2": "pig", "coop3": "chicken"}
// 	for key, value := range animals_map {
// 		fmt.Printf("%s -> %s\n", key, value)
// 		fmt.Println("key:", key) //get key
// 	}

// 	// for i, cap := range "go" {
// 	// 	fmt.Println(i, cap)
// 	// }
// }
