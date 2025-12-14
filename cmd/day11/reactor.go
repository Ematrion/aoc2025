package main

import (
	"fmt"

	"strings"
	"aoc2025/utils"
)


func LoadReactors(lines []string) utils.Graph {
	var reactors utils.Graph
	for _, line := range lines {
		nodes := strings.Split(line, " ")
		v1, _ := strings.CutSuffix(nodes[0], ":")
		reactors.Vertices = append(reactors.Vertices, v1)
		for _, v2 := range nodes[1:] {
			reactors.Edges = append(reactors.Edges, utils.Pair[string]{v1, v2})
		}
	}
	return reactors
}


type PathSearch struct {
	graph utils.Graph
	from string
	to string
	path []utils.Pair[string]
}

func NewPathSearch(graph utils.Graph, from, to string) PathSearch {
	var paths []utils.Pair[string]
	return PathSearch{graph, "you", "out", paths}
}

func (ps *PathSearch) CheckGoal() bool {
	if len(ps.path) == 0 {
		return false
	}
	return ps.path[0][0] == ps.from && ps.path[len(ps.path)-1][1] == ps.to
}

func (ps *PathSearch) firstSearch() []*PathSearch {
	var options []*PathSearch
	for _, edge := range ps.graph.Edges {
		if edge[0] == ps.from {
			option := PathSearch{ps.graph, ps.from, ps.to, []utils.Pair[string]{edge}}
			options = append(options, &option)
		}
	}
	return options
}

func (ps *PathSearch) deepSearch() []*PathSearch {
	var options []*PathSearch
	for _, edge :=  range ps.graph.Edges {
		if edge[0] == ps.path[len(ps.path)-1][1] {
			new_path := append(ps.path, edge)
			if ps.graph.IsPath(new_path) {
				option := PathSearch{ps.graph, ps.from, ps.to, new_path}
				options = append(options, &option)
			}
		}
	}
	return options
}

func (ps *PathSearch) Extend() []*PathSearch {
	if len(ps.path) == 0 {
		return ps.firstSearch()
	} else {
		return ps.deepSearch()
	}
}



func main() {
	lines, err := utils.ReadFileToLines("inputs/day11.txt")
	utils.CheckError(err)
	reactors := LoadReactors(lines)
	//fmt.Println(reactors)
	
	fromYouToOut := NewPathSearch(reactors, "you", "out")
	solutions := utils.AllSolutions(&fromYouToOut)
	fmt.Println(len(solutions))

	/*for _, solution := range solutions {
		fmt.Println((*solution).path)
	}*/
}