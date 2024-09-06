package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Validate input arguments
	if len(os.Args) != 2 || os.Args[1] != "data.txt" {
		fmt.Println("Run the command:\n\tgo run main.go data.txt\nto get the expected result")
		return
	}

	fileName := os.Args[1]
	// Read numbers from the file
	numbers := readAndRetData(fileName)

	if len(numbers) == 0 {
		fmt.Println("the data file is empty")
		os.Exit(1)
	}

	avr := Average(numbers)
	variance := Variance(numbers, avr)

	
	fmt.Println("Average:", int(math.Round(avr)))
	fmt.Println("Median:", int(math.Round(Median(numbers))))
	fmt.Println("Variance:", int(math.Round(variance)))
	fmt.Println("Standard Deviation:", int(math.Round(StandardDeviation(variance))))
}

func Average(nums []float64) float64 {
	var sum float64
	for _, val := range nums {
		sum += val
	}
	// Calculate and return average without rounding at this step
	avr := sum / float64(len(nums))
	return avr
}

// Median calculates the median of a sorted list of float64 numbers and returns it as a float64
func Median(nums []float64) float64 {
	sort.Float64s(nums) // Sort the slice
	n := len(nums)
	if n%2 == 0 {
		// For an even number of elements, return the average of the two middle elements
		return (nums[n/2] + nums[n/2-1]) / 2
	}
	// For an odd number of elements, return the middle element
	return nums[n/2]
}

// Variance calculates the variance based on the provided numbers and their average
func Variance(nums []float64, avr float64) float64 {
	var sum float64
	for _, n := range nums {
		sum += (n - avr) * (n - avr)
	}
	
	return sum / float64(len(nums))
}

// StandardDeviation calculates the square root of the variance
func StandardDeviation(variance float64) float64 {
	// Return the square root of the variance
	return math.Sqrt(variance)
}

// reads the data from the file and return a slice of floats
func readAndRetData(fileName string) []float64 {
	
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	dataSplited := strings.Split(string(data), "\n")

	var numbers []float64
	
	for i, v := range dataSplited {

		if v == "" {
			continue
		}

		num, err := strconv.ParseFloat(v, 64)

		if err == nil {
			numbers = append(numbers, num)
		} else {
			fmt.Println(err)
			fmt.Printf("This value is not a number (%s) in the file \"data.txt\" line %d \n", v, i+1)
			os.Exit(1)
		}
	}
	return numbers
}