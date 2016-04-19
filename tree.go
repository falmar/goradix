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

func (n *node) lookUp(s string) (*node, map[string]string) {

	var traverseNode = n
	var bs = []byte(s)
	var params = map[string]string{}

	var i = 0
	var matches = 0
	var bsiMatches = 0
	var bsi = -1
	var v byte

	for i = 0; i < len(traverseNode.path); i++ {
		v = traverseNode.path[i]
		bsi++

		if bsi >= len(bs) {
			break
		}

		if string(v) == ":" {
			var param = []byte{}
			var value = []byte{}

			for f := bsi; f < len(bs); f++ {
				if string(bs[f]) != "/" {
					bsiMatches++
					value = append(value, bs[f])
					continue
				}
				bsi += f
				break
			}

			for f := i; f < len(traverseNode.path); f++ {
				if string(traverseNode.path[f]) != "/" {
					matches++
					param = append(param, traverseNode.path[f])
					continue
				}
				params[string(param)] = string(value)
				i += f
				break
			}

			v = traverseNode.path[i]
		}

		if bs[bsi] == v && matches == i {
			matches++
			bsiMatches++
		} else if bs[bsi] != v {
			break
		}

		continue
	}

	if matches == len(traverseNode.path) {
		if bsiMatches < len(bs) {
			for _, c := range traverseNode.children {
				if tn, ps := c.lookUp(string(bs[bsi+1:])); tn != nil {
					for i, v := range ps {
						params[i] = v
					}
					return tn, params
				}
			}
		}

		return traverseNode, params
	}

	return nil, nil
}
