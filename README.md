# gitdate
[![Build Status](https://travis-ci.org/WindomZ/gitdate.svg?branch=master)](https://travis-ci.org/WindomZ/gitdate)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/gitdate)](https://goreportcard.com/report/github.com/WindomZ/gitdate)

Nothing to say...

![v1.0.3](https://img.shields.io/badge/version-v1.0.3-blue.svg)
![status](https://img.shields.io/badge/status-stable-green.svg)

## Installation

To get the package, execute:

```bash
go get gopkg.in/WindomZ/gitdate.v1
```

To import this package, add the following line to your code:

```go
import "gopkg.in/WindomZ/gitdate.v1"
```

## Usage
```bash
$ gitdate -h

  Usage:
    gitdate <date> [-m=<minutes>|--minute=<minutes>] [-H=<hours>|--hour=<hours>] [-d=<days>|--day=<days>] [-M=<months>|--month=<months>] [-y=<years>|--year=<years>]
    gitdate -h|--help
    gitdate -v|--version

  Options:
    -m --minute <minutes>
                  date offset +/-number of minutes
    -H --hour <hours>
                  date offset +/-number of hours
    -d --day <days>
                  date offset +/-number of days
    -M --month <months>
                  date offset +/-number of months
    -y --year <years>
                  date offset +/-number of years
    -h --help     output usage information
    -v --version  output the version number

  Argument:
    <date>        two formats: '15:04' or 'now'
```

## License

The [MIT License](https://github.com/WindomZ/gitclone/blob/master/LICENSE)
