package goradix

import "testing"

func TestAutoComplete(t *testing.T) {
	radix := New()
	insertData(radix, sampleData)
	var expectedData = []string{"ster", "sting"}
	words, err := radix.AutoComplete("toa")

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
