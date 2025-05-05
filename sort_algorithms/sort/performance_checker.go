package sort

import (
	"algorithms/utils"
	"fmt"
	"os"
	"time"
)

// PerformanceCheckerResult holds the details of each individual performance check
type PerformanceCheckerResult struct {
	ArrSize       int
	Order         utils.OrderType
	Iteration     int
	Comparisons   int
	Swaps         int
	ExecutionTime time.Duration
}

// PerformanceChecker struct to manage multiple runs and record results
type PerformanceChecker struct {
	Algorithm Sorter
	ArrSize   int
	Order     utils.OrderType
	Runs      int
	Results   []PerformanceCheckerResult
}

// RunPerformanceCheck performs the performance check and records each iteration results
func (pc *PerformanceChecker) RunPerformanceCheck() {
	for i := 0; i < pc.Runs; i++ {
		arr := utils.GenerateArr(pc.ArrSize, pc.Order)
		result := pc.Algorithm.Sort(arr)

		pc.Results = append(pc.Results, PerformanceCheckerResult{
			ArrSize:       pc.ArrSize,
			Order:         pc.Order,
			Iteration:     i + 1,
			Comparisons:   result.Comparisons,
			Swaps:         result.Swaps,
			ExecutionTime: result.ExecutionTime,
		})
	}
}

// WriteResultsToFile writes the performance check results to a file
func (pc *PerformanceChecker) WriteResultsToFile(file *os.File) {
	pc.RunPerformanceCheck()

	var totalComparisons, totalSwaps int
	var totalTime time.Duration

	for _, result := range pc.Results {
		totalComparisons += result.Comparisons
		totalSwaps += result.Swaps
		totalTime += result.ExecutionTime
	}

	avgComparisons := totalComparisons / pc.Runs
	avgSwaps := totalSwaps / pc.Runs
	avgTime := totalTime / time.Duration(pc.Runs)

	mainRecord := fmt.Sprintf(
		"Algorithm: %T | Array Size: %d | Order: %s | Avg Comparisons: %d | Avg Swaps: %d | Avg Time: %d ns\n",
		pc.Algorithm, pc.ArrSize, pc.Order, avgComparisons, avgSwaps, avgTime.Nanoseconds(),
	)

	if _, err := file.WriteString(mainRecord); err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	for _, result := range pc.Results {
		iterationRecord := fmt.Sprintf(
			"Iteration: %d | Comparisons: %d | Swaps: %d | Time: %d ns\n",
			result.Iteration, result.Comparisons, result.Swaps, result.ExecutionTime.Nanoseconds(),
		)

		if _, err := file.WriteString(iterationRecord); err != nil {
			fmt.Println("Error writing to file: ", err)
			continue
		}
	}
}
