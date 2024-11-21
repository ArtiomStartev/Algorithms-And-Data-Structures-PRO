package search

import (
	"algorithms/utils"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// PerformanceCheckerResult holds the details of each individual performance check
type PerformanceCheckerResult struct {
	SliceSize     int
	Order         utils.OrderType
	Iteration     int
	ExecutionTime time.Duration
}

// PerformanceChecker struct to manage multiple runs and record results
type PerformanceChecker struct {
	Algorithm Searcher
	SliceSize int
	Order     utils.OrderType
	Runs      int
	Results   []PerformanceCheckerResult
}

// RunPerformanceCheck performs the performance check and records each iteration result
func (pc *PerformanceChecker) RunPerformanceCheck() time.Duration {
	var totalDuration time.Duration
	var failures int

	// Get the data and target for the search
	data, target := pc.getDataAndTarget()

	// Run the performance check for the specified number of iterations
	for i := 0; i < pc.Runs; i++ {
		startTime := time.Now()
		found := pc.Algorithm.Search(data, target)
		duration := time.Since(startTime)
		totalDuration += duration

		// Record result and check if target was found
		if !found {
			failures++
		}

		pc.Results = append(pc.Results, PerformanceCheckerResult{
			SliceSize:     pc.SliceSize,
			Order:         pc.Order,
			Iteration:     i + 1,
			ExecutionTime: duration,
		})
	}

	// Log any failures in finding the target
	if failures > 0 {
		fmt.Printf("Warning: %d/%d searches failed to find the target\n", failures, pc.Runs)
	}

	// Return average execution time
	return totalDuration / time.Duration(pc.Runs)
}

// getDataAndTarget generates the data and target for the performance check based on the algorithm type
func (pc *PerformanceChecker) getDataAndTarget() (any, int) {
	// Generate slice based on size and order type
	slice := utils.GenerateSlice(pc.SliceSize, pc.Order)

	// Select a random element from the slice as the target
	target := slice[rand.Intn(len(slice))]

	switch pc.Algorithm.(type) {
	case BinarySearchTreeSearcher:
		// Create a binary search tree from the slice
		return MapSliceToTree(slice), target
	default:
		// Return the slice and target for other algorithms
		return slice, target
	}
}

// WriteResultsToFile outputs both the average result and individual iteration results to a file
func (pc *PerformanceChecker) WriteResultsToFile(file *os.File) {
	averageTime := pc.RunPerformanceCheck()

	// Write the average result to file
	avgResult := fmt.Sprintf("Algorithm: %T | Slice Size: %d | Order: %s | Avg Time: %d ns\n",
		pc.Algorithm, pc.SliceSize, pc.Order, averageTime.Nanoseconds())

	if _, err := file.WriteString(avgResult); err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	// Write detailed results for each run to file
	for _, result := range pc.Results {
		detailedResult := fmt.Sprintf("Slice Size: %d | Order: %s | Iteration: %d | Time: %d ns\n",
			result.SliceSize, result.Order, result.Iteration, result.ExecutionTime.Nanoseconds())

		if _, err := file.WriteString(detailedResult); err != nil {
			fmt.Println("Error writing to file: ", err)
			continue
		}
	}
}
