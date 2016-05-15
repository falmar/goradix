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

		_, wordSlice, err := radix.Autocomplete(textToLook)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(wordSlice)
		fmt.Print("Text to Search: ")
	}
}

func exampleData(radix *goradix.Radix) {
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
}
