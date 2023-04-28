package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//Fetching the argument, and checking for validity
	Arg := os.Args[1]
	mot := splitText(Arg)
	if len(Arg) < 2 {
		return
	}
	for i := 0; i < len(mot); i++ {
		for _, r := range mot[i] {
			if r < 32 || r > 126 {
				fmt.Println("rhooo meecc")
				return
			}
		}
	}

	//Creating array of strings from the standard
	bytes, err := ioutil.ReadFile("../standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(bytes), "\n")

	//Creating the art itself
	var arr []rune
	for i := 0; i < len(mot); i++ {
		for _, r := range mot[i] {
			arr = append(arr, r)
		}
		if mot[i] == "" {
			fmt.Println()
		}else{
			printArt(arr, lines)
		}
		arr = []rune{}
	}
	fmt.Println()
	fmt.Println(splitText(Arg))
}

// Printing given rune array, based on lines art
func printArt(arr []rune, lines []string) {
	for line := 0; line < 8; line++ {
		for _, r := range arr {
			skip := (r - 32) * 9
			fmt.Print(lines[line+int(skip)])
		}
		fmt.Println()
	}
}

func splitText(text string) []string {
	return strings.Split(text, "\\n")

}
