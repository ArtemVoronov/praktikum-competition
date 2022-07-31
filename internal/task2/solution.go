package task2

import (
	"fmt"
	"log"
	"time"
)

const (
	WEEK            string = "WEEK"            // неделя с понедельника по воскресенье.
	MONTH           string = "MONTH"           // месяц
	QUARTER         string = "QUARTER"         // интервалы в три месяца: январь — март, апрель — июнь, июль — сентябрь,  октябрь — декабрь.
	YEAR            string = "YEAR"            // год c 1 января по 31 декабря.
	FRIDAY_THE_13TH string = "FRIDAY_THE_13TH" // интервал с пятницы 13-го по ближайший четверг 12-го.
)

type Period struct {
	Start string
	End   string
}

const suffix string = "T00:00:00+00:00"
const formatTemplate string = "2006-01-02"

func parse(dateIsoFormatted string) time.Time {
	result, err := time.Parse(time.RFC3339, dateIsoFormatted+suffix)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
func format(date time.Time) string {
	return date.Format(formatTemplate)
}

func getAddingTime(now time.Time, unit string) (time.Duration, error) {
	switch unit {
	case WEEK:
		return getAddingTimeForWeek(now)
	case MONTH:
		return getAddingTimeForMonth(now)
	case QUARTER:
		return getAddingTimeForQuarter(now)
	case YEAR:
		return getAddingTimeForYear(now)
	case FRIDAY_THE_13TH:
		return getAddingTimeForFriday13(now)
	default:
		return 0, fmt.Errorf("unknown unit: %s", unit)
	}
}

func getAddingTimeForWeek(now time.Time) (time.Duration, error) {
	weekday := now.Weekday()
	switch weekday {
	case time.Sunday:
		return 0, nil
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday:
		return time.Duration((7-weekday)*24) * time.Hour, nil
	default:
		return 0, fmt.Errorf("unkown weekday: %v", weekday)
	}
}

func getAddingTimeForMonth(now time.Time) (time.Duration, error) {
	day := now.Day()
	month := now.Month()
	year := now.Year()
	switch month {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return time.Duration((31-day)*24) * time.Hour, nil
	case time.April, time.June, time.September, time.November:
		return time.Duration((30-day)*24) * time.Hour, nil
	case time.February:
		daysInFebruary := 28
		if year%4 == 0 {
			daysInFebruary = 29
		}
		return time.Duration((daysInFebruary-day)*24) * time.Hour, nil
	default:
		return 0, fmt.Errorf("unkown month: %v", month)
	}
}

func getAddingTimeForQuarter(now time.Time) (time.Duration, error) {
	toAdd := time.Duration(0)
	month := now.Month()
	t := now.Add(toAdd)
	switch month {
	case time.January, time.February, time.March:
		for t.Month() != time.March || t.Day() != 31 {
			toAdd += time.Duration(24) * time.Hour
			t = now.Add(toAdd)
		}
	case time.April, time.May, time.June:
		for t.Month() != time.June || t.Day() != 30 {
			toAdd += time.Duration(24) * time.Hour
			t = now.Add(toAdd)
		}
	case time.July, time.August, time.September:
		for t.Month() != time.September || t.Day() != 30 {
			toAdd += time.Duration(24) * time.Hour
			t = now.Add(toAdd)
		}
	case time.October, time.November, time.December:
		for t.Month() != time.December || t.Day() != 31 {
			toAdd += time.Duration(24) * time.Hour
			t = now.Add(toAdd)
		}
	default:
		return 0, fmt.Errorf("unkown month: %v", month)
	}

	return toAdd, nil
}

func getAddingTimeForYear(now time.Time) (time.Duration, error) {
	toAdd := time.Duration(0)
	t := now.Add(toAdd)
	for t.Month() != time.December || t.Day() != 31 {
		toAdd += time.Duration(24) * time.Hour
		t = now.Add(toAdd)
	}
	return toAdd, nil
}

func getAddingTimeForFriday13(now time.Time) (time.Duration, error) {
	weekday := now.Weekday()
	toAdd := time.Duration(0)
	switch weekday {
	case time.Sunday, time.Monday, time.Tuesday, time.Wednesday:
		toAdd = time.Duration((4-weekday)*24) * time.Hour
	case time.Thursday:
		toAdd = time.Duration(0)
	case time.Friday:
		toAdd = time.Duration(6*24) * time.Hour
	case time.Saturday:
		toAdd = time.Duration(5*24) * time.Hour
	default:
		return 0, fmt.Errorf("unkown weekday: %v", weekday)
	}

	thurday := now.Add(toAdd)

	for thurday.Day() != 12 {
		toAdd += time.Duration(7*24) * time.Hour
		thurday = now.Add(toAdd)
	}

	return toAdd, nil
}

func print(t time.Time) {
	fmt.Printf("date: %v, year: %v, month: %v, weekday: %v\n", format(t), t.Year(), t.Month(), t.Weekday())
}

func Solution(inputPeriodType string, inputPeriod Period) ([]Period, error) {
	var result []Period

	start := parse(inputPeriod.Start)
	end := parse(inputPeriod.End)

	var t time.Time = start
	for t.Before(end) {

		// fmt.Println("periodStart:")
		// print(t)
		period := Period{Start: format(t)}

		toAdd, err := getAddingTime(t, inputPeriodType)
		if err != nil {
			log.Fatal(err)
		}

		if t.Add(toAdd).After(end) {
			t = end
			// fmt.Println("periodEnd:")
			// print(t)
			period.End = format(t)
			result = append(result, period)
			break
		}

		t = t.Add(toAdd)

		// fmt.Println("periodEnd:")
		// print(t)
		period.End = format(t)
		result = append(result, period)

		t = t.Add(time.Hour * 24)

	}

	return result, nil
}
