package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello from Go!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter a text")
	input , _ := reader.ReadString('\n')

	fmt.Println(input)

	fmt.Println("enter a number")

	numInput , _ := reader.ReadString('\n')

	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	if err != nil {
			fmt.Println(err)
		}	 

		fmt.Println(aFloat)

}
