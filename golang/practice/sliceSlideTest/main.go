package main

import "fmt"

func main() {
	// test for queue
	fmt.Println("==== queue ====")
	sl := make([]int, 1, 5)
	fmt.Printf("%v, len : %d, cap : %d\n", sl, len(sl), cap(sl))

	sl = append(sl, 1)
	sl = append(sl, 2)
	sl = append(sl, 3)

	fmt.Printf("%v, len : %d, cap : %d\n", sl, len(sl), cap(sl))

	temp := sl[0]
	// capacity reduced
	sl = sl[1:]

	fmt.Printf("temp : %d\n", temp)
	fmt.Printf("%v, len : %d, cap : %d\n", sl, len(sl), cap(sl))

	fmt.Println("==== stack ====")
	sl2 := make([]int, 1, 5)
	fmt.Printf("%v, len : %d, cap : %d\n", sl2, len(sl2), cap(sl2))

	sl2 = append(sl2, 1)
	sl2 = append(sl2, 2)
	sl2 = append(sl2, 3)

	fmt.Printf("%v, len : %d, cap : %d\n", sl2, len(sl2), cap(sl2))

	index := len(sl2) - 1
	temp2 := sl2[index]
	// capacity maintained
	sl2 = sl2[:index]

	fmt.Printf("temp : %d\n", temp2)
	fmt.Printf("%v, len : %d, cap : %d\n", sl2, len(sl2), cap(sl2))
}
