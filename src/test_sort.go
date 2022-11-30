package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 3, 1, 6, 4}

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println("<: ", arr)

	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println(">: ", arr)

}
