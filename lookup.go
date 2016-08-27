// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package goradix

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

	if !node.key {
		return nil, ErrNoMatchFound
	}

	return node.get(), err
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
