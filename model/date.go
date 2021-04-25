package model

type Date struct {
	year  int `db:"year"  json:"year"`
	month int `db:"month" json:"month"`
	day   int `db:"day"   json:"day"`

	hour   int `db:"hour"   json:"hour"`
	minute int `db:"minute" json:"minute"`
	second int `db:"second" json:"second"`
}

func (date *Date) IsBefore(other Date) bool {
	predicate := func(left int, right int) bool {
		return left < right
	}

	return date.compare(other, predicate)
}

func (date *Date) IsAfter(other Date) bool {
	predicate := func(left int, right int) bool {
		return left > right
	}

	return date.compare(other, predicate)
}

func (date *Date) Equals(other Date) bool {
	predicate := func(left int, right int) bool {
		return left == right
	}

	return date.compare(other, predicate)
}

func (date *Date) compare(other Date, predicate func(int, int) bool) bool {
	thisArray := date.toArray()
	otherArray := other.toArray()

	for i, thisElem := range thisArray {
		otherElem := otherArray[i]

		if predicate(thisElem, otherElem) {
			return true
		}
	}

	return false
}

func (date *Date) toArray() []int {
	return []int{date.year, date.month, date.day, date.hour, date.minute, date.second}
}
