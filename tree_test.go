package main

import "testing"

func exampleData1(radix *Radix) {
	radix.Insert("test")
	radix.Insert("toaster")
	radix.Insert("toasting")
	radix.Insert("slow")
	radix.Insert("slowly")
}

func TestInsert(t *testing.T) {
	radix := New()

	radix.Insert("test")
	bs := []byte("test")

	for i, v := range radix.path {
		if v != bs[i] {
			t.Logf("Expected: %d; got: %d", v, bs[i])
		}
	}

	radix = New()

	radix.Insert("test")
	bs = []byte("te2t")

	for i, v := range radix.path {
		if v != bs[i] && i == 3 {
			t.Logf("Expected: %d; got: %d", v, bs[i])
		}
	}
}

func TestInsertSeparation(t *testing.T) {
	radix := New()
	radix.Insert("toaster")
	radix.Insert("toasting")

	masterText := []byte("toast")
	nodesText := [][]byte{[]byte("er"), []byte("ing")}

	for i, v := range radix.path {
		if v != masterText[i] {
			t.Logf("Expected: %d; got: %d", v, masterText[i])
		}
	}

	for i, n := range radix.nodes {
		for ii, v := range n.path {
			if v != nodesText[i][ii] {
				t.Logf("Expected: %d; got: %d", v, nodesText[i][ii])
			}
		}
	}
}

func checkNodes(t *testing.T, nodes []*Radix, TextLevels [][]byte, level int) int {
	for _, n := range nodes {
		for i, v := range n.path {
			if v != TextLevels[level][i] {
				t.Fail()
				t.Logf("Expected: %s; got: %s", string(v), string(TextLevels[level][i]))
			}
		}
		level = checkNodes(t, n.nodes, TextLevels, level+1)
	}
	return level
}

func TestInsertSeparationComplex(t *testing.T) {
	radix := New()
	exampleData1(radix)

	TextLevels := [][]byte{
		[]byte("t"),
		[]byte("est"),
		[]byte("oast"),
		[]byte("er"),
		[]byte("ing"),
		[]byte("slow"),
		[]byte("ly"),
	}

	if radix.path != nil {
		t.Fail()
		t.Logf("Expected: %v; got: %v", nil, radix.path)
	}

	checkNodes(t, radix.nodes, TextLevels, 0)
}
