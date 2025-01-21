package main

import (
	"fmt"
	"log"
)

// bst graph
// graph condaint collection of vertex called vertces

type graph struct {
	vertces []*vertex
}

// vertex have adjuscent that means they have connection between the vertex
type vertex struct {
	val int

	adjuscent []*vertex
}

// insert into graph
func (g *graph) insertGraph(val int) {

	if !g.contain(val) {

		g.vertces = append(g.vertces, &vertex{val: val})

	}

}

func (g *graph) getVertex(val int) *vertex {

	for _, ver := range g.vertces {

		if ver.val == val {

			return ver
		}
	}

	return nil

}

func (g *graph) displayGraph() {

	for _, ver := range g.vertces {

		fmt.Print(ver.val)

		for _, val := range ver.adjuscent {

			fmt.Print(" ", val.val)
		}

		println()

	}

}

func (g *graph) addEdges(from, to int) {

	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if !g.contain(fromVertex.val) || !g.contain(toVertex.val) {

		log.Println("can't make edge's, no vertex is there in th graph ")

		return
	}

	for _, val := range fromVertex.adjuscent {

		if val == toVertex {

			log.Println("already make edge's")
			return
		}
	}

	fromVertex.adjuscent = append(fromVertex.adjuscent, toVertex)
	toVertex.adjuscent = append(toVertex.adjuscent, fromVertex)

}

// checking the vertex contain in the graph ..
func (g *graph) contain(v int) bool {

	for _, ver := range g.vertces {

		if ver.val == v {

			return true

		}
	}

	return false

}

func main() {

	g := graph{}

	g.insertGraph(10)
	g.insertGraph(20)
	g.insertGraph(30)
	g.insertGraph(40)
	g.insertGraph(50)
	g.insertGraph(60)
	g.insertGraph(70)
	g.insertGraph(80)
	g.insertGraph(90)

	g.addEdges(10, 20)
	g.addEdges(10, 60)
	g.addEdges(10, 70)

	g.addEdges(20, 40)
	g.addEdges(20, 50)
	g.addEdges(20, 60)

	g.dfs(10)
}

type queue struct {
	arr []*vertex
}

func (g *graph) bfs(key int) {

	q := &queue{}

	seen := make(map[int]bool)

	start := g.getVertex(key)

	if start == nil {

		log.Println("there is no vertex to start ..")
		return
	}

	seen[start.val] = true
	q.arr = append(q.arr, start)

	for len(q.arr) != 0 {

		vertex := q.arr[0]

		q.arr = q.arr[1:]

		fmt.Print("vertex val :", vertex.val, " ")

		for _, near := range vertex.adjuscent {

			fmt.Print(" ", near.val)

			if !seen[near.val] {

				q.arr = append(q.arr, near)
				seen[near.val] = true
			}
		}

		fmt.Println()

	}

}

type stack struct {
	arr []*vertex
}

func (g *graph) dfs(key int) {

	s := &stack{}

	start := g.getVertex(key)

	seen := make(map[int]bool)

	s.arr = append(s.arr, start)

	seen[start.val] = true

	fmt.Print("vertex ", start.val)

	for len(s.arr) != 0 {

		ver := s.arr[len(s.arr)-1]

		s.arr = s.arr[:len(s.arr)-1]

		for _, near := range ver.adjuscent {

			if !seen[near.val] {

				fmt.Print(" ", near.val)

				s.arr = append(s.arr, near)
				seen[near.val] = true
			}
		}

		fmt.Println()

	}

}
