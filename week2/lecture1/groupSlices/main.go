package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	groupSlices(citiesAndPrices())

}

func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100

	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)
	}

	return cities, prices
}
func groupSlices(keySlice []string, valueSlice []int) map[string][]int {

	keySliceAndValueSlice := make(map[string][]int)
	for k := range keySlice {
		values := make([]int, 4)
		for i := range values {
			valueSliceIndex := rand.Intn(100)
			values[i] = valueSlice[valueSliceIndex]
		}
		keySliceAndValueSlice[keySlice[k]] = values
	}

	for key, val := range keySliceAndValueSlice {
		fmt.Println(key, val)
	}

	return keySliceAndValueSlice

}
