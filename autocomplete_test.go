// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

func TestAutoComplete(t *testing.T) {
	radix := New(false)
	insertData(radix, sampleData)

	test := func(toLook string, expectedData []string, wholeWord bool) {
		words, err := radix.AutoComplete(toLook, wholeWord)

		if err != nil {
			t.Fail()
			return
		}

		if len(words) != len(expectedData) {
			t.Logf("Word count mismatch, Expected: %d; Got: %d; words: %v; tl: %s", len(expectedData), len(words), words, toLook)
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

	test("t", []string{"est", "oaster", "oasting"}, false)
	test("t", []string{"test", "toaster", "toasting"}, true)
	test("toa", []string{"ster", "sting"}, false)
	test("te", []string{"st"}, false)
	test("slo", []string{"w", "wly"}, false)
	test("slow", []string{"ly"}, false)
	test("toa", []string{"toaster", "toasting"}, true)
	test("slo", []string{"slow", "slowly"}, true)
	test("slow", []string{"slowly"}, true)
	test("te", []string{"test"}, true)
	test("test", nil, true)
}
