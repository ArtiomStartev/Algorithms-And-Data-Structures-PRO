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
	ArrSize       int
	Order         utils.OrderType
	Iteration     int
	ExecutionTime time.Duration
}

// PerformanceChecker struct to manage multiple runs and record results
type PerformanceChecker struct {
	Algorithm Searcher
	ArrSize   int
	Order     utils.OrderType
	Runs      int
	Results   []PerformanceCheckerResult
}

// RunPerformanceCheck performs the performance check and records each iteration result
func (pc *PerformanceChecker) RunPerformanceCheck() time.Duration {
	var totalDuration time.Duration
	var failures int

	for i := 0; i < pc.Runs; i++ {
		data, x := pc.getDataAndTarget()

		startTime := time.Now()
		found := pc.Algorithm.Search(data, x)
		duration := time.Since(startTime)
		totalDuration += duration

		if !found {
			failures++
		}

		pc.Results = append(pc.Results, PerformanceCheckerResult{
			ArrSize:       pc.ArrSize,
			Order:         pc.Order,
			Iteration:     i + 1,
			ExecutionTime: duration,
		})
	}

	if failures > 0 {
		fmt.Printf("Warning: %d/%d searches failed to find the target\n", failures, pc.Runs)
	}

	return totalDuration / time.Duration(pc.Runs)
}

// getDataAndTarget generates the data and target for the performance check based on the algorithm type
func (pc *PerformanceChecker) getDataAndTarget() (any, int) {
	arr := utils.GenerateArr(pc.ArrSize, pc.Order)
	x := arr[rand.Intn(len(arr))]

	switch pc.Algorithm.(type) {
	case BinarySearchTreeSearcher:
		return MapSliceToTree(utils.Unique(arr)), x
	default:
		return arr, x
	}
}

// WriteResultsToFile outputs both the average result and individual iteration results to a file
func (pc *PerformanceChecker) WriteResultsToFile(file *os.File) {
	averageTime := pc.RunPerformanceCheck()

	avgResult := fmt.Sprintf("Algorithm: %s | Array Size: %d | Order: %s | Avg Time: %f ms\n",
		getAlgorithmName(pc.Algorithm), pc.ArrSize, pc.Order, float64(averageTime.Nanoseconds())/1e6)

	if _, err := file.WriteString(avgResult); err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	// Write detailed results for each run to file
	for _, result := range pc.Results {
		detailedResult := fmt.Sprintf("Array Size: %d | Order: %s | Iteration: %d | Time: %f ms\n",
			result.ArrSize, result.Order, result.Iteration, float64(result.ExecutionTime.Nanoseconds())/1e6)

		if _, err := file.WriteString(detailedResult); err != nil {
			fmt.Println("Error writing to file: ", err)
			continue
		}
	}
}

// getAlgorithmName returns the name of the algorithm based on its type
func getAlgorithmName(alg Searcher) string {
	switch alg.(type) {
	case LinearSearcher:
		return "Linear Search"
	case BinarySearcher:
		return "Binary Search"
	case ExponentialSearcher:
		return "Exponential Search"
	case BinarySearchTreeSearcher:
		return "Binary Search Tree"
	default:
		return "Unknown Algorithm"
	}
}
