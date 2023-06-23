package src

import (
	"strings"
	utils "AsciiArtWeb/src/utils"
)

func CreatingArtStandard(str1 string) string {
	mot := utils.SplitText(str1)
	output := ""
	if !utils.IsAsciiCaractere(str1) {
		bytes, err, _ := utils.OpenAndReadFile("fonts/standard.txt")
		utils.HandleError(err)
		lines := strings.Split(string(bytes), "\n")
		//Creating the art itself
		var arr []rune
		for i := 0; i < len(mot); i++ {
			if mot[i] != "" {
				for _, r := range mot[i] {
					arr = append(arr, r)
				}
				if mot[i] == "" {
					output += "\n"
				} else {
					output += printArtStandard(arr, lines)
				}
				arr = []rune{}
			} else if i < len(mot)-1 {
				output += "\n"
			}
		}
	} else {
		return "it's not a ascii caractere"
	}
	return output
}

// Printing given rune array, based on lines art
func printArtStandard(arr []rune, lines []string) string {
	output := ""
	for line := 0; line < 8; line++ {
		for _, r := range arr {
			skip := (r - 32) * 9
			output += (lines[line+int(skip)])
		}
		output += "\n"
	}
	return output
}