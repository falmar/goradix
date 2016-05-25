// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

func TestAutoComplete(t *testing.T) {
	radix := New()
	insertData(radix, sampleData)

	test := func(toLook string, expectedData []string, wholeWord bool) {
		words, err := radix.AutoComplete(toLook, wholeWord)

		if err != nil {
			t.Fail()
			return
		}

		if len(words) != len(expectedData) {
			t.Logf("Word count mismatch, Expected: %d; Got: %d", len(expectedData), len(words))
			t.FailNow()
		}

		for _, es := range expectedData {
			var match bool

			for _, s := range words {
				if es == s {
					match = true
				}
			}

			if !match {
				t.Logf("Word not found; Look for: %s; Contains only: %v", es, words)
				t.FailNow()
			}
		}
	}

	test("toa", []string{"ster", "sting"}, false)
	test("te", []string{"st"}, false)
	test("slo", []string{"wly"}, false)
	test("toa", []string{"toaster", "toasting"}, true)
	test("slo", []string{"slowly"}, true)
	test("te", []string{"test"}, true)

	// TODO: Fix autocomplete to return not only leaf
	// test("slo", []string{"w", "ly"})

}
