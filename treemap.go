// TreeMap -- balanced binary search tree implementation of a key/value map
//
// @version 2018-07-09
// @author  Robert Altnoeder (r.altnoeder@gmx.net)
//
// Copyright (C) 2018 Robert ALTNOEDER
//
// Redistribution and use in source and binary forms,
// with or without modification, are permitted provided that
// the following conditions are met:
//
//  1. Redistributions of source code must retain the above copyright notice,
//     this list of conditions and the following disclaimer.
//  2. Redistributions in binary form must reproduce the above copyright
//     notice, this list of conditions and the following disclaimer in
//     the documentation and/or other materials provided with the distribution.
//  3. The name of the author may not be used to endorse or promote products
//     derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
// IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
// OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
// TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
// LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
// NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE,
// EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package dsaext

type treeNode struct {
	key     interface{}
	value   interface{}
	parent  *treeNode
	less    *treeNode
	greater *treeNode
	level   int
}

func newtreeNode(key, value interface{}, parent, less, greater *treeNode) *treeNode {
	return &treeNode{key, value, parent, less, greater, 1}
}

func (node *treeNode) successor() *treeNode {
	retNode := node
	for retNode != nil {
		if retNode.greater != nil {
			// Find the item with the lowest key in the subtree that
			// contains items with higher keys
			retNode = retNode.greater
			for retNode.less != nil {
				retNode = retNode.less
			}
			break
		} else {
			// Find a parent node with a higher value
			for retNode.parent != nil {
				if retNode.parent.less == retNode {
					break
				}
				retNode = retNode.parent
			}
			retNode = retNode.parent
			break
		}
	}
	return retNode
}

func (node *treeNode) predecessor() *treeNode {
	retNode := node
	for retNode != nil {
		if retNode.less != nil {
			// Find the item with the highest key in the subtree that
			// contains items with lower keys
			retNode = retNode.less
			for retNode.greater != nil {
				retNode = retNode.greater
			}
			break
		} else {
			// Find a parent node with a lower value
			for retNode.parent != nil {
				if retNode.parent.less == retNode {
					break
				}
				retNode = retNode.parent
			}
			retNode = retNode.parent
			break
		}
	}
	return retNode
}

type TreeMapIterator struct {
	nextNode *treeNode
}

func (iter *TreeMapIterator) Next() (interface{}, interface{}, bool) {
	var key interface{} = nil
	var value interface{} = nil
	flag := false
	if iter.nextNode != nil {
		key = iter.nextNode.key
		value = iter.nextNode.value
		flag = true
		iter.nextNode = iter.nextNode.successor()
	}
	return key, value, flag
}

type TreeMap struct {
	root  *treeNode
	size  int
	cmpFn compareFn
}

func NewTreeMap(cmpFn compareFn) *TreeMap {
	return &TreeMap{nil, 0, cmpFn}
}

func (tree *TreeMap) Iterator() *TreeMapIterator {
	node := tree.root
	if node != nil {
		// Find the item with the lowest key
		for node.less != nil {
			node = node.less
		}
	}
	return &TreeMapIterator{node}
}

func (tree *TreeMap) Insert(key, value interface{}) {
	if tree.root == nil {
		// Insert at the tree's root
		tree.root = newtreeNode(key, value, nil, nil, nil)
		tree.size++
	} else {
		// Insert below the tree's root
		tree.root = tree.insertWalk(tree.root, key, value)
	}
}

func (tree *TreeMap) insertWalk(node *treeNode, key, value interface{}) *treeNode {
	retNode := node
	dir := tree.cmpFn(key, node.key)
	if dir < 0 {
		if node.less == nil {
			node.less = newtreeNode(key, value, node, nil, nil)
			tree.size++
		} else {
			node.less = tree.insertWalk(node.less, key, value)
		}
		// Rebalance the tree after an insertion
		retNode = tree.skew(retNode)
		retNode = tree.split(retNode)
	} else if dir > 0 {
		if node.greater == nil {
			node.greater = newtreeNode(key, value, node, nil, nil)
			tree.size++
		} else {
			node.greater = tree.insertWalk(node.greater, key, value)
		}
		// Rebalance the tree after an insertion
		retNode = tree.skew(retNode)
		retNode = tree.split(retNode)
	} else {
		// Update existing item
		// No rebalancing is required
		node.value = value
	}
	return retNode
}

func (tree *TreeMap) Remove(key interface{}) {
	if tree.root != nil {
		tree.root = tree.removeWalk(tree.root, key)
	}
}

