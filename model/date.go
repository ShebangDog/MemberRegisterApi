package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Date struct {
	year  int
	month int
	day   int

	hour   int
	minute int
	second int
}

func ParseDate(str string) (Date, error) {
	dateTime := strings.Split(str, "-")
	dateStr := strings.Split(dateTime[0], "/")
	timeStr := strings.Split(dateTime[1], ":")

	stringsToInts := func(strings []string) ([]int, error) {
		var dates []int
		for _, str := range strings {
			parseInt, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return nil, err
			}

			dates = append(dates, int(parseInt))
		}

		return dates, nil
	}

	dates, _ := stringsToInts(dateStr)
	times, _ := stringsToInts(timeStr)

	result := Date{
		year:   dates[0],
		month:  dates[1],
		day:    dates[2],
		hour:   times[0],
		minute: times[1],
		second: times[2],
	}

	return result, nil
}

func (d *Date) ToString() string {
	result := fmt.Sprintf(
		`%02d/%02d/%02d-%02d:%02d:%02d`,
		d.year,
		d.month,
		d.day,
		d.hour,
		d.minute,
		d.second,
	)

	return result
}

func (d *Date) IsBefore(other Date) bool {
	predicate := func(left int, right int) bool {
		return left < right
	}

	return d.compare(other, predicate)
}

func (d *Date) IsAfter(other Date) bool {
	predicate := func(left int, right int) bool {
		return left > right
	}

	return d.compare(other, predicate)
}

func (d *Date) Equals(other Date) bool {
	predicate := func(left int, right int) bool {
		return left == right
	}

	return d.compare(other, predicate)
}

func (d *Date) compare(other Date, predicate func(int, int) bool) bool {
	thisArray := d.toArray()
	otherArray := other.toArray()

	for i, thisElem := range thisArray {
		otherElem := otherArray[i]

		if predicate(thisElem, otherElem) {
			return true
		}
	}

	return false
}

func (d *Date) toArray() []int {
	return []int{d.year, d.month, d.day, d.hour, d.minute, d.second}
}
