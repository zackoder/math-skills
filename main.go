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
	if len(os.Args) != 2 ||  os.Args[1] != "data.txt" {
		fmt.Println("Run the command:\n\tgo run main.go data.txt\nto get the expected result")
		return 
	}

	fileName := os.Args[1]
	numbers := readAndRetData(fileName)
	avr := Average(numbers)
	variance := Variance(numbers, int(avr))

	fmt.Println("Average:", avr)
	fmt.Println("Median:", Median(numbers))
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", StandardDeviation(variance))
}



func Average(nums []float64) int {
	var sum float64
	for _, val := range nums {
		sum += val
	}
	avr := sum / float64(len(nums)) // Calculate the average
	return int(math.Round(avr))     // Round and return as int
}

// Median calculates the median of a sorted list of float64 numbers and returns it as an integer.

func Median(nums []float64) int {
	sort.Float64s(nums) // Sort the slice
	n := len(nums)
	if n%2 == 0 {
		// For an even number of elements, average the two middle elements.
		return int(math.Round((nums[n/2] + nums[n/2-1]) / 2))
	}
	// For an odd number of elements, return the middle element.
	return int(math.Round(nums[n/2]))
}

// Variance calculates the variance based on the provided numbers and their average.
func Variance(nums []float64, avr int) int {
	mean := float64(avr) // Convert average to float64 for accurate variance calculation.
	var sum float64
	for _, n := range nums {
		sum += math.Round(math.Pow(n-mean, 2))
	}
	variance := sum / float64(len(nums)) // Divide the sum by the number of elements.
	return int(math.Round(variance))     // Return the rounded variance.
}

// StandardDeviation calculates the square root of the variance.
func StandardDeviation(variance int) int {
	return int(math.Round(math.Sqrt(float64(variance)))) // Return the square root of the variance.
}


func readAndRetData(fileName string) []float64 {
	
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err.Error())
	}

	dataSplited := strings.Split(string(data), "\n")

	var numbers []float64
	
	for _, v := range dataSplited {

		if v == "" {
			continue
		}

		num, err := strconv.ParseFloat(v, 64)

		if err == nil {
			numbers = append(numbers, num)
		} else {
			fmt.Println(err)
			fmt.Printf("This value is not a number (%s)\n", v)
			os.Exit(1)
		}
	}
	return numbers
}