func (tree *TreeMap) removeWalk(node *treeNode, key interface{}) *treeNode {
	retNode := node
	dir := tree.cmpFn(key, node.key)
	if dir < 0 {
		if node.less != nil {
			node.less = tree.removeWalk(node.less, key)
			if node.less != nil {
				node.less.parent = node
			}
		}
	} else if dir > 0 {
		if node.greater != nil {
			node.greater = tree.removeWalk(node.greater, key)
			if node.greater != nil {
				node.greater.parent = node
			}
		}
	} else {
		if node.less != nil {
			// Find predecessor node
			delNode := node.less
			for delNode.greater != nil {
				delNode = delNode.greater
			}
			// Copy value and remove leaf
			node.key = delNode.key
			node.value = delNode.value
			node.less = tree.removeWalk(node.less, delNode.key)
			if node.less != nil {
				node.less.parent = node
			}
		} else if node.greater != nil {
			// Find successor node
			delNode := node.greater
			for delNode.less != nil {
				delNode = delNode.less
			}
			// Copy value and remove leaf
			node.key = delNode.key
			node.value = delNode.value
			node.greater = tree.removeWalk(node.greater, delNode.key)
			if node.greater != nil {
				node.greater.parent = node
			}
		} else {
			retNode = nil
			tree.size--
		}
	}

	if retNode != nil {
		// Adjust the balance level
		if retNode.less != nil || retNode.greater != nil {
			var maxLevel int
			if retNode.less != nil {
				maxLevel = retNode.less.level
			}
			if retNode.greater != nil && retNode.greater.level < maxLevel {
				maxLevel = retNode.greater.level
			}
			maxLevel++
			if retNode.level >= maxLevel {
				retNode.level = maxLevel
				if retNode.greater != nil && retNode.greater.level >= maxLevel {
					retNode.greater.level = maxLevel
				}
			}
		}

		// Rebalance the tree after a deletion
		retNode = tree.skew(retNode)
		if retNode.greater != nil {
			retNode.greater = tree.skew(retNode.greater)
			subNode := retNode.greater
			if subNode.greater != nil {
				subNode.greater = tree.skew(subNode.greater)
			}
		}
		retNode = tree.split(retNode)
		if retNode.greater != nil {
			retNode.greater = tree.split(retNode.greater)
		}
	}

	return retNode
}

func (tree *TreeMap) Get(key interface{}) (interface{}, bool) {
	var retValue interface{} = nil
	retNode, retFlag := tree.getWalk(tree.root, key)
	if retFlag {
		retValue = retNode.value
	}
	return retValue, retFlag
}

func (tree *TreeMap) getWalk(node *treeNode, key interface{}) (*treeNode, bool) {
	var retNode *treeNode = nil
	var retFlag bool = false
	if node != nil {
		dir := tree.cmpFn(key, node.key)
		if dir < 0 {
			retNode, retFlag = tree.getWalk(node.less, key)
		} else if dir > 0 {
			retNode, retFlag = tree.getWalk(node.greater, key)
		} else {
			retNode = node
			retFlag = true
		}
	}
	return retNode, retFlag
}

func (tree *TreeMap) GetFirstKey() (interface{}, bool) {
	var retKey interface{} = nil
	var retFlag bool = false
	node := tree.root
	if node != nil {
		for node.less != nil {
			node = node.less
		}
		retKey = node.key
		retFlag = true
	}
	return retKey, retFlag
}

func (tree *TreeMap) GetLastKey() (interface{}, bool) {
	var retKey interface{} = nil
	var retFlag bool = false
	node := tree.root
	if node != nil {
		for node.greater != nil {
			node = node.greater
		}
		retKey = node.key
		retFlag = true
	}
	return retKey, retFlag
}

func (tree *TreeMap) GetSize() int {
	return tree.size
}

func (tree *TreeMap) skew(node *treeNode) *treeNode {
	rotNode := node
	if node.less != nil && node.level == node.less.level {
		rotNode = node.less
		rotNode.parent = node.parent
		if node.parent != nil {
			if node.parent.less == node {
				node.parent.less = rotNode
			} else {
				node.parent.greater = rotNode
			}
		} else {
			tree.root = rotNode
		}
		node.less = rotNode.greater
		if node.less != nil {
			node.less.parent = node
		}
		rotNode.greater = node
		node.parent = rotNode
	}
	return rotNode
}

func (tree *TreeMap) split(node *treeNode) *treeNode {
	rotNode := node
	if node.greater != nil && node.greater.greater != nil &&
		node.level == node.greater.greater.level {
		rotNode = node.greater
		rotNode.parent = node.parent
		if node.parent != nil {
			if node.parent.less == node {
				node.parent.less = rotNode
			} else {
				node.parent.greater = rotNode
			}
		} else {
			tree.root = rotNode
		}
		node.greater = rotNode.less
		if node.greater != nil {
			node.greater.parent = node
		}
		rotNode.less = node
		node.parent = rotNode
		rotNode.level++
	}
	return rotNode
}
