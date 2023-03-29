package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Récupérer le nom du fichier à ouvrir en ligne de commande
	filename := os.Args[1]
	outputFile := os.Args[2]

	// Ouvrir le fichier en lecture seule
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()

	//recuperer le contenu du fichier dans une variable
	text := TraitementDeText(filename)

	// Faire quelque chose avec le fichier ouvert...

	// Crée le fichier de sortie
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Erreur en créant le fichier de sortie :", err)
		return
	}
	defer output.Close()

	// Écrire une chaîne de caractères dans le fichier
	_, err = output.WriteString(text)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TraitementDeText(filename string) string {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return "nil"
	}
	// Convertir la phrase en un tableau de mots
	words := strings.Fields(string(contents))

	var upperCount, lowerCount, capCount int

	for i := 1; i < len(words); i++ {
		switch {
		case strings.HasPrefix(words[i], "(up,"):
			m := words[i+1]
			k := m[:len(m)-1]
			n := toInt(k, "k")
			if i-n >= 0 {
				for j := 1; j <= n; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
				i -= n + 1
				upperCount = 0
				lowerCount = 0
				capCount = 0
			} else {
				// If the number of words to uppercase is greater than the available words,
				// uppercase all the available words.
				for j := 1; j <= i; j++ {
					words[i-j] = strings.ToUpper(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
				i -= i
				upperCount = 0
				lowerCount = 0
				capCount = 0
			}
		case strings.HasPrefix(words[i], "(low,"):
			k := words[i+1][:len(words[i+1])-1]
			n := toInt(k, "k")
			if i-n >= 0 {
				for j := 1; j <= n; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
				i -= n + 1
				upperCount = 0
				lowerCount = 0
				capCount = 0
			} else {
				// If the number of words to lowercase is greater than the available words,
				// lowercase all the available words.
				for j := 1; j <= i; j++ {
					words[i-j] = strings.ToLower(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
				i -= i
				upperCount = 0
				lowerCount = 0
				capCount = 0
			}
		case strings.HasPrefix(words[i], "(cap,"):
			k := words[i+1][:len(words[i+1])-1]
			n := toInt(k, "k")
			if i-n >= 0 {
				for j := 1; j <= n; j++ {
					words[i-j] = Capitalize(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
				i -= n + 1
				upperCount = 0
				lowerCount = 0
				capCount = 0
			} else {
				// If the number of words to capitalize is greater than the available words,
				// capitalize all the available words.
				for j := 1; j <= i; j++ {
					words[i-j] = Capitalize(words[i-j])
				}
				words = append(words[:i], words[i+2:]...)
				i -= i
				upperCount = 0
				lowerCount = 0
				capCount = 0
			}
		case words[i] == "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
			upperCount = 1
			lowerCount = 0
			capCount = 0
		case strings.Contains(words[i], "(up)"):
			words[i] = strings.ToUpper(words[i])
		case words[i] == "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
			upperCount = 0
			lowerCount = 1
			capCount = 0
		case strings.Contains(words[i], "(low)"):
			words[i] = strings.ToLower(words[i])
		case words[i] == "(cap)":
			words[i-1] = Capitalize(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
			upperCount = 0
			lowerCount = 0
			capCount = 1
		case strings.Contains(words[i], "(cap)"):
			words[i] = Capitalize(words[i])
		case words[i] == "a" && i < len(words)-1:
			nextWord := words[i+1]
			if isVowel(nextWord[0]) {
				words[i] = "an"
			}
		case words[i] == "(hex)":
			decimal, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Println(err)
				return "nil"
			}
			decimalFormatted := strconv.FormatInt(decimal, 10)
			words[i-1] = decimalFormatted
			words = append(words[:i], words[i+1:]...)
			i--
		case words[i] == "(bin)":
			decimal, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println(err)
				return "nil"
			}
			decimalFormatted := strconv.FormatInt(decimal, 10)
			words[i-1] = decimalFormatted
			words = append(words[:i], words[i+1:]...)
			i--
		default:
			if upperCount > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				upperCount--
			} else if lowerCount > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				lowerCount--
			} else if capCount > 0 {
				words[i-1] = Capitalize(words[i-1])
				capCount--
			}
		}
	}

	output := strings.Join(words, " ")

	output = Apostrophe(output)

	output = strings.ReplaceAll(output, "(up)", "")
	output = strings.ReplaceAll(output, "(UP)", "")
	output = strings.ReplaceAll(output, "(LOW)", "")
	output = strings.ReplaceAll(output, "(low)", "")
	output = strings.ReplaceAll(output, "(Cap)", "")
	output = strings.ReplaceAll(output, "(cap)", "")
	output = replaceMultipleSpaces(output)
	output = Ponctuation(output)
	return output

}

func replaceMultipleSpaces(s string) string {
	words := strings.Fields(s)
	return strings.Join(words, " ")
}

func toInt(s, name string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Invalid %s: %s", name, s)
	}
	return i
}

func isVowel(c byte) bool {
	vowels := "aeihouAEIHOU"
	for i := 0; i < len(vowels); i++ {
		if c == vowels[i] {
			return true
		}
	}
	return false
}

func Ponctuation(s string) string {
	// Définit la regex pour trouver les ponctuations.
	re := regexp.MustCompile(`[\,\.\!\?\:\;]+`)
	// Applique la regex pour trouver toutes les ponctuations.
	matches := re.FindAllString(s, -1)
	// Remplace chaque ponctuation avec elle-même avec une espace à gauche et supprime les espaces inutiles.
	for _, match := range matches {
		s = strings.ReplaceAll(s, match, match+" ")
		// Supprime l'espace après la virgule.
		s = strings.ReplaceAll(s, " "+match, match)
	}
	return replaceMultipleSpaces(s)
}

func Apostrophe(s string) string {
	re := regexp.MustCompile(`'\s*(\w+)\s*'`)
	s = string(re.ReplaceAll([]byte(s), []byte("'$1'")))
	return s
}

func Capitalize(s string) string {
	tab := []rune(s)
	if len(tab) > 0 {
		if tab[0] != ' ' && tab[0] >= 'a' && tab[0] <= 'z' {
			tab[0] = tab[0] - 32
		}
	}
	for i := 0; i < len(tab)-1; i++ {
		if (tab[i] < 'a' || tab[i] > 'z') && (tab[i] < 'A' || tab[i] > 'Z') && (tab[i] < '0' || tab[i] > '9') {
			if tab[i+1] >= 'a' && tab[i+1] <= 'z' {
				tab[i+1] = tab[i+1] - 32
			}
		} else if tab[i+1] >= 'A' && tab[i+1] <= 'Z' {
			tab[i+1] = tab[i+1] + 32
		}
	}
	return string(tab)
}
