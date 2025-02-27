package main

import (
	"fmt"
	"sort"
)

// source: https://gist.github.com/dammyng/c76f937faf80dbe4ddac64de9063a509
type Pair struct {
	Key   string
	Value int
}

func (p Pair) String() string { return fmt.Sprintf("(%s,%d)", p.Key, p.Value) }

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortedByValue(m map[string]int) PairList {
	p := make(PairList, len(m))

	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	return p
}
