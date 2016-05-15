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
			continue
		}

		fmt.Println("Found Node:", string(node.Path))

		fmt.Print("Text to Search: ")
	}
}

func exampleData(radix *goradix.Radix) {
	radix.Insert("test")
	radix.Insert("toaster")
	radix.Insert("toasting")
	radix.Insert("slow")
	radix.Insert("slowly")
}
