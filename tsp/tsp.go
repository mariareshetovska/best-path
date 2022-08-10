package tsp

import (
	"best-path/utility"
	"fmt"
)

var permutation [][]int

const N int = 6

// slice of tops
var tops []int

func TspDuration(graph [][]int) (int, [][]int) {
	// temp subgraphs of tops
	var topCover [][]int
	var topCoverTemp [][]int

	// Total time path
	var currentTotalTime = 0

	// start from
	var minPath = 1000000000

	// Insert vertices of graph into array
	tops = createTopSlice()

	// Get all the permutation
	getPermutation(tops, nil)
	for _, node := range permutation {
		// Reset
		currentTotalTime = 0
		topCoverTemp = nil

		for i, n := 0, len(tops)-1; i < n; i++ {
			// Get arch of graph weight
			archWeight := graph[node[i]][node[i+1]]
			// Increment path weight
			currentTotalTime += archWeight
			// Push node into graph path
			topCoverTemp = append(topCoverTemp, []int{node[i], node[i+1], archWeight})
		}

		if currentTotalTime < minPath {
			minPath = currentTotalTime
			topCover = nil
			for i := range topCoverTemp {
				topCover = append(topCover, topCoverTemp[i])
			}
		}

	}
	return minPath, topCover
}

func TspPrice(graph [][]float64) (float64, [][]float64) {
	// temp subgraphs of tops
	var topCover [][]float64
	var topCoverTemp [][]float64

	// Total path weight
	var currentPriceWeight float64 = 0

	// start from
	var minPriceWeight float64 = 1000000000

	// Insert vertices of graph into array
	tops = createTopSlice()

	// Get all the permutation
	getPermutation(tops, nil)
	for _, node := range permutation {
		currentPriceWeight = 0
		topCoverTemp = nil

		for i, n := 0, len(tops)-1; i < n; i++ {
			//Get arch of graph weight
			archWeight := graph[node[i]][node[i+1]]
			// Increment path weight
			currentPriceWeight += archWeight
			// Push node into graph path
			topCoverTemp = append(topCoverTemp, []float64{float64(node[i]), float64(node[i+1]), archWeight})
		}

		if currentPriceWeight < minPriceWeight {
			minPriceWeight = currentPriceWeight
			topCover = nil
			for i := range topCoverTemp {
				topCover = append(topCover, topCoverTemp[i])
			}
		}

	}
	return minPriceWeight, topCover

}

func PrintResDuration(minPathWeight int, topCover [][]int, leaveSt []int, arriveSt []int) {
	fmt.Println("THE BEST PATH FOR THE DURATION WILL CONTINUE " + utility.ParseDurationToHour(minPathWeight) + " \n")
	for _, top := range topCover {
		fmt.Println("From station ", leaveSt[top[0]], " to station ", arriveSt[top[1]], "and takes ", utility.ParseDurationToHour(top[2]))

	}

}

func PrintResPrice(minPriceWeight float64, topCover [][]float64, leaveSt []int, arriveSt []int) {
	total := fmt.Sprintf("%.2f", minPriceWeight)
	fmt.Println("THE BEST PATH FOR THE PRICE WILL COST " + total)
	for _, top := range topCover {
		fmt.Println("From station ", leaveSt[int(top[0])], " to station ", arriveSt[int(top[1])], "and costs ", (top[2]))
	}
}

func createTopSlice() []int {
	var new []int
	for i := 0; i < N; i++ {
		new = append(new, i)
	}
	return new
}

func getPermutation(items []int, perms []int) []int {
	if len(items) == 0 {
		permutation = append(permutation, perms)
		return perms
	} else {
		var newPerms []int
		var newItems []int
		for i := len(items) - 1; i >= 0; i-- {
			newItems = items
			newPerms = perms
			f := arraySplice(&newItems, i)
			newPerms = append([]int{f}, newPerms...)
			getPermutation(newItems, newPerms)
		}
	}
	return nil
}

func arraySplice(items *[]int, pos int) int {
	var res int
	var newItems []int

	for i, item := range *items {
		if i == pos {
			res = item
		} else {
			newItems = append(newItems, item)
		}
		i++
	}
	*items = newItems
	return res
}
