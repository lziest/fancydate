package fancydate

import (
	"time"
)

func Anniversary(d time.Time, years int) time.Time {
	yy := d.Year()
	mm := d.Month()
	dd := d.Day()

	ret := time.Date(yy+years, mm, dd, 0, 0, 0, 0, d.Location())
	if ret.Day() != dd {
		// date gets normalized, skip it: for example,
		// Feb 29th 2016 doesn't get an anniversary in 2017
		var z time.Time
		return z
	}
	return ret
}

func SimpleAnniversaries(d time.Time) (ret []time.Time) {
	for i := 1; i < 100; i += 1 {
		ann := Anniversary(d, i)
		if !ann.IsZero() {
			ret = append(ret, Anniversary(d, i))
		}
	}
	return ret

}

func AfterManyZeros(d time.Time) (ret []time.Time) {
	for i := 1; i <= 50; i += 1 {
		ret = append(ret, AfterDays(d, i*100))
	}
	return ret
}

func AfterLuckyNumbers(d time.Time) (ret []time.Time) {
	var luckyNumbers = []int{168, 666, 888, 1688, 1888, 6666, 6688, 8888}
	for _, n := range luckyNumbers {
		ret = append(ret, AfterDays(d, n))
	}
	return ret
}

func AfterDays(d time.Time, days int) time.Time {
	return d.AddDate(0, 0, days)

}
