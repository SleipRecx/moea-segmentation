package main

// TODO: Improve running time using "union by rank" and "path compression" heuristics.
type Element struct {
	Parent *Element
	Value  interface{}
}

func makeSet(value interface{}) *Element {
	e := new(Element)
	e.Parent = e
	e.Value = value
	return e
}

func findSet(e *Element) *Element {
	if e == e.Parent {
		return e
	}
	e.Parent = findSet(e.Parent)
	return e.Parent
}

func union(e1, e2 *Element) {
	root1 := findSet(e1)
	root2 := findSet(e2)
	root1.Parent = root2
}
