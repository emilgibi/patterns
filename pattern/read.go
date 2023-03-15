package pattern

import (
    "fmt"
    "strings"
    "bufio"
    "strconv"
    "os"
)

func Read() {
    reader := bufio.NewReader(os.Stdin)
     fmt.Print("Text to send: ")
     text, _ := reader.ReadString('\n')
     newText := strings.TrimSpace(text)
     fmt.Print(strconv.Itoa(len(newText))) 
    
}

func CountSpecials() {
	cons := 0
	vowels := 0
	specialCharacters := 0

	read := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	string, _ := read.ReadString('\n')

	for _, ch := range string {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			if ch != 'a' && ch != 'e' && ch != 'i' && ch != 'o' && ch != 'u' && ch != 'A' && ch != 'E' && ch != 'I' && ch != 'O' && ch != 'U' {
				cons++
			} else {
				vowels++
			}
		} else {
			specialCharacters++
		}
	}

	fmt.Printf("Number of consonants: %d\n", cons)
	fmt.Printf("Number of vowels: %d\n", vowels)
	fmt.Printf("Number of special characters: %d\n", specialCharacters)
}

func LowerCase() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Enter a string: ")
	string, _ := reader.ReadString('\n')

	lowerCase := strings.ToLower(string)

	fmt.Println("Lowercase string:", lowerCase)
}

func SentencetoWords() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	sentence, _ := reader.ReadString('\n')

	words := strings.Fields(sentence)

	Count := len(words)

	fmt.Printf("Number of words: %d\n", Count)
}
