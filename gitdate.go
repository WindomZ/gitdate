package main

import (
	"github.com/WindomZ/gitdate/gitdate"
	"github.com/WindomZ/go-commander"
)

func main() {
	// gitdate
	commander.Program.
		Version("v1.0.5")

	// gitdate [options] <date>
	commander.Program.
		Command("<date>").
		Action(gitdate.DateArgvAction).
		Option("-m --minute <minutes>", "date offset +/-number of minutes").
		Option("-H --hour <hours>", "date offset +/-number of hours").
		Option("-d --day <days>", "date offset +/-number of days").
		Option("-M --month <months>", "date offset +/-number of months").
		Option("-y --year <years>", "date offset +/-number of years")

	commander.Program.Annotation(
		"Argument",
		[]string{commander.Format.Description("<date>", "two formats: '15:04' or 'now'")},
	)

	if _, err := commander.Program.Parse(); err != nil {
		panic(err)
	}
}
