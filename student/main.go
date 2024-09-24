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
	res := math.Round(float64(sum) / float64(len(arrInt)))
	return res
}



func Variance(arrInt []float64) int64 {
	var variance, average float64
	average = Average(arrInt)
	for i := 0; i < len(arrInt); i++ {
		variance += (float64(arrInt[i])-average)*(float64(arrInt[i])-average)/ float64(len(arrInt))
	}
	res := int64(math.Round(variance))
	return res
}

func Deviation(arrData []float64) float64 {
	variance := Variance(arrData)
	deviation := (math.Round(math.Sqrt(float64(variance))))
	return deviation
}
func Result(arrData []float64) (int64, int64) {
	average := Average(arrData)
	deviation := Deviation(arrData)
	lower :=  int64(average - deviation)
	limit := int64(average + deviation)
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
		fmt.Println(lower, limit)
	}

}