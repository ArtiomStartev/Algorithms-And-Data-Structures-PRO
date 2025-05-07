package main

import (
	"algorithms/search_algorithms/search"
	"algorithms/utils"
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	algorithms := []search.Searcher{
		search.LinearSearcher{},
		search.BinarySearcher{},
		search.ExponentialSearcher{},
		search.BinarySearchTreeSearcher{},
	}

	arrSizes := []int{100000, 1000000, 10000000, 100000000, 1000000000}
	orders := []utils.OrderType{utils.Ascending, utils.Descending, utils.Random}

	// Open or create a file to store the results
	fileName := "./search_algorithms/results/performance_results_" + time.Now().Format("2006-01-02_15:04:05") + ".txt"
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
				// Skip binary search and exponential search for descending and random orders
				if reflect.TypeOf(alg) == reflect.TypeOf(search.BinarySearcher{}) || reflect.TypeOf(alg) == reflect.TypeOf(search.ExponentialSearcher{}) {
					if order == utils.Descending || order == utils.Random {
						continue
					}
				}

				pc := search.PerformanceChecker{
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
