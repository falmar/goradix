// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

// ----------------------- Benchmarks ------------------------ //

func BenchmarkInsertNTS(b *testing.B) {
	rx := New(false)
	tn := 0
	sd2 := sampleData3()
	sdLen := len(sd2) - 1

	for i := 0; i < b.N; i++ {
		if tn == sdLen {
			rx = New(false)
			tn = 0
		}

		rx.InsertBytes(sd2[tn], i)

		tn++
	}
}

func BenchmarkInsertTS(b *testing.B) {
	rx := New(true)
	tn := 0
	sd2 := sampleData3()
	sdLen := len(sd2) - 1

	for i := 0; i < b.N; i++ {
		if tn == sdLen {
			rx = New(true)
			tn = 0
		}

		rx.InsertBytes(sd2[tn], i)

		tn++
	}
}
