package src

import (
	utils "AsciiArtWeb/src/utils"
	"strings"
)

// CreatingArt creates an art based on file passed in arguments this function work when we have two arguments
func CreatingArt(str1, str2 string) string {
	mot := utils.SplitText(str1)
	output := ""
	if !utils.IsAsciiCaractere(str1) {
		bytes, err, _ := utils.OpenAndReadFile("fonts/thinkertoy.txt")
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
					output += printArt(arr, lines)
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
func printArt(arr []rune, lines []string) string {
	output := ""
	for line := 1; line <= 8; line++ {
		for _, r := range arr {
			skip := (r - 32) * 9
			output += (lines[line+int(skip)])
		}
		output += "\n"
	}
	return output
}



// func CreateANewFile(fileName string) (string, error) {
// 	// Create a new file for output
// 	f, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}
// 	defer f.Close()
// 	return fileName, err
// }

// func writeOutputToFile(outputFile, output string) {
// 	err := ioutil.WriteFile(outputFile, []byte(output), 0644)
// 	HandleError(err)
// 	// fmt.Println("Output file written successfully!")
// }

// func GetOutputFilename(arg string) string {
// 	fileName := ""
// 	if strings.HasPrefix(arg, "--output=") {
// 		fileName = strings.Split(arg, "=")[1]
// 	}
// 	return fileName
// }
