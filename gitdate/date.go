package gitdate

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

const DATE_FORMAT string = time.RFC3339

func formatDate(date string) (time.Time, error) {
	var _time = time.Now()
	if date = strings.TrimSpace(strings.ToLower(date)); date != "now" {
		date = regexp.MustCompile(`\d{1,2}(:\d{1,2}){1,2}`).FindString(date)
		date = regexp.MustCompile(`\d{1,2}`).ReplaceAllStringFunc(date,
			func(s string) string {
				if len(s) == 1 {
					return "0" + s
				}
				return s
			})
		offsetStr := _time.Format(DATE_FORMAT)
		if strings.Count(date, ":") == 1 {
			offsetStr = offsetStr[:11] + date + offsetStr[16:]
		} else {
			offsetStr = offsetStr[:11] + date + offsetStr[19:]
		}
		_time, _ = time.Parse(DATE_FORMAT, offsetStr)
		if _time.IsZero() {
			return _time, errors.New("ERROR: <date> formats: '15:04' or '15:04:05' or 'now'.")
		}
	}
	return _time, nil
}
