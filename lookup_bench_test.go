// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "testing"

// ----------------------- Benchmarks ------------------------ //

func BenchmarkLookUpNTS(b *testing.B) {
	rx := New(false)
	insertDataBytes(rx, sampleData2Bytes)
	sd3 := sampleData2Bytes()

	for i := 0; i < b.N; i++ {
		rx.LookUpBytes(randomBytes(sd3))
	}
}

func BenchmarkLookUpTS(b *testing.B) {
	rx := New(true)
	insertDataBytes(rx, sampleData2Bytes)
	sd3 := sampleData2Bytes()

	for i := 0; i < b.N; i++ {
		rx.LookUpBytes(randomBytes(sd3))
	}
}
