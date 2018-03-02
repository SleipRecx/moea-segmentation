package graph

type Element struct {
	Parent *Element
	Rank   int
	Value  interface{}
}

func MakeSet(value interface{}) *Element {
	e := new(Element)
	e.Parent = e
	e.Rank = 0
	e.Value = value
	return e
}

func FindSet(e *Element) *Element {
	if e == e.Parent {
		return e
	}
	e.Parent = FindSet(e.Parent)
	return e.Parent
}

func Union(e1, e2 *Element) {
	root1 := FindSet(e1)
	root2 := FindSet(e2)

	if root1 == root2 {
		return
	}

	if root1.Rank > root2.Rank {
		root2.Parent = root1
	} else if root1.Rank < root2.Rank {
		root1.Parent = root2
	} else {
		root1.Parent = root2
		root2.Rank = root2.Rank + 1
	}

}
