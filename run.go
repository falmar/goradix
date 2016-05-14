package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	radix := Radix{master: true}

	insertToRadix(&radix)

	printRecursive(&radix, 0)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Text to Search: ")
	for scanner.Scan() {

		textToLook := scanner.Text()

		_, wordSlice, err := radix.Autocomplete(textToLook)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(wordSlice)
		fmt.Print("Text to Search: ")
	}
}

func insertToRadix(radix *Radix) {

	// Trivial Example 1:
	radix.Insert("romane")
	radix.Insert("romanus")
	radix.Insert("romulus")
	radix.Insert("rubens")
	radix.Insert("ruber")
	radix.Insert("rubicon")
	radix.Insert("rubicundus")
	radix.Insert("rinicundus")
	radix.Insert("repicundus")
	radix.Insert("lepicundus")
	radix.Insert("lepocundus")
	radix.Insert("lomulus")
	radix.Insert("lupus")
	radix.Insert("huber")
	radix.Insert("pepicundus")
	radix.Insert("pepicundas")
	radix.Insert("pepocundus")
	radix.Insert("pomulus")
	radix.Insert("pupus")
	radix.Insert("yuber")
	radix.Insert("yubel")
	radix.Insert("yubo")

	// Trivial Example 2:
	//radix.Insert("")
	//radix.Insert("test")
	//radix.Insert("toaster")
	//radix.Insert("toasting")
	//radix.Insert("slow")
	//radix.Insert("slowly")
}

func printRecursive(n *Radix, level int) {
	fmt.Println(string(n.path), " - ", level)
	if len(n.nodes) > 0 {
		for _, c := range n.nodes {
			printRecursive(c, level+1)
		}
	}
}
