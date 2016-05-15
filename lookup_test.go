package goradix

import "testing"

func TestLookUp(t *testing.T) {
	radix := New()

	exampleData(radix)

	checkLookUp := func(toLook, expected string) {

		node, err := radix.LookUp([]byte(toLook))

		if err != nil {
			t.Fail()
			t.Log("Error:", err)
			return
		}

		expectedPath := []byte(expected)

		for i, v := range node.Path {
			if v != expectedPath[i] {
				t.Fail()
				t.Logf("Expected: %d; Got: %d", expectedPath[i], v)
			}
		}
	}

	checkLookUp("s", "oast")
}
