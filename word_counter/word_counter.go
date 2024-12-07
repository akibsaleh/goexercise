package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	articleFile := flag.String("article", "demo.txt", "This is a text file")
	flag.Parse()

	if(*articleFile == "") {
		exit("Please provide a text file")
	}

	// fmt.Println(file)
	file, err := os.Open(*articleFile)
	if err != nil {
		exit(fmt.Sprintf("Error opening txt file: %s\n", *articleFile))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			exit(fmt.Sprintf("Error reading txt file: %s\n", *articleFile))
		}
		regEx := regexp.MustCompile("[^a-zA-Z0-9 ]")
		content += regEx.ReplaceAllString(strings.ToLower(scanner.Text()), "")
	}
	contentSlice := strings.Fields(content)
	uniqueWords := make(map[string]int)
	for _, word := range contentSlice {
		uniqueWords[word]++
	}
	wordSlice := make([]string, len(uniqueWords))
	i := 0
	for word:= range uniqueWords {
		wordSlice[i] = word
		i++
	}
	sort.Slice(wordSlice, func (i, j int) bool {
		return uniqueWords[wordSlice[i]] > uniqueWords[wordSlice[j]]
	})
	for idx, word := range wordSlice {
		fmt.Printf("%s: %d\n", word, uniqueWords[word])
		if idx == 10 {
			break
		}
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// Process :
// 1. Create Flags
// 2. handle error if no file is provided
// 3. open file using os.Open
// 4. handle error opening file
// 5. close file with defer
// 6. create scanner with bufio.NewScanner and iterate over each line
// 7. handle error with scanner.Err()
// 8. To remove special characters a regexp is created
// 9. Using.regex.replaceallstring to replace special characters with empty string
// 10. convert to lowercase using strings.ToLower
// 11. split into words using strings.Fields
// 12. create map to store unique words
// 13. iterate over each word and add to map. words are the keys and count is the value
// 14. create slice to store unique words
// 15. iterate over map and add words to slice. if word is repeated, the value increments
// 16. sort slice using sort.Slice to sort words by count
// 17. print top 10 words
