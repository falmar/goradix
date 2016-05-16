package goradix

import "testing"

func TestLookUp(t *testing.T) {
	radix := New()

	insertData(radix, sampleData)

	checkLookUp := func(toLook, expected string, expectedError error) {
		node, err := radix.LookUp([]byte(toLook))

		if err != expectedError {
			t.Logf("Expected Error: %v; Got: %v", expectedError, err)
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
	checkLookUp("t", "t", nil)
	checkLookUp("toast", "oast", nil)
	checkLookUp("toasting", "ing", nil)
	checkLookUp("test", "est", nil)
	checkLookUp("slow", "slow", nil)
	checkLookUp("slowly", "ly", nil)

	//Intentional fails
	checkLookUp("toastar", "", ErrNoMatchFound)
	checkLookUp("toasto", "", ErrNoMatchFound)
	checkLookUp("tast", "", ErrNoMatchFound)
	checkLookUp("slowe", "", ErrNoMatchFound)
}

// ----------------------- Benchmarks ------------------------ //

func BenchmarkLookUp(b *testing.B) {
	radix := New()
	insertData(radix, sampleData2)

	for i := 0; i < b.N; i++ {
		radix.LookUp([]byte("pepocundus"))
	}
}

func randomBytes() []byte {
	strings := sampleData2()
	return []byte(strings[random(0, len(strings))])
}
