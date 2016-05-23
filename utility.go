// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import (
	"bytes"
	"fmt"
	"math/rand"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

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

func insertData(radix *Radix, sd func() []string) {
	for i, s := range sd() {
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

func buildWords(rt *Radix, bs, strip []byte, words chan<- []byte) {
	var npath []byte
	npath = append(bs, rt.Path...)

	if len(rt.nodes) > 0 {
		for _, n := range rt.nodes {
			buildWords(n, npath, strip, words)
		}
	} else {
		words <- bytes.Replace(npath, strip, nil, 1)
	}
}

func buildWordsWorker(inWords <-chan []byte, outWords chan<- [][]byte) {
	var wordSlice [][]byte

	for v := range inWords {
		wordSlice = append(wordSlice, v)
	}

	outWords <- wordSlice
}
