package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	mySlice1 := make([]int, 3)

	fmt.Println("enter an integer")
	var x string
	fmt.Scan(&x)
	for x != "X" {

		i, err := strconv.Atoi(x)
		if err == nil {
			mySlice1 = append(mySlice1, i)
			sort.Ints(mySlice1)
			fmt.Println(mySlice1)
			fmt.Println("enter an integer")
			fmt.Scan(&x)
		}

	}

}
