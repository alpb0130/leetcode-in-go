package dfs

import (
	"strconv"
	"strings"
)

type Vertex struct {
	Val       int
	Neighbors []*Vertex
}

func NewVertex(val int) *Vertex {
	return &Vertex{val, []*Vertex{}}
}

//func DeepCopyGraph2(vertex *Vertex) *Vertex {
//	if vertex == nil {
//		return nil
//	}
//
//	v := &Vertex{Val: vertex.Val}
//	isVisited := map[int]*Vertex{v.Val: v}
//	copyHelper2(vertex, v, isVisited)
//	return v
//}
//func copyHelper2(old, new *Vertex, isVisited map[int]*Vertex) {
//	for _, neighbor := range old.Neighbors {
//		newNeighbor, ok := isVisited[neighbor.Val]
//		if !ok {
//			newNeighbor = &Vertex{Val: neighbor.Val}
//			new.Neighbors = append(new.Neighbors, newNeighbor)
//			isVisited[newNeighbor.Val] = newNeighbor
//			copyHelper2(neighbor, newNeighbor, isVisited)
//		} else {
//			new.Neighbors = append(new.Neighbors, newNeighbor)
//		}
//	}
//}

func DeepCopyGraph(vertex *Vertex) *Vertex {
	return copyHelper(vertex, map[int]*Vertex{})
}

func copyHelper(old *Vertex, vertexMap map[int]*Vertex) *Vertex {
	if old == nil {
		return nil
	}
	if v, ok := vertexMap[old.Val]; ok {
		return v
	}
	new := &Vertex{Val: old.Val}
	vertexMap[new.Val] = new
	for _, neighbor := range old.Neighbors {
		new.Neighbors = append(new.Neighbors, copyHelper(neighbor, vertexMap))
	}
	return new
}

func Encode(vertex *Vertex) string {
	if vertex == nil {
		return ""
	}
	return encodeHelper(vertex, map[int]bool{})
}

func encodeHelper(vertex *Vertex, isVisited map[int]bool) string {
	var str strings.Builder
	isVisited[vertex.Val] = true
	str.WriteString("{ Val: " + strconv.Itoa(vertex.Val) + ", ")
	str.WriteString("Neighbors: [")
	for i, neighbor := range vertex.Neighbors {
		str.WriteString(strconv.Itoa(neighbor.Val))
		if i != len(vertex.Neighbors)-1 {
			str.WriteString(", ")
		}
	}
	str.WriteString("] }")
	for _, neighbor := range vertex.Neighbors {
		str.WriteString(", ")
		if !isVisited[neighbor.Val] {
			str.WriteString(encodeHelper(neighbor, isVisited))
		}
	}
	return str.String()
}
