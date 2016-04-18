package main

type node struct {
	path     []byte
	children []*node
}

func (n *node) insert(s string) bool {

	bs := []byte(s)

	if len(n.path) == 0 {
		n.path = bs
		return true
	}

	match := 0
	i := 0
	var v byte

	for i, v = range n.path {

		if i >= len(bs) {
			return false
		}

		if v != bs[i] && match > 0 {

			if n.children == nil {
				n.addChildren(bs[i:], nil)
				n.addChildren(n.path[i:], nil)
				n.path = n.path[:i]
			} else {
				children := n.children
				n.children = nil
				n.addChildren(n.path[i:], children)
				n.path = n.path[:i]
				n.addChildren(bs[i:], nil)
			}

			return true
		} else if v == bs[i] && match == i {
			match++
		}
	}

	if match > 0 {

		for _, c := range n.children {
			if c.insert(string(bs[i+1:])) {
				return true
			}
		}

		// no match found on childrens
		n.addChildren(bs[i+1:], nil)

		return true
	}

	return false
}

func (n *node) addChildren(bs []byte, c []*node) {
	n.children = append(n.children, &node{path: bs, children: c})
}

func (n *node) lookUp(s string) *node {

	var traverseNode = n
	var matches = 0
	var bs = []byte(s)

	var i = 0
	var v byte

	for i, v = range traverseNode.path {

		if i >= len(bs) {
			break
		}

		if bs[i] == v && matches == i {
			matches++
		} else if bs[i] != v {
			break
		}

	}

	if matches == len(traverseNode.path) {

		if matches < len(bs) {
			for _, c := range traverseNode.children {
				if tn := c.lookUp(string(bs[i+1:])); tn != nil {
					return tn
				}
			}
		}

		return traverseNode
	}

	return nil
}
