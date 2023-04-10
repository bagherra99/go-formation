package main

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Ouvrir le fichier texte
	file, err := os.Open("standard.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Lire le contenu du fichier "monfichier.txt"
	contenu, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	var tableau []string

	parts := strings.Split(string(contenu), "\n")
	for i := 0; i < len(parts); i += 9 {
		end := i + 9
		if end > len(parts) {
			end = len(parts)
		}
		ascii := strings.Join(parts[i:end], "\n")
		tableau = append(tableau, ascii)

	}

	a := tableau[40]
	b := tableau[41]
	c := strings.Split(a, "\n")
	d := strings.Split(b, "\n")

	for i := 0; i < 6; i++ {
		// fmt.Println(string(c[0][0]))
	}
	fmt.Println(c[1]+d[1])
	fmt.Println(len(tableau))

	
}

// for j := 0; j < len(tableau); j++ {
	// 	a := tableau[j]
	// 	c := strings.Split(a, "\n")
	// 	for i := 0; i < 8; i++ {
	// 		fmt.Println(c[i])
	// 		// fmt.Println(d[i])
	// 	}
	// }