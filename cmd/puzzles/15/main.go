package main

import (
	"advent-of-code-2021/pkg/util"
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"strconv"
	"strings"
	"time"
)

func main() {
	matrix := readMatrix()

	part1Result, part1Duration := part1(matrix)
	fmt.Printf("Part 1: %4d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2(matrix)
	fmt.Printf("Part 2: %4d (duration: %s)\n", part2Result, part2Duration)
}

func part1(matrix [][]int) (int, time.Duration) {
	start := time.Now()
	graph := convertToGraph(matrix)
	lowestRiskLevel := determineLowestRiskLevel(graph)
	return lowestRiskLevel, time.Since(start)
}

func part2(matrix [][]int) (int, time.Duration) {
	start := time.Now()
	resizedMatrix := increaseMatrix(matrix)
	graph := convertToGraph(resizedMatrix)
	lowestRiskLevel := determineLowestRiskLevel(graph)
	return lowestRiskLevel, time.Since(start)
}

func increaseMatrix(matrix [][]int) [][]int {
	const ResizeFactor = 5

	lenY := len(matrix)
	lenX := len(matrix[0])
	resizedMatrix := make([][]int, lenY*ResizeFactor)
	for y := 0; y < lenY*ResizeFactor; y++ {
		resizedMatrix[y] = make([]int, lenX*ResizeFactor)
	}

	for yf := 0; yf < ResizeFactor; yf++ {
		for xf := 0; xf < ResizeFactor; xf++ {
			for y := 0; y < lenY; y++ {
				for x := 0; x < lenX; x++ {
					value := matrix[y][x] + yf + xf
					if value > 9 {
						value -= 9
					}
					resizedMatrix[y+yf*lenY][x+xf*lenX] = value
				}
			}
		}
	}
	return resizedMatrix
}

func convertToGraph(matrix [][]int) graph.Graph {
	directedGraph := simple.NewWeightedDirectedGraph(0, 0)

	lenY := len(matrix)
	lenX := len(matrix[0])
	nodes := make([][]graph.Node, lenY)
	for y, row := range matrix {
		nodes[y] = make([]graph.Node, lenX)
		for x, _ := range row {
			node := directedGraph.NewNode()
			directedGraph.AddNode(node)
			nodes[y][x] = node
		}
	}

	for y, row := range nodes {
		for x, _ := range row {
			// edge 1 to down: x,y <-> x,y+1
			if y+1 < lenY {
				downByOne := directedGraph.NewWeightedEdge(nodes[y][x], nodes[y+1][x], float64(matrix[y+1][x]))
				directedGraph.SetWeightedEdge(downByOne)
				backByOne := directedGraph.NewWeightedEdge(nodes[y+1][x], nodes[y][x], float64(matrix[y][x]))
				directedGraph.SetWeightedEdge(backByOne)
			}
			// edge 1 to right: x,y <-> x+1,y
			if x+1 < lenX {
				rightByOne := directedGraph.NewWeightedEdge(nodes[y][x], nodes[y][x+1], float64(matrix[y][x+1]))
				directedGraph.SetWeightedEdge(rightByOne)
				backByOne := directedGraph.NewWeightedEdge(nodes[y][x+1], nodes[y][x], float64(matrix[y][x]))
				directedGraph.SetWeightedEdge(backByOne)
			}
		}
	}
	return directedGraph
}

func determineLowestRiskLevel(g graph.Graph) int {
	start := g.Node(0)
	end := g.Node(int64(g.Nodes().Len() - 1))
	weight := path.DijkstraFrom(start, g).WeightTo(end.ID())
	return int(weight)
}

func readMatrix() [][]int {
	lines := util.ReadFile("15")

	matrix := make([][]int, len(lines))
	for y, line := range lines {
		digits := strings.Split(line, "")
		row := make([]int, len(digits))
		for x, char := range digits {
			digit, _ := strconv.Atoi(char)
			row[x] = digit
		}
		matrix[y] = row
	}
	return matrix
}
