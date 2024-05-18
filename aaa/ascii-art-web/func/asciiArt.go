package asciiArt

import (
	"fmt"
	"os"
)

func generate(s string, banner string) string {
	input := s
	for i := 0; i < len(input); i++ {
		if input[i] < 32 || input[i] > 128 {
			return "error"
		}
	}
	word := split(input)
	BANNER := banner
	fileContent, err := os.ReadFile(BANNER)
	if err != nil {

		return "error in file\n"
	}
	if fileContent == nil {
		fmt.Print("no Text")
	}
	lettres := getLettres(fileContent)
	return output(lettres, word)

}
func output(lettres [][]string, word []string) string {
	output := ""
	bl := false
	for l := 0; l < len(word); l++ {
		if word[l] == "" {
			continue
		}
		if word[l] == "\n" {
			if l != len(word)-1 && bl && word[l+1] != "\n" {
				continue
			}
			output += "\n"
			continue
		}
		for i := 1; i < 9; i++ {
			for j := 0; j < len(word[l]); j++ {
				output += lettres[word[l][j]-32][i]
			}
			output += "\n"
		}
		bl = true
	}
	return output
}

func split(str string) []string {
	word := ""
	splitedword := []string{}
	skip := false
	for i := 0; i < len(str); i++ {
		if skip {
			skip = false
			continue
		}
		if i != len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			if word != "" {
				splitedword = append(splitedword, word)
			}
			word = ""
			skip = true
			splitedword = append(splitedword, "\n")
			continue
		}
		word = word + string(str[i])
	}
	if word != "" {
		splitedword = append(splitedword, word)
	}
	return splitedword
}
func getLettres(fileContent []byte) [][]string {
	Content := []byte{}
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] != 13 {
			Content = append(Content, fileContent[i])
		}
	}
	fileContent = Content
	lettres := [][]string{}
	lettre := []string{}
	line := []byte{}
	for i := 0; i < len(fileContent); i++ {
		if i != len(fileContent)-1 && fileContent[i] == '\n' && fileContent[i+1] == '\n' {
			lettre = append(lettre, string(line))
			lettres = append(lettres, lettre)
			lettre = nil
			line = nil
			continue
		}
		if fileContent[i] == '\n' {
			lettre = append(lettre, string(line))
			line = nil
			continue
		}
		line = append(line, fileContent[i])
	}
	lettres = append(lettres, lettre)
	return lettres
}
