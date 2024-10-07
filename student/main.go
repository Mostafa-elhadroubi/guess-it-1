package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	// "strings"
)


// func DataTraitment(s string) []float64 {
// 	arrData := strings.Split(s, "\n")
// 	var arrInt = []float64{}
// 	arrData = arrData[:len(arrData)-1]
// 	for i := 0; i < len(arrData); i++ {
// 		num, _:= strconv.Atoi(arrData[i])
// 		arrInt = append(arrInt, float64(num))
// 	}
// 	return arrInt
// }

func Average(arrInt []float64) float64 {
	sum := 0.0
	for i := 0; i < len(arrInt); i++ {
		sum += arrInt[i]
	}
	res := math.Round((sum) / float64(len(arrInt)))
	return res
}



func Variance(arrInt []float64) float64 {
	var variance, average float64
	average = Average(arrInt)
	for i := 0; i < len(arrInt); i++ {
		variance += ((arrInt[i])-average)*((arrInt[i])-average)
	}
	res := (variance/ float64(len(arrInt)))
	return res
}

func Deviation(arrData []float64) float64 {
	variance := Variance(arrData)
	deviation := math.Sqrt((variance))
	return deviation
}
func Result(arrData []float64) (int, int) {
	i := len(arrData)-10
	if len(arrData) == 0 {
		return 0, 0
	}
	newSlice := arrData[i:]
	average := Average(newSlice)
	deviation := Deviation(newSlice)
	lower :=  int(math.Round((average - deviation)))
	limit := int(math.Round((average + deviation)))
	// fmt.Println(lower, limit)
	return lower, limit
}

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
		lower, limit := Result(arrData)
		fmt.Printf("%d %.d\n", lower, limit)
	}

}