package goradix

import (
	"fmt"
	"math/rand"
)

func sampleData() []string {
	return []string{
		"test",
		"toaster",
		"toasting",
		"slow",
		"slowly",
	}
}

func sampleData2() []string {
	return []string{
		"romane",
		"romanus",
		"romulus",
		"rubens",
		"ruber",
		"rubicon",
		"rubicundus",
		"rinicundus",
		"repicundus",
		"lepicundus",
		"lepocundus",
		"lomulus",
		"lupus",
		"huber",
		"pepicundus",
		"pepicundas",
		"pepocundus",
		"pomulus",
		"pupus",
		"yuber",
		"yubel",
		"yubo",
	}
}

func insertData(radix *Radix, cb func() []string) {
	for i, s := range cb() {
		radix.Insert(s, i)
	}
}

func printRecursive(n *Radix, level int) {
	fmt.Println(string(n.Path), " - ", level)
	if len(n.nodes) > 0 {
		for _, c := range n.nodes {
			printRecursive(c, level+1)
		}
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
