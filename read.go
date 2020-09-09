package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type name struct {
	firstname string
	lastname  string
}

func main() {

	mySlice := make([]name, 0)

	var file string

	fmt.Println("enter the filename")
	fmt.Scan(&file)
	file, err := os.Open(file)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	

	for _, eachline := range txtlines {
		s := eachline.splt(" ")
		var tmp name
		tmp.firstname = s[0]
		tmp.lastname = s[1]
		mySlice = append(mySlice, tmp)
	}
	fmt.Println(mySlice)

}
