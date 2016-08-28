// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

func TestRemove(t *testing.T) {
	rx := New()
	insertData(rx, sampleData)

	rx.Remove("slow")

	if rx.nodes[1].key {
		t.Fatal("Expected node to not be key")
	}

	if rx.nodes[1].value != nil {
		t.Fatalf("Expected node value to be nil; Got: %v", rx.nodes[1].value)
	}
}

func TestRemoveChild(t *testing.T) {
	rx := New()
	insertData(rx, sampleData)

	rx.Remove("test")

	if len(rx.nodes[0].nodes) != 1 {
		t.Fatalf("Expected node to have only 1 child; Got: %d", len(rx.nodes[0].nodes))
	}
}

func TestRemoveChildMerge(t *testing.T) {
	rx := New()
	insertData(rx, sampleData)

	rx.Remove("slow")

	if string(rx.nodes[1].Path) != "slowly" {
		t.Fatal("Expected child node to be merged")
	}
}
