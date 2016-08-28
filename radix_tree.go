// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import (
	"errors"
	"sync"
)

// ErrNoMatchFound self explanatory
var ErrNoMatchFound = errors.New("No Match Found")

// Radix Radix
type Radix struct {
	Path      []byte
	nodes     []*Radix
	parent    *Radix
	master    bool
	value     interface{}
	leaf, key bool
	mu        *sync.RWMutex
	cs        bool
}

// New return a Radix Tree
// cs bool - Concurrent Safe
func New(cs bool) *Radix {
	return &Radix{master: true, mu: &sync.RWMutex{}, cs: cs}
}

// ----------------------- Basic ------------------------ //

// Set a value to the Radix Tree node
func (r *Radix) set(v interface{}) {
	r.value = v
}

// Get a value from Radix Tree node
func (r *Radix) getNonBlocking() interface{} {
	return r.value
}

// Get a value from Radix Tree node
func (r *Radix) get() interface{} {
	r.mu.RLock()
	v := r.value
	r.mu.RUnlock()
	return v
}

// ----------------------- Locks ------------------------ //
// in order to make concurrent safety optional

func (r *Radix) lock() {
	if r.cs {
		r.mu.Lock()
	}
}

func (r *Radix) unlock() {
	if r.cs {
		r.mu.Unlock()
	}
}

func (r *Radix) rLock() {
	if r.cs {
		r.mu.RLock()
	}
}

func (r *Radix) rUnlock() {
	if r.cs {
		r.mu.RUnlock()
	}
}
