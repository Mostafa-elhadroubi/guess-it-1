package functions

import (
	"math"
)

func Average(arrInt []float64) float64 {
	sum := 0.0
	for _,  num := range arrInt {
		sum += num
	}
	res := float64(sum) / float64(len(arrInt))
	roundRes := math.Round(res)
	return math.Round(roundRes)
}



func Variance(arrInt []float64) float64 {
	var variance float64
	average := Average(arrInt)
	for  _, num := range arrInt {
		variance += (num-average)*(num-average)
	}
	return (variance/ float64(len(arrInt)))
}

func Deviation(arrData []float64) float64 {
	return math.Sqrt(Variance(arrData))
}
func Result(arrData []float64) (int, int) {
	i := len(arrData)-8
	if i < 0 {
		i = 0
	}
	newSlice := arrData[i:]
	average := Average(newSlice)
	deviation := Deviation(newSlice)
	lower :=  average - (0.8*deviation)
	limit := average + (0.8*deviation)
	return int(math.Round(lower)), int(math.Round(limit))
}