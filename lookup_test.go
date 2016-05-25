// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

func TestSourceLookUp(t *testing.T) {
	radix := New()
	insertData(radix, sampleData)

	checkLookUp := func(toLook, expected string, expectedLeaf bool, expectedError error) {
		node, err := radix.sLookUp([]byte(toLook))

		if err != expectedError {
			t.Logf("Expected Error: %v; Got: %v. Node: %s", expectedError, err, node.Path)
			t.Fail()
		}

		if expectedError == ErrNoMatchFound && node == nil {
			return
		}

		if node == nil {
			t.Fail()
			t.Logf("Expected node to not be: %v; Got: %v", nil, node)
			return
		}

		if node.leaf != expectedLeaf {
			t.Fail()
			t.Logf("Expected Leaf: %v; Got: %v", expectedLeaf, node.leaf)
		}

		expectedPath := []byte(expected)

		if len(node.Path) != len(expectedPath) {
			t.Fail()
			t.Logf("Expected path lenght: %d; Got: %d", len(expectedPath), len(node.Path))
			return
		}

		for i, v := range node.Path {
			if v != expectedPath[i] {
				t.Fail()
				t.Logf("Expected: %d; Got: %d", expectedPath[i], v)
			}
		}
	}

	// Correct
	checkLookUp("t", "t", false, nil)
	checkLookUp("toast", "oast", false, nil)
	checkLookUp("toasting", "ing", true, nil)
	checkLookUp("test", "est", true, nil)
	checkLookUp("slow", "slow", false, nil)
	checkLookUp("slowly", "ly", true, nil)
	checkLookUp("toaster", "er", true, nil)

	// Intentional fails
	checkLookUp("toastar", "", false, ErrNoMatchFound)
	checkLookUp("toasto", "", false, ErrNoMatchFound)
	checkLookUp("tast", "", false, ErrNoMatchFound)
	checkLookUp("slowe", "", false, ErrNoMatchFound)
}

func TestLookUp(t *testing.T) {
	radix := New()
	insertData(radix, sampleData)

	toLookUp := []string{"t", "test", "toast", "toasting", "slowly"}
	expectedValues := []interface{}{nil, 0, nil, 2, 4}
	expectedError := []error{ErrNoMatchFound, nil, ErrNoMatchFound, nil, nil}

	for i, s := range toLookUp {
		v, err := radix.LookUp(s)

		if expectedError[i] != err {
			t.Logf("Expected Error: %v; Got: %v", expectedError[i], err)
			t.FailNow()
		}

		if expectedValues[i] != v {
			t.Logf("Expected Value: %v; Got: %v", expectedValues[i], v)
			t.FailNow()
		}
	}
}

// ----------------------- Benchmarks ------------------------ //

func BenchmarkLookUp(b *testing.B) {
	radix := New()
	insertData(radix, sampleData2)

	toLookUp := randomBytes()

	for i := 0; i < b.N; i++ {
		radix.LookUpBytes(toLookUp)
	}
}

func randomBytes() []byte {
	strings := sampleData2()
	return []byte(strings[random(0, len(strings))])
}
