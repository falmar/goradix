// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

func TestInsert(t *testing.T) {
	radix := New()
	radix.Insert("test", 1)

	expectedBytes := []byte("test")
	expectedValue := 1

	for i, v := range radix.Path {
		if v != expectedBytes[i] {
			t.Fail()
			t.Logf("Expected: %d; got: %d", v, expectedBytes[i])
		}
	}

	if radix.Get() != expectedValue {
		t.Fail()
		t.Logf("Expected value: %v; got: %v", radix.Get(), expectedValue)
	}
}

func TestInsertSeparation(t *testing.T) {
	radix := New()
	radix.Insert("toaster", "value1")
	radix.Insert("toasting", "value2")
	radix.Insert("toast", "value3")

	expectedText := [][]byte{
		[]byte("toast"), []byte("er"), []byte("ing"),
	}
	expectedValues := []string{"value3", "value1", "value2"}

	for i, v := range radix.Path {
		if v != expectedText[0][i] {
			t.Fail()
			t.Logf("Expected: %d; got: %d", v, expectedText[0][i])
		}
	}

	if radix.Get() != expectedValues[0] {
		t.Fail()
		t.Logf("Expected value: %v; got: %v", radix.Get(), expectedValues[0])
	}

	if radix.key != true {
		t.Fail()
		t.Logf("Expected key: %v; Got: %v", true, radix.key)
	}

	for i, n := range radix.nodes {
		for ii, v := range n.Path {
			if v != expectedText[i+1][ii] {
				t.Fail()
				t.Logf("Expected: %d; got: %d", v, expectedText[i+1][ii])
			}
		}

		if n.Get() != expectedValues[i+1] {
			t.Fail()
			t.Logf("Expected value: %v; got: %v", n.Get(), expectedValues[i+1])
		}

		if n.key != true {
			t.Fail()
			t.Logf("Expected key: %v; Got: %v", true, n.key)
		}
	}
}

func checkNodes(t *testing.T, nodes []*Radix, expectedText [][]byte, expectedValue []interface{}, expectedKeys []bool, level int) int {
	for _, n := range nodes {
		for i, v := range n.Path {
			if v != expectedText[level][i] {
				t.Fail()
				t.Logf("Expected: %s; got: %s", string(expectedText[level][i]), string(v))
			}
		}

		if n.Get() != expectedValue[level] {
			t.Fail()
			t.Logf("Expected value: %v; got: %v", expectedValue[level], n.Get())
		}

		if n.key != expectedKeys[level] {
			t.Fail()
			t.Logf("Expected Key: %v; Got: %v.", expectedKeys[level], n.key)
		}

		level = checkNodes(t, n.nodes, expectedText, expectedValue, expectedKeys, level+1)
	}
	return level
}

func TestInsertSeparationComplex(t *testing.T) {
	radix := New()
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

	if radix.Path != nil {
		t.Fail()
		t.Logf("Expected: %v; got: %v", nil, radix.Path)
	}

	checkNodes(t, radix.nodes, expectedTexts, expectedValues, expectedKeys, 0)
}
