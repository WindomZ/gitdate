package gitdate

import (
	"errors"
	"github.com/WindomZ/go-commander"
	"strconv"
	"strings"
	"time"
)

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

func execCommandGitDate(unix int64) error {
	script, err := httpGet(script_url)
	if err != nil {
		return err
	}
	if _, _, err = commander.Exec.ExecPipeCommand("git", "stash"); err != nil {
		return err
	}
	defer func() {
		commander.Exec.ExecPipeCommand("git", "stash", "pop")
	}()
	_, err = commander.Exec.ExecStdStatementCommand(
		strings.Replace(script, "$1",
			strconv.FormatInt(unix, 10), -1),
	)
	return err
}
