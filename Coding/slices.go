/*
Slice được xây dựng dựa trên khái niệm của dynamic arrays,
là một array có chiều dài phụ thuộc vào số lượng thành phần mà ta truyền vào
và chiều dài của nó có thể tăng và giảm tùy ý được.
**/

package main

// import "fmt"

// func main() {
// 	var animals []string
// 	// nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type.
// 	fmt.Println("uninit:", animals, animals == nil, len(animals) == 0)

// 	// The cap built-in function returns the capacity of v
// 	animals = make([]string, 3)
// 	fmt.Println("emp:", animals, "len:", len(animals), "cap:", cap(animals))
// 	// set value
// 	animals[0] = "tiger"
// 	animals[1] = "fish"
// 	animals[2] = "chicken"

// 	fmt.Println("set:", animals)
// 	fmt.Println("get element index 1 is:", animals[1])
// 	fmt.Println("Length is", len(animals))

// 	animals = append(animals, "bird")
// 	animals = append(animals, "pig")
// 	fmt.Println(animals) //get new list

// 	// copy array in tmp
// 	tmp := make([]string, len(animals))
// 	copy(tmp, animals)
// 	fmt.Println("list tmp:", tmp)

// 	// sclice array by index
// 	new_array := tmp[2:len(animals)]
// 	fmt.Println("new array:", new_array)

// 	new_array = tmp[:5]
// 	fmt.Println("list tmp[:5]:", new_array)

// 	new_array = tmp[2:]
// 	fmt.Println("list tmp[:2]:", new_array)

// 	//array list hiragana with set default value
// 	hiragana := []string{"a", "i", "u", "e", "o"}
// 	fmt.Println("list hiragana:", hiragana)

// 	twoD := make([][]int, 3)
// 	for i := 0; i < 3; i++ {
// 		innerLen := i + 1
// 		twoD[i] = make([]int, innerLen)
// 		for j := 0; j < innerLen; j++ {
// 			twoD[i][j] = i + j
// 		}
// 	}
// 	fmt.Println("2d:", twoD)
// }
