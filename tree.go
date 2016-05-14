package main

import (
	"bytes"
	"errors"
	"fmt"
)

// ErrNoMatch self explanatory
const ErrNoMatch = "No Match Found"

// Radix Radix
type Radix struct {
	path   []byte
	nodes  []*Radix
	master bool
}

// ----------------------- Inserts ------------------------ //

// Insert new string to the Radix Tree
func (r *Radix) Insert(s string) bool {
	return r.insertByteString([]byte(s))
}

func (r *Radix) insertByteString(bs []byte) bool {

	if len(r.path) == 0 && len(r.nodes) == 0 {
		r.path = bs
		return true
	}

	match := 0
	i := 0
	var v byte

	for i, v = range r.path {

		if i >= len(bs) {
			return false
		}

		if v != bs[i] && match > 0 {
			if r.nodes == nil {
				r.addChildren(bs[i:], nil)
				r.addChildren(r.path[i:], nil)
				r.path = r.path[:i]
			} else {
				r.pushChildren(bs, i, false)
			}

			return true
		} else if v == bs[i] && match == i {
			match++
		}
	}

	if match > 0 {

		for _, c := range r.nodes {
			if c.insertByteString(bs[i+1:]) {
				return true
			}
		}

		// no match found on childrens
		r.addChildren(bs[i+1:], nil)

		return true
	} else if r.master {
		if len(r.path) > 0 {
			r.pushChildren(bs, i, true)
			return true
		}

		for _, c := range r.nodes {
			if c.insertByteString(bs) {
				return true
			}
		}

		r.addChildren(bs, nil)
		return true
	}

	return false
}

func (r *Radix) pushChildren(bs []byte, i int, master bool) {
	nodes := r.nodes
	r.nodes = nil
	r.addChildren(r.path[i:], nodes)

	if master {
		r.path = []byte{}
		r.addChildren(bs, nil)
	} else {
		r.path = r.path[:i]
		r.addChildren(bs[i:], nil)
	}
}

func (r *Radix) addChildren(bs []byte, c []*Radix) {
	r.nodes = append(r.nodes, &Radix{path: bs, nodes: c})
}

// ----------------------- Look Up ------------------------ //

// LookUp will return the node matching
func (r *Radix) LookUp(bs []byte) (*Radix, error) {
	var traverseNode = r

	lbs, matches := r.match(traverseNode.path, bs)

	if matches == len(traverseNode.path) {
		if matches < len(bs) {
			for _, c := range traverseNode.nodes {
				if tn, err := c.LookUp(lbs); tn != nil {
					return tn, err
				}
			}

			return nil, errors.New(ErrNoMatch)
		}

		return traverseNode, nil
	}

	return nil, errors.New(ErrNoMatch)
}

func (r Radix) match(path, bs []byte) (lbs []byte, matches int) {
	var i = 0
	var v byte

	for i < len(path) {
		v = path[i]
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

	lbs = bs[i:]

	return
}

// ----------------------- Autocomplete ------------------------ //

// Autocomplete will return the words out of Looked Up string
func (r Radix) Autocomplete(s string) (node *Radix, words []string, err error) {
	var lbs []byte
	bs := []byte(s)
	lbs, node, err = r.acLookUp(bs)

	if err != nil {
		return
	}

	fmt.Println(node)

	if node.master && len(bs) > 0 {
		err = errors.New(ErrNoMatch)
	}

	wordChan := make(chan []byte)
	wordChanOut := make(chan []string)

	go buildWordsWorker(wordChan, wordChanOut)

	fmt.Println(string(lbs))

	if len(lbs) > 0 {
		for _, n := range node.nodes {
			if _, matches := r.match(n.path, lbs); matches > 0 {
				if matches == len(lbs) {
					buildWords(n, []byte{}, n.path, wordChan)
				}
			}
		}
	} else {
		buildWords(node, []byte{}, node.path, wordChan)
	}

	close(wordChan)
	words = <-wordChanOut

	fmt.Println(len(words))

	if len(words) == 0 {
		err = errors.New(ErrNoMatch)
	}

	return
}

func (r *Radix) acLookUp(bs []byte) ([]byte, *Radix, error) {
	var traverseNode = r

	lbs, matches := r.match(traverseNode.path, bs)

	fmt.Println("Look for:", string(bs), "in", string(traverseNode.path), "matches", matches, "turn to", string(lbs))

	if matches == len(traverseNode.path) {
		if matches <= len(lbs) {
			for _, n := range traverseNode.nodes {
				if nlbs, tn, err := n.acLookUp(lbs); tn != nil {
					return nlbs, tn, err
				}
			}
		}

		return lbs, traverseNode, nil
	}

	return lbs, nil, errors.New(ErrNoMatch)
}

func buildWords(rt *Radix, bs, strip []byte, words chan<- []byte) {
	var npath []byte

	npath = append(bs, rt.path...)

	if len(rt.nodes) > 0 {
		for _, n := range rt.nodes {
			buildWords(n, npath, strip, words)
		}
	} else {
		fmt.Println("Write word: ", string(npath), string(strip), "Becomes:", string(bytes.Replace(npath, strip, []byte(""), -1)))
		words <- bytes.Replace(npath, strip, []byte(""), 1)
	}
}

func buildWordsWorker(inWords <-chan []byte, outWords chan<- []string) {
	var wordSlice []string

	for v := range inWords {
		wordSlice = append(wordSlice, string(v))
	}

	outWords <- wordSlice
}
