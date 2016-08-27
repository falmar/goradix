// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

// ----------------------- Benchmarks ------------------------ //

func BenchmarkInsertString(b *testing.B) {
	rx := New()
	tn := 0
	sd2 := sampleData2()
	sdLen := len(sd2) - 1

	for i := 0; i < b.N; i++ {
		if tn == sdLen {
			rx = New()
			tn = 0
		}

		rx.Insert(sd2[tn], i)

		tn++
	}
}

func BenchmarkInsertBytes(b *testing.B) {
	rx := New()
	tn := 0
	sd2 := sampleData3()
	sdLen := len(sd2) - 1

	for i := 0; i < b.N; i++ {
		if tn == sdLen {
			rx = New()
			tn = 0
		}

		rx.InsertBytes(sd2[tn], i)

		tn++
	}
}
