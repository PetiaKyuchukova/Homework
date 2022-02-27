package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please, enter the month value (in the range 1-12): ")
	scanner.Scan()
	month, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Please, enter the year: ")
	scanner.Scan()
	year, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	output_days, output_bool := daysInMonth(month, year)
	nameOfMonth := nameOfMonth(month)
	switch nameOfMonth {
	case "Invalid input!":
		break

	default:
		switch output_bool {
		case true:
			fmt.Println("The year", year, "is a leap year. Days of month", nameOfMonth, "are", output_days, ".")
		case false:
			fmt.Println("The year", year, "is not a leap year. Days of month", nameOfMonth, "are", output_days, ".")
		}
	}

}

func nameOfMonth(month int64) string {
	switch month {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "Octomber"
	case 11:
		return "November"
	case 12:
		return "December"

	default:
		return "Invalid input!"
	}
}
func daysInMonth(month int64, year int64) (int, bool) {

	switch year % 4 {
	case 0:
		{
			switch month {
			case 1, 3, 5, 7, 8, 10, 12:
				return 31, true
			case 2:
				return 29, true
			case 4, 6, 9, 11:
				return 30, true
			default:
				fmt.Printf("Invalid input! Please, enter month in the range 1-12!")
				return 0, true
			}

		}
	default:
		{
			switch month {
			case 1, 3, 5, 7, 8, 10, 12:
				return 31, false
			case 2:
				return 28, false
			case 4, 6, 9, 11:
				return 30, false
			default:
				fmt.Printf("Invalid input! Please, enter month in the range 1-12!")
				return 0, true
			}
		}
	}
}
