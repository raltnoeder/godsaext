// VMap -- double ended queue implementation of a key/value map
//
// @version 2018-07-11
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

type vMapNode struct {
	key   interface{}
	value interface{}
	prev  *vMapNode
	next  *vMapNode
}

func newVMapNode(key, value interface{}, prev, next *vMapNode) *vMapNode {
	return &vMapNode{key, value, prev, next}
}

type VMap struct {
	head  *vMapNode
	tail  *vMapNode
	size  int
	cmpFn compareFn
}

type VMapIterator struct {
	nextNode *vMapNode
}

func (iter *VMapIterator) Next() (interface{}, interface{}, bool) {
	var key interface{} = nil
	var value interface{} = nil
	retFlag := false
	if iter.nextNode != nil {
		key = iter.nextNode.key
		value = iter.nextNode.value
		retFlag = true
		iter.nextNode = iter.nextNode.next
	}
	return key, value, retFlag
}

func NewVMap(cmpFn compareFn) *VMap {
	return &VMap{nil, nil, 0, cmpFn}
}

func (mapObj *VMap) Iterator() *VMapIterator {
	return &VMapIterator{mapObj.head}
}

func (mapObj *VMap) Clear() {
	mapObj.head = nil
	mapObj.tail = nil
	mapObj.size = 0
}

func (mapObj *VMap) GetSize() int {
	return mapObj.size
}

func (mapObj *VMap) Prepend(key, value interface{}) {
	node := newVMapNode(key, value, nil, mapObj.head)
	if mapObj.head != nil {
		mapObj.head.prev = node
	}
	mapObj.head = node
	if mapObj.tail == nil {
		mapObj.tail = node
	}
	mapObj.size++
}

func (mapObj *VMap) Append(key, value interface{}) {
	node := newVMapNode(key, value, mapObj.tail, nil)
	if mapObj.tail != nil {
		mapObj.tail.next = node
	}
	mapObj.tail = node
	if mapObj.head == nil {
		mapObj.head = node
	}
	mapObj.size++
}

func (mapObj *VMap) Get(key interface{}) (interface{}, bool) {
	var retValue interface{} = nil
	node, retFlag := mapObj.findNode(key)
	if retFlag {
		retValue = node.value
	}
	return retValue, retFlag
}

func (mapObj *VMap) GetFirst() (interface{}, interface{}, bool) {
	var retKey interface{} = nil
	var retValue interface{} = nil
	retFlag := false
	if mapObj.head != nil {
		retKey = mapObj.head.key
		retValue = mapObj.head.value
		retFlag = true
	}
	return retKey, retValue, retFlag
}

func (mapObj *VMap) GetLast() (interface{}, interface{}, bool) {
	var retKey interface{} = nil
	var retValue interface{} = nil
	retFlag := false
	if mapObj.tail != nil {
		retKey = mapObj.tail.key
		retValue = mapObj.tail.value
		retFlag = true
	}
	return retKey, retValue, retFlag
}

func (mapObj *VMap) Remove(key interface{}) {
	node, found := mapObj.findNode(key)
	if found {
		if mapObj.head == node {
			mapObj.head = node.next
		} else {
			node.prev.next = node.next
		}
		if mapObj.tail == node {
			mapObj.tail = node.prev
		} else {
			node.next.prev = node.prev
		}
		mapObj.size--
	}
}

func (mapObj *VMap) findNode(key interface{}) (*vMapNode, bool) {
	node := mapObj.head
	for node != nil {
		if mapObj.cmpFn(key, node.key) == 0 {
			break
		}
		node = node.next
	}
	return node, node != nil
}
