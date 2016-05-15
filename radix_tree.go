package goradix

import (
	"bytes"
	"errors"
	"fmt"
)

// ErrNoMatchFound self explanatory
var ErrNoMatchFound = errors.New("No Match Found")

// Radix Radix
type Radix struct {
	Path   []byte
	nodes  []*Radix
	parent *Radix
	master bool
}

// New return a Radix Tree
func New() *Radix {
	return &Radix{master: true}
}

// ----------------------- Inserts ------------------------ //

// Insert new string to the Radix Tree
func (r *Radix) Insert(s string) bool {
	return r.insertByteString([]byte(s))
}

func (r *Radix) insertByteString(bs []byte) bool {

	if len(r.Path) == 0 && len(r.nodes) == 0 {
		r.Path = bs
		return true
	}

	match := 0
	i := 0
	var v byte

	for i, v = range r.Path {

		if i >= len(bs) {
			return false
		}

		if v == bs[i] && match == i {
			match++
			continue
		}

		if v != bs[i] && match > 0 {
			if r.nodes == nil {
				r.addChildren(r.Path[i:], nil)
				r.addChildren(bs[i:], nil)
				r.Path = r.Path[:i]
			} else {
				r.pushChildren(bs, i, false)
			}

			return true
		}
	}

	if match > 0 {
		for _, c := range r.nodes {
			if c.insertByteString(bs[i+1:]) {
				return true
			}
		}
		// no match found on nodes
		r.addChildren(bs[i+1:], nil)

		return true
	} else if r.master {
		// Add independent nodes while the master Radix Tree
		// path becomes nil
		if r.Path != nil {
			r.pushChildren(bs, i, true)
			return true
		}

		for _, c := range r.nodes {
			if c.insertByteString(bs) {
				return true
			}
		}

		// no match found on nodes
		// add new node
		r.addChildren(bs, nil)
		return true
	}

	return false
}

func (r *Radix) pushChildren(bs []byte, i int, master bool) {
	nodes := r.nodes
	r.nodes = nil
	r.addChildren(r.Path[i:], nodes)

	if master {
		r.Path = nil
		r.addChildren(bs, nil)
	} else {
		r.Path = r.Path[:i]
		r.addChildren(bs[i:], nil)
	}
}

func (r *Radix) addChildren(bs []byte, c []*Radix) {
	r.nodes = append(r.nodes, &Radix{Path: bs, nodes: c, parent: r})
}

// ----------------------- Match ------------------------ //

func (r Radix) match(bs []byte) (lbs []byte, matches int, plbs []byte) {
	var i = 0
	var v byte

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

	lbs = bs[i:]
	plbs = r.Path[i:]

	return
}

// ----------------------- Look Up ------------------------ //

// TODO: If first letter do not match, continue
// TODO: Optimize match function

// LookUp will return the node matching
func (r *Radix) LookUp(bs []byte) (*Radix, error) {
	var traverseNode = r

	lbs, matches, _ := traverseNode.match(bs)

	if matches == len(traverseNode.Path) {
		if matches < len(bs) {
			for _, n := range traverseNode.nodes {
				if tn, err := n.LookUp(lbs); tn != nil {
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
		err = ErrNoMatchFound
	}

	wordChan := make(chan []byte)
	wordChanOut := make(chan []string)

	go buildWordsWorker(wordChan, wordChanOut)

	fmt.Println(string(lbs))

	if len(lbs) > 0 {
		for _, n := range node.nodes {
			if _, matches, _ := n.match(lbs); matches > 0 {
				if matches == len(lbs) {
					buildWords(n, []byte{}, n.Path, wordChan)
				}
			}
		}
	} else {
		buildWords(node, []byte{}, node.Path, wordChan)
	}

	close(wordChan)
	words = <-wordChanOut

	fmt.Println(len(words))

	if len(words) == 0 {
		err = ErrNoMatchFound
	}

	return
}

func (r *Radix) acLookUp(bs []byte) ([]byte, *Radix, error) {
	var traverseNode = r

	lbs, matches, _ := traverseNode.match(bs)

	fmt.Println("Look for:", string(bs), "in", string(traverseNode.Path), "matches", matches, "turn to", string(lbs))

	if matches == len(traverseNode.Path) {
		if matches <= len(lbs) {
			for _, n := range traverseNode.nodes {
				if nlbs, tn, err := n.acLookUp(lbs); tn != nil {
					return nlbs, tn, err
				}
			}
		}

		return lbs, traverseNode, nil
	}

	return lbs, nil, ErrNoMatchFound
}

func buildWords(rt *Radix, bs, strip []byte, words chan<- []byte) {
	var npath []byte

	npath = append(bs, rt.Path...)

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
