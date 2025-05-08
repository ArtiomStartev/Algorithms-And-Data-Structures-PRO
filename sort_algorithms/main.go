package main

import (
	"algorithms/sort_algorithms/sort"
	"algorithms/utils"
	"fmt"
	"os"
	"time"
)

func main() {
	algorithms := []sort.Sorter{
		sort.SelectionSorter{},
		sort.BubbleSorter{},
		sort.QuickSorter{},
		sort.MergeSorter{},
	}

	arrSizes := []int{1000, 10000, 100000, 250000}
	orders := []utils.OrderType{utils.Ascending, utils.Descending, utils.Random}

	// Open or create a file to store the results
	fileName := "./sort_algorithms/results/performance_results_" + time.Now().Format("2006-01-02_15:04:05") + ".txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()

	// Run tests for each algorithm, order size and combination
	for _, alg := range algorithms {
		for _, size := range arrSizes {
			for _, order := range orders {
				pc := sort.PerformanceChecker{
					Algorithm: alg,
					ArrSize:   size,
					Order:     order,
					Runs:      5, // Number of runs for averaging
				}
				pc.WriteResultsToFile(file)
			}
		}
	}
}
