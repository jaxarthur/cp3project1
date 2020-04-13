package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		inputError()
	}

	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		inputError()
	}
	if size < 1 {
		inputError()
	}

	start := time.Now()

	pyrimid := genPyrimid(size)

	end := time.Now()
	elapsedNs := end.Sub(start).Nanoseconds()
	elapsedSec := float64(elapsedNs) * 1e-9

	printPyrimid(pyrimid)
	fmt.Printf("Elapsed Time: %d ns or %2.9f seconds\n", elapsedNs, elapsedSec)
}

func genPyrimid(size int) [][]float64 {
	arr := make([][]float64, 0)
	arr = append(arr, make([]float64, 0))
	arr[0] = append(arr[0], 200)
	return genPyrimidLoop(size-1, arr, 1, 0)
}

func genPyrimidLoop(size int, data [][]float64, x int, y int) [][]float64 {
	if x > y {
		if y > size-1 {
			return data
		}
		return genPyrimidLoop(size, data, 0, y+1)
	}

	if x == 0 {
		data = append(data, make([]float64, 0))
	}

	val := 200.0
	if x > 0 {
		val += data[y-1][x-1] / 2
	}

	if x < y {
		val += data[y-1][x] / 2
	}

	data[y] = append(data[y], val)
	return genPyrimidLoop(size, data, x+1, y)

}

func printPyrimid(data [][]float64) {
	for _, y := range data {
		for _, x := range y {
			fmt.Printf("%-8.2f", x-200)
		}
		fmt.Println()
	}
}

func inputError() {
	fmt.Println("Must enter non zero postive int as the first command line argument!!")
	os.Exit(1)
}
