package utils


type Searchable[T any] interface {
	CheckGoal() bool
	Extend() []T
}


func BFS[S Searchable[S]](s S) (S, int) {
/*  Breadth-First-Search from:
	https://en.wikipedia.org/wiki/Breadth-first_search

	implements Pseudocode:
	1  procedure BFS(G, root) is
	2      let Q be a queue
	3      label root as explored
	4      Q.enqueue(root)
	5      while Q is not empty do
	6          v := Q.dequeue()
	7          if v is the goal then
	8              return v
	9          for all edges from v to w in G.adjacentEdges(v) do
	10              if w is not labeled as explored then
	11                  label w as explored
	12                  w.parent := v
	13                  Q.enqueue(w)

	Except:
		- we do not label visoted state as explored
		- we do track the depth of search, and return it
*/
	Q := Queue[S]{s}
	I := Queue[int]{0}
	for !Q.IsEmpty() {
		v, _ := Q.Dequeue()
		i, _ := I.Dequeue()
		if v.CheckGoal() {
			return v, i
		}
		for _, w := range v.Extend() {
			Q.Enqueue(w)
			I.Enqueue(i+1)
		}
	}
	var zero S
	return zero, -1
}

func AllSolutions[S Searchable[S]](s S) []S {
	// BFS based
	Q := Queue[S]{s}
	I := Queue[int]{0}
	var solutions []S
	for !Q.IsEmpty() {
		v, _ := Q.Dequeue()
		i, _ := I.Dequeue()
		if v.CheckGoal() {
			solutions = append(solutions, v)
		}
		for _, w := range v.Extend() {
			Q.Enqueue(w)
			I.Enqueue(i+1)
		}
	}
	return solutions
}