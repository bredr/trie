package trie

//go:generate go run github.com/a8m/syncmap -name Trie -pkg trie "map[rune]*Trie"

func (t *Trie) Insert(x string) {
	t.insert([]rune(x))
}

func (t *Trie) insert(x []rune) {
	if len(x) == 0 {
		return
	}
	next, _ := t.LoadOrStore(x[0], &Trie{})
	if next != nil {
		if len(x) == 1 && x[0] != rune('\n') {
			next.insert([]rune{'\n'})
			return
		}
		next.insert(x[1:])
	}
}

func (t *Trie) Remove(x string) {
	t.remove([]rune(x))
}

func (t *Trie) remove(x []rune) {
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
	return string(t.prefixSearch([]rune{}, []rune(x)))
}

func (t *Trie) isTerm() bool {
	_, ok := t.Load(rune('\n'))
	return ok
}

func (t *Trie) next() (r rune, next *Trie) {
	t.Range(func(key rune, value *Trie) bool {
		r = key
		next = value
		return false
	})
	return r, next
}

func (t *Trie) prefixSearch(agg []rune, x []rune) []rune {
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
	}
	if next != nil {
		return next.prefixSearch(append(agg, x[0]), x[1:])
	}
	return []rune{}
}
