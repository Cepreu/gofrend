package vars

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

const timeRe = "^((0?[1-9])|(1[0-2])):([0-5][0-9]) ?(AM|PM|am|pm)$|(([01][0-9])|(2[0-3])):([0-5][0-9]$)"

func vuStringToMinutes(strValue string) (int, error) {
	var minutes int64
	r, _ := regexp.Compile(timeRe)
	res := r.FindAllStringSubmatch(strValue, -1)
	if len(res) > 0 {
		i64, _ := strconv.ParseInt(res[0][1], 10, 32)
		minutes += i64 % 12
		i64, _ = strconv.ParseInt(res[0][7], 10, 32)
		minutes += i64
		if res[0][5] == "PM" || res[0][5] == "pm" {
			minutes += 12
		}
		minutes *= 60
		i64, _ = strconv.ParseInt(res[0][4], 10, 32)
		minutes += i64
		i64, _ = strconv.ParseInt(res[0][9], 10, 32)
		minutes += i64
		return int(minutes), nil
	}
	return 0, errors.New("Cannot convert string to time")
}

func vuStringToDate(strValue string) (int, int, int, error) {
	const template = "2006-01-02"
	if t, e := time.Parse(template, strValue); e == nil {
		return t.Year(), int(t.Month()), t.Day(), e
	}
	return 0, 0, 0, errors.New("Cannot convert string to date")
}
