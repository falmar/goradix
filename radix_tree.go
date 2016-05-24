// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

import "errors"

// ErrNoMatchFound self explanatory
var ErrNoMatchFound = errors.New("No Match Found")

// Radix Radix
type Radix struct {
	Path   []byte
	nodes  []*Radix
	parent *Radix
	master bool
	value  interface{}
	leaf   bool
}

// New return a Radix Tree
func New() *Radix {
	return &Radix{master: true}
}

// ----------------------- Basic ------------------------ //

// Set a value to the Radix Tree node
func (r *Radix) Set(v interface{}) {
	r.value = v
}

// Get a value from Radix Tree node
func (r *Radix) Get() interface{} {
	return r.value
}

// Nodes children
func (r *Radix) Nodes() []*Radix {
	return r.nodes
}

// ----------------------- Inserts ------------------------ //

// Insert new string to the Radix Tree
func (r *Radix) Insert(s string, v ...interface{}) bool {
	return r.InsertBytes([]byte(s), v...)
}

// InsertBytes to the Radix Tree
func (r *Radix) InsertBytes(bs []byte, val ...interface{}) bool {
	var value interface{}

	if len(val) > 0 {
		value = val[0]
	}

	if len(r.Path) == 0 && len(r.nodes) == 0 {
		r.Path = bs
		r.Set(value)
		r.leaf = true
		return true
	}

	match := 0
	i := 0
	var v byte

	for i, v = range r.Path {
		if i >= len(bs) {
			// No more matchs to check
			return false
		}

		if v == bs[i] && match == i {
			// continue as long it match the path
			match++
			continue
		}

		if v != bs[i] && match > 0 {
			// If the byte string does not match anymore but had
			// previous matches to the path then add the byte string
			// as children node
			if r.nodes == nil {
				// If there is no existing nodes then slice the path
				// until the last occurrence, add what is left of the path as
				// children and also add the byte string.
				r.addChildren(r.Path[i:], r.Get(), nil)
				r.Set(nil)
				r.addChildren(bs[i:], value, nil)
				r.Path = r.Path[:i]
			} else {
				// Otherwise just add the new byte string as
				r.pushChildren(bs, value, i, false)
			}

			return true
		}
	}

	if match > 0 {
		// If it matches all current node path and the byte string
		for _, c := range r.nodes {
			if c.InsertBytes(bs[i+1:], value) {
				return true
			}
		}
		// no match found on nodes
		r.addChildren(bs[i+1:], value, nil)

		return true
	}

	if r.master {
		// If there is NO match and the current node is the master Radix

		if r.Path != nil {
			r.pushChildren(bs, value, i, true)
			return true
		}

		for _, c := range r.nodes {
			if c.InsertBytes(bs, value) {
				return true
			}
		}

		// no match found on children nodes
		// add new byte string as node
		r.addChildren(bs, value, nil)
		return true
	}

	return false
}

// Add children node to the current Radix Tree node
func (r *Radix) addChildren(bs []byte, v interface{}, c []*Radix) *Radix {
	nNode := &Radix{Path: bs, nodes: c, parent: r, value: v}
	r.nodes = append(r.nodes, nNode)

	return nNode
}

// Push the current children nodes to a new node with the path
// of what is left from slicing of the current path
// and add the new byte string as children node
func (r *Radix) pushChildren(bs []byte, v interface{}, i int, master bool) {
	nodes := r.nodes
	r.nodes = nil
	r.addChildren(r.Path[i:], r.Get(), nodes)
	r.Set(nil)

	if master {
		r.Path = nil
		r.addChildren(bs, v, nil)
	} else {
		r.Path = r.Path[:i]
		r.addChildren(bs[i:], v, nil)
	}
}

// ----------------------- Match ------------------------ //

func (r Radix) match(bs []byte) ([]byte, int, []byte) {
	var i int
	var v byte
	var matches int

	for i < len(r.Path) {
		v = r.Path[i]
		if i >= len(bs) {
			break
		}

		if bs[i] == v && matches == i {
			matches++
		} else if bs[i] != v {
			break
		}
		i++
	}

	return bs[i:], matches, r.Path[i:]
}

// ----------------------- Look Up ------------------------ //

// LookUp will return the node matching
func (r *Radix) LookUp(s string) (interface{}, error) {
	return r.LookUpBytes([]byte(s))
}

// LookUpBytes will return the node matching
func (r *Radix) LookUpBytes(bs []byte) (interface{}, error) {
	node, err := r.sLookUp(bs)

	if err != nil {
		return nil, err
	}

	return node.Get(), err
}

func (r *Radix) sLookUp(bs []byte) (*Radix, error) {
	var traverseNode = r
	lbs, matches, _ := traverseNode.match(bs)

	// && ((!r.master && matches > 0) || r.master)
	if matches == len(traverseNode.Path) {
		if matches < len(bs) {
			for _, n := range traverseNode.nodes {
				if tn, err := n.sLookUp(lbs); tn != nil {
					return tn, err
				}
			}

			// Do not jump back to parent node
			return nil, ErrNoMatchFound
		}

		return traverseNode, nil
	}

	return nil, ErrNoMatchFound
}

// ----------------------- Autocomplete ------------------------ //

// AutoComplete will complete the word you are looking for
// and return them as string
func (r Radix) AutoComplete(s string) ([]string, error) {
	var stringWords []string
	byteWords, err := r.AutoCompleteBytes([]byte(s))

	if err == nil {
		for _, bs := range byteWords {
			stringWords = append(stringWords, string(bs))
		}
	}

	return stringWords, err
}

// AutoCompleteBytes will complete the word you are looking for
// and return them as slice of bytes
func (r Radix) AutoCompleteBytes(bs []byte) ([][]byte, error) {
	node, strip, err := r.acLookUp(bs)

	if err != nil {
		return nil, err
	}

	inWord := make(chan []byte)
	outWords := make(chan [][]byte)

	go buildWordsWorker(inWord, outWords)
	buildWords(node, []byte{}, strip, inWord)
	close(inWord)
	wordSlice := <-outWords
	close(outWords)

	return wordSlice, err
}

func (r *Radix) acLookUp(bs []byte) (*Radix, []byte, error) {
	var traverseNode = r

	lbs, matches, _ := traverseNode.match(bs)

	if matches > 0 || (matches == 0 && traverseNode.master) {
		if matches == len(traverseNode.Path) && matches < len(bs) {
			for _, n := range traverseNode.nodes {
				if tn, nlbs, err := n.acLookUp(lbs); tn != nil {
					return tn, nlbs, err
				}
			}
		}

		if len(lbs) == 0 {
			return traverseNode, bs, nil
		}
	}

	return nil, nil, ErrNoMatchFound
}
