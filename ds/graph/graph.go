package main

import "fmt"

type grpaph struct {
	vertices []*vertex
}

type vertex struct {
	val int

	adajucent []*vertex
}

// insert to grpah

func (g *grpaph) insertGraph(val int) {

	if !contains(g.vertices, val) {

		g.vertices = append(g.vertices, &vertex{val: val})

	}

}

//check vertex is inside the grpah

func contains(vertices []*vertex, num int) bool {

	for _, v := range vertices {

		if v.val == num {

			return true
		}
	}

	return false
}

// make edges between them

func (g *grpaph) addEdges(from, to int) {

	fromVertex := GetVertex(g.vertices, from)
	toVertex := GetVertex(g.vertices, to)

	if fromVertex == nil || toVertex == nil {

		fmt.Println("there is no vertex for connection ")
		return
	}

	for _, val := range fromVertex.adajucent {

		if val.val == to {

			fmt.Println("duplicate spoted")
			return
		}
	}

	fromVertex.adajucent = append(fromVertex.adajucent, toVertex)
	toVertex.adajucent = append(toVertex.adajucent, fromVertex)
}

// get the vertex

func GetVertex(vertces []*vertex, value int) *vertex {

	for _, val := range vertces {

		if val.val == value {
			return val
		}
	}

	return nil
}

func (g *grpaph) DisplayGraph() {

	for _, vertices := range g.vertices {

		fmt.Print(vertices.val, " ")

		for _, ver := range vertices.adajucent {

			fmt.Print(" ", ver.val)

		}

		println()
	}

}

func main() {

	g := grpaph{}

	g.insertGraph(10)
	g.insertGraph(20)
	g.insertGraph(30)
	g.insertGraph(40)
	g.insertGraph(50)
	g.insertGraph(60)
	g.insertGraph(70)

	g.addEdges(10, 20)
	g.addEdges(10, 50)

	g.addEdges(40, 30)

	g.bfs(10)

}

// BFS
// Is stand for breath first search tree
// fifo first in first out

type queue struct {
	arr []*vertex
}

func (g *grpaph) bfs(key int) {

	q := &queue{}

	vertex := GetVertex(g.vertices, key)

	q.arr = append(q.arr, vertex)
	seen := make(map[int]bool)

	for len(q.arr) != 0 {

		vertex := q.arr[0]

		q.arr = q.arr[1:]

		if seen[vertex.val] {
			continue
		}
		seen[vertex.val] = true

		fmt.Println(vertex.val, " ")

		for _, ver := range vertex.adajucent {

			if !seen[ver.val] {

				fmt.Println(ver.val)
				seen[ver.val] = true
				q.arr = append(q.arr, ver)

			}

		}

	}
	println()

}
