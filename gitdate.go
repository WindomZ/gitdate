package main

import (
	"github.com/WindomZ/gitdate/gitdate"
	"github.com/WindomZ/go-commander"
)

func main() {
	// gitdate
	commander.Program.
		Version("0.0.1")

	// gitdate [options] <date>
	commander.Program.
		Command("<date>").
		Action(gitdate.DateArgvAction).
		Option("-m --minute <minute>", "date offset +/-number of minute").
		Option("-H --hour <hour>", "date offset +/-number of hour").
		Option("-d --day <day>", "date offset +/-number of day").
		Option("-M --month <month>", "date offset +/-number of month").
		Option("-y --year <year>", "date offset +/-number of year")

	if _, err := commander.Program.Parse(); err != nil {
		panic(err)
	}
}
