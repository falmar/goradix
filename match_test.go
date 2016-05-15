package goradix

import "testing"

func TestMatch(t *testing.T) {

	match := func(t *testing.T, insert []string, matchString, lbsMatch, plbsMatch []byte, matchesInt int) {
		//Complex N. 1
		radix := New()

		for _, v := range insert {
			radix.Insert(v)
		}

		lbs, matches, plbs := radix.match([]byte(matchString))

		if matches != matchesInt {
			t.Fail()
			t.Logf("matches fail. Expected %d; Got: %d", matchesInt, matches)
		}

		if len(lbs) != len(lbsMatch) {
			t.Fail()
			t.Logf("left string length fail. Expected %d; Got: %d", len(lbsMatch), len(lbs))
		}

		for i, v := range lbs {
			if v != lbsMatch[i] {
				t.Fail()
				t.Logf("left byte string fail. Expected: %d; Got: %d", lbsMatch[i], v)
			}
		}

		if len(plbs) != len(plbsMatch) {
			t.Fail()
			t.Logf("left string length fail. Expected %d; Got: %d", len(plbsMatch), len(lbs))
		}

		for i, v := range plbs {
			if v != plbsMatch[i] {
				t.Fail()
				t.Logf("left byte string fail. Expected: %d; Got: %d", plbsMatch[i], v)
			}
		}
	}

	match(t, []string{"test"}, []byte("tea"), []byte("a"), []byte("st"), 2)

	match(t, []string{"yubel", "yuber", "yubo"}, []byte("y"), []byte{}, []byte("ub"), 1)
}
