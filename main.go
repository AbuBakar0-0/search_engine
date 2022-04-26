package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	ALBHABET_SIZE = 26
)

type trieNode struct {
	childrens [ALBHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &trieNode{}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
}

func (t *trie) find(word string) bool {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	if current.isWordEnd {
		return true
	}
	return false
}

func main() {

	// Step 1 : Read FILE.txt

	search := "dear"

	_, err := os.Stat("words.txt")

	// if errors.Is(err, os.ErrNotExist) {
	// 	fmt.Println("file does not exist")
	// } else {
	// 	fmt.Println("file exists")
	// }

	content, err := ioutil.ReadFile("words.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := string(content)

	// Step 2 : Convert in Array of lines

	arrlines := strings.Split(lines, "\n")

	// for i := 0; i < len(arrlines); i++ {
	// 	fmt.Println(arrlines[i])
	// }

	//Step 3 : Create LineMap

	linemap := make(map[int][]string)

	for i := 0; i < len(arrlines); i++ {
		linemap[i+1] = strings.Fields(arrlines[i])
	}

	// for k, v := range linemap {
	// 	fmt.Println(k, " VALUE IS ", v)
	// }

	//Step 4 : Create token map

	tokenMap := make(map[string][]int)

	lineNumbers := make([]int, 0)

	for k, v := range linemap {
		for i := 0; i < len(v); i++ {
			if search == v[i] {
				lineNumbers = append(lineNumbers, k)
			}
		}
	}

	tokenMap[search] = lineNumbers

	// for k, v := range tokenMap {
	// 	fmt.Println(k, " VALUE IS ", v)
	// }

	//Step 5: Token Array

	tokens := make([]string, 0)

	for i := 0; i < len(linemap); i++ {
		for j := 0; j < len(linemap[i]); j++ {
			if !contains(tokens, linemap[i][j]) {
				tokens = append(tokens, linemap[i][j])
			}
		}
	}

	// for k, v := range tokens {
	// 	fmt.Println(k, " VALUE IS ", v)
	// }

	//Step 6: Create Trie

	trie := initTrie()

	for i := 0; i < len(tokens); i++ {
		trie.insert(tokens[i])
	}

	//Step 7: Search in Trie

	found := trie.find(search)
	if found {
		fmt.Printf("Word \"%s\" found in trie\n", search)
		fmt.Println("FOUND : ")

		//Step 8: looking for word in token map

		for i := 0; i < len(tokenMap[search]); i++ {
			fmt.Println("LINE :", tokenMap[search][i], " -> ", linemap[tokenMap[search][i]])
		}

	} else {
		fmt.Printf("Word \"%s\" not found in trie\n", search)
	}

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
