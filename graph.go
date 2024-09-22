package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// graph
type graph struct {
	vextexes []*vertex
}
type vertex struct {
	key      string
	adjacent []*vertex
	visited  bool
}

// vertex
func (g *graph) Add(k string) {
	if g.Contains(k) {
		err := fmt.Errorf("vertex %v already created", k)
		fmt.Println(err.Error())
	} else {
		g.vextexes = append(g.vextexes, &vertex{key: k})
	}
}

// contains
func (g *graph) Contains(k string) bool {
	for i := 0; i < len(g.vextexes); i++ {
		if g.vextexes[i].key == k {
			return true
		}
	}
	return false
}

// print
func (g *graph) Print() {
	for i := 0; i < len(g.vextexes); i++ {
		fmt.Println("this is a vertex: ", g.vextexes[i].key)
		for j := 0; j < len(g.vextexes[i].adjacent); j++ {
			fmt.Printf("%v ", g.vextexes[i].adjacent[j].key)
		}
		fmt.Println()
	}
}

// add adjacent
func (g *graph) Addedge(from, to string) {
	if from == to {
		fmt.Printf("Cannot Add Edge: Vertex %v To Itself\n", from)
		return
	}
	fromvertex := g.Getv(from)
	tovertex := g.Getv(to)
	if fromvertex == nil || tovertex == nil {
		fmt.Printf("Cannot Add Edge: Vertex %v Or %v Does Not Exist\n", from, to)
		return
	} else if fromvertex.Has(to) {
		err := fmt.Errorf("%v-->%v Is Already An Existing Edge", from, to)
		fmt.Println(err.Error())
	} else {
		fromvertex.adjacent = append(fromvertex.adjacent, tovertex)
	}
	if !tovertex.Has(from) {
		g.Addedge(to, from)
	}
}

// edge checker
func (v *vertex) Has(to string) bool {
	for _, x := range v.adjacent {
		if x.key == to {
			return true
		}
	}
	return false
}

// get vertex
func (g *graph) Getv(k string) *vertex {
	for _, x := range g.vextexes {
		if x.key == k {
			return x
		}
	}
	return nil
}

func Filetograph(filename []byte) (*graph, string, string, int) {
	graph := &graph{}
	start := 0
	end := 0
	first := strings.Split(string(filename), "\n")
	ants, err := strconv.Atoi(first[0])
	if err != nil {
		ants = 0
	}
	for i := 1; i < len(first); i++ {
		if first[i] == "##start" || first[i] == "##end" {
			if first[i] == "##start" {
				start = i + 1
			} else if first[i] == "##end" {
				end = i + 1
			}
		}
		sample := strings.Split(first[i], " ")
		if len(sample) == 3 {
			graph.Add(sample[0])
		} else if len(sample) == 1 {
			test := strings.Split(first[i], "-")
			if len(test) == 2 {
				graph.Addedge(test[0], test[1])
			}
		}
	}
	startvertex := strings.Split(first[start], " ")[0]
	endvertex := strings.Split(first[end], " ")[0]
	return graph, startvertex, endvertex, ants
}

func (g *graph) Stoe(s, e, path string, paths *[]string) {
	start := g.Getv(s)
	end := g.Getv(e)
	start.visited = true
	path += start.key + "->"
	if start.key == end.key {
		*paths = append(*paths, path[:len(path)-2])
	} else {
		for _, v := range start.adjacent {
			if !v.visited {
				g.Stoe(v.key, end.key, path, paths)
			}
		}
	}
	start.visited = false
	path = path[:len(path)-1]
}

func (g *graph) Paths(s, e string) []string {
	var paths []string
	g.Stoe(s, e, "*", &paths)
	return paths
}

func SortByLength(strings []string) {
	sort.Slice(strings, func(i, j int) bool {
		return len(strings[i]) < len(strings[j])
	})
}

func main() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("An Err Has Occured While Reading The File")
		return
	}
	graph, start, end, _ := Filetograph(file)
	x := graph.Paths(start, end)
	SortByLength(x)
	fmt.Println(x)
}
