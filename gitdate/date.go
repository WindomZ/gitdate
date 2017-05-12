package gitdate

import (
	"errors"
	"github.com/WindomZ/go-commander"
	"regexp"
	"strings"
	"time"
)

const DATE_FORMAT string = time.RFC3339

var DateArgvAction = func(c commander.Context) error {
	date, ok := c.GetString("<date>")
	if !ok || len(date) == 0 {
		return errors.New("ERROR: <date> should not be empty.")
	}

	offset, err := formatDate(date)
	if err != nil {
		return err
	}

	if c.Contain("--day") || c.Contain("--month") || c.Contain("--year") {
		offset = offset.AddDate(c.MustInt("--year"), c.MustInt("--month"), c.MustInt("--day"))
	}
	if i := c.MustInt("--hour"); i != 0 {
		offset = offset.Add(time.Duration(i) * time.Hour)
	}
	if i := c.MustInt("--minute"); i != 0 {
		offset = offset.Add(time.Duration(i) * time.Minute)
	}

	if offset.Equal(time.Now()) {
		return nil
	}

	return execCommandGitDate(offset.Unix())
}

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
