package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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

	for k, v := range tokens {
		fmt.Println(k, " VALUE IS ", v)
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
