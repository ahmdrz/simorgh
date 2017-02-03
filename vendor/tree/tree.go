package tree

import (
	"tree/radix"
)

func NewTree() *Tree {
	return &Tree{
		tree: radix.New(),
	}
}

func (t *Tree) Del(key string) {
	t.tree.Delete(key)
}

func (t *Tree) Get(key string) (interface{}, string) {
	br, exists, mode := t.tree.Get(key)
	if exists {
		return br, mode
	}
	return "UNDEFIEND", radix.UNDEFIEND
}

func (t *Tree) Clr() int {
	m := t.tree.ToMap()
	for key := range m {
		t.tree.Delete(key)
	}
	return len(m)
}

func (t *Tree) Set(key string, value interface{}) {
	_, ok := t.tree.Set(key, value)
	if !ok {
		t.tree.Insert(key, value)
	}
}
