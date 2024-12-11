package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		input string
		max int
	)
	fmt.Scan(&input)
	for _,el := range input{
		intEl,_ := strconv.Atoi(string(el))
		if intEl > max{
			max = intEl
		}
	}
	fmt.Println(max)
}
