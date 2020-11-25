package trie

import "sync"

type Trie struct {
	mu sync.Mutex
	t  *trie
}

type trie struct {
	m map[rune]*trie
}

func New() *Trie {
	return &Trie{sync.Mutex{}, newTrie()}
}

func newTrie() *trie {
	return &trie{make(map[rune]*trie)}
}

func (t *trie) LoadOrStore(x rune, n *trie) (*trie, bool) {
	next, ok := t.m[x]
	if !ok {
		t.m[x] = n
		return n, true
	}
	return next, false
}

func (t *trie) Load(key rune) (value *trie, ok bool) {
	v, ok := t.m[key]
	return v, ok
}

func (t *trie) Range(f func(key rune, value *trie) bool) {
	for k, v := range t.m {
		if f(k, v) {
			return
		}
	}
}

func (t *trie) Delete(key rune) {
	delete(t.m, key)
}

func (t *Trie) Insert(x string) {
	t.mu.Lock()
	t.t.insert([]rune(x))
	t.mu.Unlock()
}

func (t *trie) insert(x []rune) {
	if len(x) == 0 {
		return
	}
	next, _ := t.LoadOrStore(x[0], newTrie())
	if next != nil {
		if len(x) == 1 && x[0] != rune('\n') {
			next.insert([]rune{'\n'})
			return
		}
		next.insert(x[1:])
	}
}

func (t *Trie) Remove(x string) {
	t.mu.Lock()
	t.t.remove([]rune(x))
	t.mu.Unlock()
}

func (t *trie) remove(x []rune) {
	if len(x) == 0 {
		return
	}
	next, ok := t.Load(x[0])
	if !ok {
		return
	}
	if len(x) == 1 {
		next.Delete(rune('\n'))
		return
	}
	next.remove(x[1:])
}

func (t *Trie) PrefixSearch(x string) string {
	return string(t.t.prefixSearch([]rune{}, []rune(x)))
}

func (t *trie) isTerm() bool {
	_, ok := t.Load(rune('\n'))
	return ok
}

func (t *trie) next() (r rune, next *trie) {
	t.Range(func(key rune, value *trie) bool {
		r = key
		next = value
		return false
	})
	return r, next
}

func (t *trie) prefixSearch(agg []rune, x []rune) []rune {
	if len(x) == 0 {
		if t.isTerm() {
			return agg
		}
		r, next := t.next()
		if next != nil {
			return next.prefixSearch(append(agg, r), []rune{})
		}
		return []rune{}
	}
	next, ok := t.Load(x[0])
	if !ok {
		r, next := t.next()
		if next != nil {
			return next.prefixSearch(append(agg, r), x[1:])
		}
		return []rune{}
	}
	if next != nil {
		return next.prefixSearch(append(agg, x[0]), x[1:])
	}
	return []rune{}
}
