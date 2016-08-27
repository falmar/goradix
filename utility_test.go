// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import (
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

// same as sampleData2 but in bytes to avoid conversion
func sampleData3() [][]byte {
	return [][]byte{
		[]byte{114, 111, 109, 97, 110, 101},
		[]byte{114, 111, 109, 97, 110, 117, 115},
		[]byte{114, 111, 109, 117, 108, 117, 115},
		[]byte{114, 117, 98, 101, 110, 115},
		[]byte{114, 117, 98, 101, 114},
		[]byte{114, 117, 98, 105, 99, 111, 110},
		[]byte{114, 117, 98, 105, 99, 117, 110, 100, 117, 115},
		[]byte{114, 105, 110, 105, 99, 117, 110, 100, 117, 115},
		[]byte{114, 101, 112, 105, 99, 117, 110, 100, 117, 115},
		[]byte{108, 101, 112, 105, 99, 117, 110, 100, 117, 115},
		[]byte{108, 101, 112, 111, 99, 117, 110, 100, 117, 115},
		[]byte{108, 111, 109, 117, 108, 117, 115},
		[]byte{108, 117, 112, 117, 115},
		[]byte{104, 117, 98, 101, 114},
		[]byte{112, 101, 112, 105, 99, 117, 110, 100, 117, 115},
		[]byte{112, 101, 112, 105, 99, 117, 110, 100, 97, 115},
		[]byte{112, 101, 112, 111, 99, 117, 110, 100, 117, 115},
		[]byte{112, 111, 109, 117, 108, 117, 115},
		[]byte{112, 117, 112, 117, 115},
		[]byte{121, 117, 98, 101, 114},
		[]byte{121, 117, 98, 101, 108},
		[]byte{121, 117, 98, 111},
	}
}

func insertData(radix *Radix, sd func() []string) {
	for i, s := range sd() {
		radix.Insert(s, i)
	}
}

func insertDataBytes(radix *Radix, sd func() [][]byte) {
	for i, s := range sd() {
		radix.InsertBytes(s, i)
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

func randomBytes(bytes [][]byte) []byte {
	return bytes[random(0, len(bytes))]
}

func randomString(strings []string) string {
	return strings[random(0, len(strings))]
}
