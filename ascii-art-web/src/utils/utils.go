package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func HandleError(err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func IsAsciiCaractere(str string) bool {
	for _, r := range str {
		if r < 32 || r > 126 {
			return true
		}
	}
	return false
}

func SplitText(text string) []string {
	return strings.Split(text, "\\n")
}

func OpenAndReadFile(s string) ([]byte, error, string) {
	file := fmt.Sprintf("%s", s)
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
		return nil, nil, "fichier errone"
	}
	return bytes, err, file
}
