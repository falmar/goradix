// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

// ----------------------- Benchmarks ------------------------ //

func BenchmarkAutocompleteStringLeftWords(b *testing.B) {
	rx := New(false)
	insertData(rx, sampleData2)
	sd2 := sampleData2()

	for i := 0; i < b.N; i++ {
		rx.AutoComplete(randomString(sd2), false)
	}
}

func BenchmarkAutocompleteStringFullWords(b *testing.B) {
	rx := New(false)
	insertData(rx, sampleData2)
	sd2 := sampleData2()

	for i := 0; i < b.N; i++ {
		rx.AutoComplete(randomString(sd2), true)
	}
}

func BenchmarkAutocompleteBytesLeftWords(b *testing.B) {
	rx := New(false)
	insertDataBytes(rx, sampleData2Bytes)
	sd3 := sampleData2Bytes()

	for i := 0; i < b.N; i++ {
		rx.AutoCompleteBytes(randomBytes(sd3), false)
	}
}

func BenchmarkAutocompleteBytesFullWords(b *testing.B) {
	rx := New(false)
	insertDataBytes(rx, sampleData2Bytes)
	sd3 := sampleData2Bytes()

	for i := 0; i < b.N; i++ {
		rx.AutoCompleteBytes(randomBytes(sd3), true)
	}
}
