package main

import "fmt"

func main() {

	var a = [10]int{11, -4, 7, 8, -10, 9, -3, 13, -23, 20}
	min, max := findMinAndMax(a)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

}

func findMinAndMax(a [10]int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
