package main

import "fmt"

func removeduplicatevalues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {

	intSliceValues := []int{1, 2, 3, 4, 5, 2, 3, 5, 7, 9, 6, 7}

	fmt.Println(intSliceValues)

	removeduplicatevaluesSlice := removeduplicatevalues(intSliceValues)

	fmt.Println(removeduplicatevaluesSlice)
}
