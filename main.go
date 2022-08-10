package main

import (
	"best-path/data"
	"best-path/tsp"
	"best-path/utility"
	"fmt"
	"sort"
)

const infInt int = 100000000

const infFl float64 = 100000000

func main() {
	csvLines, err := data.ReadFile("./test_task_data.csv")
	if err != nil {
		fmt.Println(err)
	}
	data := data.MapData(csvLines)

	// get array of unique sorted stations
	var leaveSt []int
	var arriveSt []int
	for i := range data {
		leaveSt = append(leaveSt, data[i].LeaveSt)
		arriveSt = append(arriveSt, data[i].ArriveSt)
	}
	uniquLeaveSt := utility.Unique(leaveSt)
	sort.Ints(uniquLeaveSt)
	uniquArriveSt := utility.Unique(arriveSt)
	sort.Ints(uniquArriveSt)

	// create matrix of time
	matrixDuration := make([][]int, 6)
	for i := range matrixDuration {
		matrixDuration[i] = make([]int, 6)
	}
	var durations []int
	for j := 0; j < len(uniquArriveSt); j++ {
		for k := 0; k < len(uniquLeaveSt); k++ {
			durations = nil
			for i := 0; i < len(data); i++ {
				if data[i].LeaveSt == uniquLeaveSt[k] && data[i].ArriveSt == uniquArriveSt[j] {
					durations = append(durations, data[i].TravelDuration)
					x := utility.FindMin(durations)
					matrixDuration[k][j] = x
				}
			}
		}
	}
	// fill empty values with infinity
	utility.FillInfInt(matrixDuration, infInt)

	// create matrix of prices
	matrixPrice := make([][]float64, 6)
	for i := range matrixPrice {
		matrixPrice[i] = make([]float64, 6)
	}
	var prices []float64
	for j := 0; j < len(uniquArriveSt); j++ {
		for k := 0; k < len(uniquLeaveSt); k++ {
			prices = nil
			for i := 0; i < len(data); i++ {
				if data[i].LeaveSt == uniquLeaveSt[k] && data[i].ArriveSt == uniquArriveSt[j] {
					prices = append(prices, data[i].Price)
					x := utility.FindMinFl(prices)
					matrixPrice[k][j] = x
				}
			}
		}
	}
	// fill empty values with infinity
	utility.FillInfFl(matrixPrice, infFl)

	for i := range matrixDuration {
		fmt.Println(matrixDuration[i])
	}

	// best path for duration
	minPathWeight, vertexCoverDur := tsp.TspDuration(matrixDuration)
	tsp.PrintResDuration(minPathWeight, vertexCoverDur, uniquLeaveSt, uniquArriveSt)
	fmt.Println("------------------------")
	// best path for price
	minPriceWeight, vertexCoverPr := tsp.TspPrice(matrixPrice)
	tsp.PrintResPrice(minPriceWeight, vertexCoverPr, uniquLeaveSt, uniquArriveSt)
}
