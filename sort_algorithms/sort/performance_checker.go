package sort

import (
	"algorithms/utils"
	"fmt"
	"os"
	"time"
)

// PerformanceCheckerResult holds the details of each individual performance check
type PerformanceCheckerResult struct {
	SliceSize     int
	Order         utils.OrderType
	Iteration     int
	Comparisons   int
	Swaps         int
	ExecutionTime time.Duration
}

// PerformanceChecker struct to manage multiple runs and record results
type PerformanceChecker struct {
	Algorithm Sorter
	SliceSize int
	Order     utils.OrderType
	Runs      int
	Results   []PerformanceCheckerResult
}

// RunPerformanceCheck performs the performance check and records each iteration result
func (pc *PerformanceChecker) RunPerformanceCheck() {
	// Run the performance check for the specified number of iterations
	for i := 0; i < pc.Runs; i++ {
		// Generate a new slice for each iteration
		slice := utils.GenerateSlice(pc.SliceSize, pc.Order)

		// Sort the slice and record metrics
		result := pc.Algorithm.Sort(slice)

		pc.Results = append(pc.Results, PerformanceCheckerResult{
			SliceSize:     pc.SliceSize,
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
	// Run the performance check
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

	// Write main record with averages
	mainRecord := fmt.Sprintf(
		"Algorithm: %T | Slice Size: %d | Order: %s | Avg Comparisons: %d | Avg Swaps: %d | Avg Time: %d ns\n",
		pc.Algorithm, pc.SliceSize, pc.Order, avgComparisons, avgSwaps, avgTime.Nanoseconds(),
	)

	if _, err := file.WriteString(mainRecord); err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	// Write results for each iteration
	for _, result := range pc.Results {
		iterationRecord := fmt.Sprintf(
			"Iteration: %d | Comparisons: %d | Swaps: %d | Time: %d ns\n",
			result.Iteration, result.Comparisons, result.Swaps, result.ExecutionTime.Nanoseconds(),
		)

		if _, err := file.WriteString(iterationRecord); err != nil {
			fmt.Println("Error writing to file: ", err)
			return
		}
	}
}
