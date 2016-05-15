package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/falmar/goradix"
)

func main() {

	radix := goradix.New()

	exampleData(radix)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Text to Search: ")
	for scanner.Scan() {

		textToLook := scanner.Text()

		node, err := radix.LookUp([]byte(textToLook))

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("\nFound Node:", string(node.Path))
		}

		fmt.Print("Text to Search: ")
	}
}

func exampleData(radix *goradix.Radix) {
	text := []string{"test", "toaster", "toasting", "slow", "slowly"}
	for _, v := range text {
		fmt.Println("Inserted:", v)
		radix.Insert(v)
	}
}
