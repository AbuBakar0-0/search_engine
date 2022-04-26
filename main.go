package main

import (
	"errors"
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

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("file does not exist")
	} else {
		fmt.Println("file exists")
	}

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

	for k, v := range tokenMap {
		fmt.Println(k, " VALUE IS ", v)
	}

}
