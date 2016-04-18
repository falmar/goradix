package main

import "fmt"

func main() {

	radix := node{}

	insertToRadix(&radix)

	printRecursive(&radix, 0)
}

func insertToRadix(radix *node) {

	// HTTP Router:
	radix.insert("/something/")
	radix.insert("/user/")
	radix.insert("/admin/")
	radix.insert("/admin/auth/")
	radix.insert("/user/profile/")
	radix.insert("/user/:name/")
	radix.insert("/uses/")
	radix.insert("/admin/products/")
	radix.insert("/admin/more/")
	radix.insert("/search/")
	radix.insert("/support/")

	// Trivial Example 1:
	//radix.insert("romane")
	//radix.insert("romanus")
	//radix.insert("romulus")
	//radix.insert("rubens")
	//radix.insert("ruber")
	//radix.insert("rubicon")
	//radix.insert("rubicundus")

	// Trivial Example 2:
	//radix.insert(" ")
	//radix.insert(" test")
	//radix.insert(" toaster")
	//radix.insert(" toasting")
	//radix.insert(" slow")
	//radix.insert(" slowly")
}

func printRecursive(n *node, level int) {
	fmt.Println(string(n.path), " - ", level)
	if len(n.children) > 0 {
		for _, c := range n.children {
			printRecursive(c, level+1)
		}
	}
}
