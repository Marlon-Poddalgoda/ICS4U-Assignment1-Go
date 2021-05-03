// This program calculates the day of the week
// Created By: Marlon Poddalgoda
// Created on: 2021-05-01

package main

import (
	"fmt"
	"math"
)

func DayOfTheWeek(yearInt int, monthInt int, dayInt int) string {
	// this function calculates the day of the week

	// variables
	var weekDay float64 = 0
	var weekDayInt int
	var century float64
	var yearInteger int
	var weekName string
	var monthFloat float64 = float64(monthInt)
	var yearFloat float64 = float64(yearInt)
	var dayFloat float64 = float64(dayInt)

	// constants
	const january, february int = 11, 12
	const two, four float64 = 2, 4
	const weekLength float64 = 7
	const sunday, monday, tuesday, wednesday int = 0, 1, 2, 3
	const thursday, friday, saturday int = 4, 5, 6

	// constants for repetitive patterns
	const monthPattern1 float64 = 2.6
	const monthPattern2 float64 = 0.2

	switch {
	case monthInt == january || monthInt == february:
		yearInteger = yearInt - 1
	default:
		yearInteger = yearInt + 0
	}

	century = yearFloat / 100

	var year float64 = float64(yearInteger % 10)

	weekDay = (dayFloat + (monthPattern1*monthFloat - monthPattern2) -
		two*century + year + (year / four) + (century / 4))

	weekDay = math.Mod(weekDay, weekLength)

	switch {
	case weekDay < 0:
		weekDay += weekLength
	default:
		weekDay += 0
	}

	weekDayInt = int(weekDay)

	switch {
	case weekDayInt == sunday:
		weekName = "Sunday"
	case weekDayInt == monday:
		weekName = "Monday"
	case weekDayInt == tuesday:
		weekName = "Tuesday"
	case weekDayInt == wednesday:
		weekName = "Wednesday"
	case weekDayInt == thursday:
		weekName = "Thursday"
	case weekDayInt == friday:
		weekName = "Friday"
	case weekDayInt == saturday:
		weekName = "Saturday"
	default:
		fmt.Println(weekDayInt)
	}

	return (weekName)
}

func main() {
	// this function takes in user input

	fmt.Println("This program calculates the day of the week for a certain date.")
	fmt.Println("")

	// variables
	var dayOfWeek string
	var yearInt int
	var monthInt int
	var dayInt int

	// constants
	const minYear, maxYear int = 1000, 9999
	const minMonth, maxMonth int = 1, 12
	const minDay, maxDay = 1, 31
	// constants used to set month values
	const march int = 3
	const two int = 2
	const ten int = 10

	fmt.Print("Enter a year: ")

	// check user input
	for {
		// user input
		_, err := fmt.Scanf("%d", &yearInt)

		// if error, then restart loop
		switch {
		case err != nil:
			// output
			fmt.Println("That's not a year, try again!")
			fmt.Print("Enter a year: ")
		case yearInt <= minYear || yearInt > maxYear:
			// output
			fmt.Println("That is out of the range, try again!")
			fmt.Print("Enter a year: ")
		default:
			// error catch
			break
		}

		if yearInt > minYear && yearInt < maxYear {
			break
		}
	}

	fmt.Print("Enter a month (01-12): ")

	// check user input
	for {
		// user input
		_, err := fmt.Scanf("%d", &monthInt)

		// if error, then restart loop
		switch {
		case err != nil:
			// output
			fmt.Println("That's not a month, try again!")
			fmt.Print("Enter a month (01-12): ")
		case monthInt < minMonth || monthInt > maxMonth:
			// output
			fmt.Println("That is out of the range, try again!")
			fmt.Print("Enter a month (01-12): ")

		default:
			break
		}

		if monthInt >= minMonth && monthInt <= maxMonth {
			switch {
			case monthInt >= march:
				monthInt -= two
			case monthInt < march:
				monthInt += ten
			default:
				monthInt += 0
			}
			break
		}
	}

	fmt.Print("Enter a day (01 - 31): ")

	// check user input
	for {
		// user input
		_, err := fmt.Scanf("%d", &dayInt)

		// if error, then restart loop
		switch {
		case err != nil:
			// output
			fmt.Println("That's not a day, try again!")
			fmt.Print("Enter a day (01 - 31): ")
		case dayInt < minDay || dayInt > maxDay:
			// output
			fmt.Println("That is out of the range, try again!")
			fmt.Print("Enter a day (01 - 31): ")
		default:
			break
		}

		if dayInt >= minDay && dayInt <= maxDay {
			break
		}
	}

	// call function
	dayOfWeek = DayOfTheWeek(yearInt, monthInt, dayInt)

	fmt.Println("")
	fmt.Println(dayOfWeek)

	// program closes
	fmt.Println("")
	fmt.Println("Done.")
}
