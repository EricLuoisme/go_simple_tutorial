package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type FullName struct {
	fname string
	lname string
}

func main() {

	var filePath string
	storeSlice := make([]FullName, 0)

	fmt.Println("Please input the file name")
	_, err := fmt.Scan(&filePath)
	if err != nil {
		return
	}

	file, err2 := os.Open(filePath)
	if err2 != nil {
		fmt.Println("File Reading Error")
		return
	}

	br := bufio.NewReader(file)
	for {
		// read single line
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}

		// split name & construct struct
		words := strings.Split(string(line), " ")
		p := FullName{fname: words[0], lname: words[1]}

		// using fmt.Sscanf(scanner.Text, "%s %s", &n.fname, &n.lname) also working

		// append
		storeSlice = append(storeSlice, p)
	}

	// traverse
	for _, person := range storeSlice {
		fmt.Printf("Fist Name:%s, Last Name:%s \n", person.fname, person.lname)
	}

}
