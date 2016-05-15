package goradix

import "fmt"

func exampleData(radix *Radix) {
	radix.Insert("test")
	radix.Insert("toaster")
	radix.Insert("toasting")
	radix.Insert("slow")
	radix.Insert("slowly")
}

func exampleData2(radix *Radix) {
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

func printRecursive(n *Radix, level int) {
	fmt.Println(string(n.Path), " - ", level)
	if len(n.nodes) > 0 {
		for _, c := range n.nodes {
			printRecursive(c, level+1)
		}
	}
}
