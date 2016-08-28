// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import (
	"sync"
	"testing"
)

func TestCC(t *testing.T) {

	var radix *Radix

	wg := &sync.WaitGroup{}

	for z := 0; z < 20; z++ {

		radix = New()

		for f := 0; f < 20; f++ {
			wg.Add(1)
			go func(r *Radix, w *sync.WaitGroup) {
				for _, s := range sampleData2() {
					r.Insert(s)
				}

				w.Done()
			}(radix, wg)

			wg.Add(1)
			go func(r *Radix, w *sync.WaitGroup) {
				for _, s := range sampleData() {
					r.Insert(s)
				}

				w.Done()
			}(radix, wg)

			wg.Add(1)
			go func(r *Radix, w *sync.WaitGroup) {
				for _, s := range sampleData2() {
					r.LookUp(s)
				}

				w.Done()
			}(radix, wg)

			wg.Add(1)
			go func(r *Radix, w *sync.WaitGroup) {
				for _, s := range sampleData() {
					r.LookUp(s)
				}

				w.Done()
			}(radix, wg)

			wg.Add(1)
			go func(r *Radix, w *sync.WaitGroup) {
				for _, s := range sampleData() {
					r.Remove(s)
				}

				w.Done()
			}(radix, wg)

		}

		wg.Wait()
	}

}
