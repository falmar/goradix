// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

func TestInsert(t *testing.T) {
	radix := New(false)
	radix.Insert("test", 1)

	expectedBytes := []byte("test")
	expectedValue := 1
	expectedLeaf := true

	for i, v := range radix.Path {
		if v != expectedBytes[i] {
			t.Fatalf("Expected: %d; got: %d", v, expectedBytes[i])
		}
	}

	if radix.get() != expectedValue {
		t.Fatalf("Expected value: %v; got: %v", radix.get(), expectedValue)
	}

	if radix.leaf != expectedLeaf {
		t.Fatalf("Expected leaf: %v; got: %v", expectedLeaf, radix.leaf)
	}
}

func TestInsertSeparation(t *testing.T) {
	radix := New(false)
	radix.Insert("toaster", "value1")
	radix.Insert("toasting", "value2")
	radix.Insert("toast", "value3")

	expectedText := [][]byte{
		[]byte("toast"), []byte("er"), []byte("ing"),
	}
	expectedValues := []string{"value3", "value1", "value2"}
	expectedLeaves := []bool{false, true, true}

	for i, v := range radix.Path {
		if v != expectedText[0][i] {
			t.Fatalf("Expected: %d; got: %d", v, expectedText[0][i])
		}
	}

	if radix.get() != expectedValues[0] {
		t.Fatalf("Expected value: %v; got: %v", radix.get(), expectedValues[0])
	}

	if radix.key != true {
		t.Fatalf("Expected key: %v; Got: %v", true, radix.key)
	}

	if radix.leaf != expectedLeaves[0] {
		t.Fatalf("Expected leaf: %v; got: %v", expectedLeaves[0], radix.leaf)
	}

	for i, n := range radix.nodes {
		for ii, v := range n.Path {
			if v != expectedText[i+1][ii] {
				t.Fatalf("Expected: %d; got: %d", v, expectedText[i+1][ii])
			}
		}

		if n.get() != expectedValues[i+1] {
			t.Fatalf("Expected value: %v; got: %v", n.get(), expectedValues[i+1])
		}

		if n.key != true {
			t.Fatalf("Expected key: %v; Got: %v", true, n.key)
		}

		if n.leaf != expectedLeaves[i+1] {
			t.Fatalf("Expected leaf: %v; got: %v", expectedLeaves[i+1], n.leaf)
		}
	}
}

func checkNodes(t *testing.T, nodes []*Radix, expectedText [][]byte, expectedValue []interface{}, expectedKeys, expectedLeaves []bool, level int) int {
	for _, n := range nodes {
		for i, v := range n.Path {
			if v != expectedText[level][i] {
				t.Fail()
				t.Logf("Expected: %s; got: %s", string(expectedText[level][i]), string(v))
			}
		}

		if n.get() != expectedValue[level] {
			t.Fail()
			t.Logf("Expected value: %v; got: %v", expectedValue[level], n.get())
		}

		if n.key != expectedKeys[level] {
			t.Fail()
			t.Logf("Expected Key: %v; Got: %v.", expectedKeys[level], n.key)
		}

		if n.leaf != expectedLeaves[level] {
			t.Fatalf("Expected leaf: %v; got: %v", expectedLeaves[level], n.leaf)
		}

		level = checkNodes(t, n.nodes, expectedText, expectedValue, expectedKeys, expectedLeaves, level+1)
	}
	return level
}

func TestInsertSeparationComplex(t *testing.T) {
	radix := New(false)
	insertData(radix, sampleData)

	expectedTexts := [][]byte{
		[]byte("t"),
		[]byte("est"),
		[]byte("oast"),
		[]byte("er"),
		[]byte("ing"),
		[]byte("slow"),
		[]byte("ly"),
	}

	expectedValues := []interface{}{
		nil, 0, nil, 1, 2, 3, 4,
	}

	expectedKeys := []bool{
		false, true, false, true, true, true, true,
	}

	expectedLeaves := []bool{
		false, true, false, true, true, false, true,
	}

	if radix.Path != nil {
		t.Fail()
		t.Logf("Expected: %v; got: %v", nil, radix.Path)
	}

	checkNodes(t, radix.nodes, expectedTexts, expectedValues, expectedKeys, expectedLeaves, 0)
}

func TestInsertFirstMatchingLetter(t *testing.T) {
	radix := New(false)
	radix.Insert("roma", "value1")
	radix.Insert("romb", "value2")
	radix.Insert("abc", "value3")

	expectedTexts := [][]byte{
		[]byte("rom"),
		[]byte("a"),
		[]byte("b"),
		[]byte("abc"),
	}
	expectedValues := []interface{}{
		nil, "value1", "value2", "value3",
	}
	expectedKeys := []bool{
		false, true, true, true,
	}
	expectedLeaves := []bool{
		false, true, true, true,
	}

	checkNodes(t, radix.nodes, expectedTexts, expectedValues, expectedKeys, expectedLeaves, 0)
}
