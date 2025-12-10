package datastructures

import "fmt"

type DisjointSet[T comparable] struct {
	parent map[T]T
	Count  int
}

func NewDisjointSet[T comparable]() *DisjointSet[T] {
	return &DisjointSet[T]{
		parent: make(map[T]T, 0),
		Count:  0,
	}
}

func (ds *DisjointSet[T]) Find(item T) T {
	for ds.parent[item] != item {
		ds.parent[item] = ds.parent[ds.parent[item]] // Path compression
		item = ds.parent[item]
	}
	return item
}

func (ds *DisjointSet[T]) Union(item1, item2 T) bool {
	root1 := ds.Find(item1)
	root2 := ds.Find(item2)

	if root1 != root2 {
		ds.parent[root1] = root2
		ds.Count--
		return true
	}

	return false
}

func (ds *DisjointSet[T]) Add(item T) {
	if _, exists := ds.parent[item]; !exists {
		ds.parent[item] = item
		ds.Count++
	}
}

func (ds *DisjointSet[T]) String() string {
	result := "DisjointSet:\n"

	for item, parent := range ds.parent {
		result += fmt.Sprintf("  %v -> %v\n", item, parent)
	}

	result += fmt.Sprintf("\nTotal Groups: %d\n", ds.Count)
	for k, v := range ds.Groups() {
		result += fmt.Sprintf("  %v -> %v\n", k, v)
	}
	return result
}

func (ds *DisjointSet[T]) Groups() map[T][]T {
	groups := make(map[T][]T)
	for item := range ds.parent {
		root := ds.Find(item)
		groups[root] = append(groups[root], item)
	}
	return groups
}
