// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

// ----------------------- Benchmarks ------------------------ //

func BenchmarkLookUpStringSingle(b *testing.B) {
	rx := New()
	insertData(rx, sampleData2)

	toLookUp := randomString(sampleData2())

	for i := 0; i < b.N; i++ {
		rx.LookUp(toLookUp)
	}
}

func BenchmarkLookUpStringRandom(b *testing.B) {
	rx := New()
	insertData(rx, sampleData2)
	sd2 := sampleData2()

	for i := 0; i < b.N; i++ {
		if i%20 == 0 {
			rx = New()
		}

		rx.LookUp(randomString(sd2))
	}
}

func BenchmarkLookUpBytesSingle(b *testing.B) {
	rx := New()
	insertDataBytes(rx, sampleData3)

	toLookUp := randomBytes(sampleData3())

	for i := 0; i < b.N; i++ {
		rx.LookUpBytes(toLookUp)
	}
}

func BenchmarkLookUpBytesRandom(b *testing.B) {
	rx := New()
	insertDataBytes(rx, sampleData3)
	sd3 := sampleData3()

	for i := 0; i < b.N; i++ {
		if i%20 == 0 {
			rx = New()
		}

		rx.LookUpBytes(randomBytes(sd3))
	}
}
