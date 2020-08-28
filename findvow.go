package main

import "fmt"

type groupChars []rune

func isVowel(c rune) bool {
	vowels := groupChars{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
	for _, value := range vowels {
		if value == c {
			return true
		}
	}
	return false
}

func main() {
	myString := "sezzle India"
	a := 0
	for _, value := range myString {
		if isVowel(value) {
			fmt.Printf("%c is Vowel\n", value)
			a++
		}
	}
	fmt.Printf("%d Vowels.", a)

}
