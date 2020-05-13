package main

import "fmt"

type LRUCache struct {
	cap		int
	now		int
	keys	*Node
	values	map[int]int
}

//下述结构效率高很多
/*type LinkNode struct {
	key,value int
	next,prev *LinkNode
}

type LRUCache struct {
	m map[int] *LinkNode
	cap int
	head,tail *LinkNode
}*/

//执行结果：通过 显示详情
//执行用时 :304 ms, 在所有 Go 提交中击败了5.13%的用户
//内存消耗 :12 MB, 在所有 Go 提交中击败了100.00%的用户  想哭。。。
func leetcode146() {
	l := Constructor(2)
	l.Put(2, 1)
	fmt.Println(l.Get(2))
	l.Put(3,2)
	l.keys.Print()
	fmt.Println(l.Get(2))
	l.keys.Print()
	fmt.Println(l.Get(3))
}

func Constructor(capacity int) LRUCache {
	o := LRUCache{}
	o.cap = capacity
	o.values = make(map[int]int)
	return o
}

func (l *LRUCache) Get(key int) int {
	o,ok := l.values[key]
	if !ok {
		return -1
	}
	tmp := l.keys
	if tmp.Value == key {
		return o
	}
	var before, next *Node
	for tmp.Next != nil {
		if tmp.Value == key {
			next = tmp.Next
			break
		}
		before = tmp
		tmp = tmp.Next
	}
	if tmp != nil {
		tmp.Next = l.keys
		if before != nil {
			before.Next = next
		}
		l.keys = tmp
	}
	return o
}

func (l *LRUCache) Put(key int, value int) {
	_, ok := l.values[key]
	l.values[key] = value
	if ok {
		tmp := l.keys
		var before, next *Node
		if tmp.Value == key {
			return
		}
		for tmp != nil {
			if tmp.Value == key {
				next = tmp.Next
				break
			}
			before = tmp
			tmp = tmp.Next
		}
		if tmp != nil {
			tmp.Next = l.keys
			if before != nil {
				before.Next = next
			}
			l.keys = tmp
		}
	}else{
		l.now++
		l.keys = &Node{key, l.keys}
		if l.now > l.cap {
			l.now = l.cap
			tmp := l.keys
			for i:=0; i<l.cap-1; i++{
				tmp = tmp.Next
			}
			c := tmp.Next
			for c != nil {
				delete(l.values, c.Value)
				c = c.Next
			}
			tmp.Next = nil
		}
	}
}