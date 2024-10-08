package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"guessit1/student/functions"
)





func main() {
	scanner := bufio.NewScanner(os.Stdin)
	arrData := []float64{}
	for scanner.Scan() {
		s := scanner.Text()
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Println("Invalid input!!!!")
			continue
		}
		arrData =append(arrData, num)
		if len(arrData) > 1 {
			lower, limit := functions.Result(arrData)
			fmt.Printf("%d %d\n", lower, limit)
		}
	}

}