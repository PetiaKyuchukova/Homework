package main

import (
	"fmt"
	"sort"
	"time"
)

func sortDates(format string, dates ...string) ([]string, error) {
	var timeSlice []time.Time
	for _, dateStr := range dates {
		time, err := time.Parse(format, dateStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		timeSlice = append(timeSlice, time)
	}

	sort.Slice(timeSlice, func(i, j int) bool { return timeSlice[i].Before(timeSlice[j]) })

	for i := 0; i < len(timeSlice); i++ {
		dates[i] = timeSlice[i].Format(format)
	}
	return dates, nil

}
func main() {
	dates := []string{"Dec-03-2021", "Mar-18-2022", "Sep-14-2008"}
	layout := "Jan-02-2006"
	fmt.Print(sortDates(layout, dates...))
}
