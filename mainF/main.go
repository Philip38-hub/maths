package main

import (
	"bufio"
	"fmt"
	"mathskill"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Oops! No file to read.")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Unable to Open file!")
		return
	}
	scanner := bufio.NewScanner(file) //read file line by line
	var data []float64
	for scanner.Scan() {
		// if len(scanner.Text()) > 10 { // check if values are within range
		// 	fmt.Printf("Value %s is out of range! Ensure every values is 10 digits or fewer and try again!\n", scanner.Text())
		// 	return
		// }
		val, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("invalid input format. Check data set and try again!")
			return
		}
		data = append(data, val)
	}
	if len(data) == 0 { //if the file is empty
		fmt.Println("Oops! No values to work with!")
		return
	}
	// %.0f will concatenate decimal points and dispay larger values
	fmt.Printf("Average: %.0f\n", mathskill.Average(data))                  // avarage
	fmt.Printf("Median: %.0f\n", mathskill.Median(data))                 // median
	fmt.Printf("Variance: %.0f\n", mathskill.Variance(data))            // variance
	fmt.Printf("Standard Deviation: %.0f\n", mathskill.StandardDev(data)) // standard deviation
}
