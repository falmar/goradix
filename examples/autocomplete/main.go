// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/falmar/goradix"
)

func main() {

	radix := goradix.New()

	fmt.Printf("%v\n", []byte(""))

	exampleData(radix)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Text to Search: ")
	for scanner.Scan() {

		wordSlice, err := radix.AutoComplete(scanner.Text())

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(wordSlice)
		}

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
