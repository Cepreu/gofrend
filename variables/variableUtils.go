package variables

import (
	"errors"
	"fmt"
	"strconv"
)

const timeRe = "^((0?[1-9])|(1[0-2])):([0-5][0-9]) ?(AM|PM|am|pm)|(([01][0-9])|(2[0-3])):([0-5][0-9])$"
func vuStringToMinutes(strValue string) (int64, error) {
	var minutes int64
	r, _ := regexp.Compile(timeRe)
	res := r.FindAllStringSubmatch(strValue, -1)
	if len(res) > 0 {
		i64, _ := strconv.ParseInt(res[0][1], 10, 64)
		minutes += i64
		i64, _ = strconv.ParseInt(res[0][7], 10, 64)
		minutes += i64
		if res[0][5] == "PM" || res[0][5] == "pm" {
			minutes += 12
		}
		minutes *= 60
		i64, _ = strconv.ParseInt(res[0][4], 10, 64)
		minutes += i64
		i64, _ = strconv.ParseInt(res[0][9], 10, 64)
		minutes += i64
		return minutes, nil
	}
	return minutes, errors.New("Cannot convert string to time")
}