package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input, output string
	fmt.Scan(&input)
	for _, n := range input{
		nInt,_ := strconv.Atoi(string(n))
		output += strconv.Itoa(nInt * nInt)
	}
	fmt.Println(output)
}
