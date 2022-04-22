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

	arrlines := strings.Split(lines, "\n")
	for i := 0; i < len(arrlines); i++ {
		fmt.Println(arrlines[i])
	}

}
