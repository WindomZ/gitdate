# gitdate

[![Build Status](https://travis-ci.org/WindomZ/gitdate.svg?branch=master)](https://travis-ci.org/WindomZ/gitdate)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/gitdate)](https://goreportcard.com/report/github.com/WindomZ/gitdate)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

> After `commit` and before `push`, the last commit date is just a variable.

![v1.1.0](https://img.shields.io/badge/version-v1.1.0-blue.svg)
![status](https://img.shields.io/badge/status-stable-green.svg)

## Installation

To get the package, execute:

```bash
go get -u github.com/WindomZ/gitdate
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
    <date>        two formats: '15:04' or '15:04:05' or 'now'
```

## Example
```bash
gitdate 8:22                       # 08:22 today
gitdate 13:45 -d -1                # 13:45 yesterday
gitdate 13:45:21 -d -1             # 13:45:21 yesterday
gitdate now                        # now, current time
gitdate now -H -3 -m 5             # now, subtract 3 hours, and plus 5 minutes
gitdate now -M -1 -d -2 -H 3 -m 5  # now, subtract 1 month 2 days, and plus 3 hours 5 minutes
```

## PowerBy

[go-commander](https://github.com/WindomZ/go-commander) - The solution for building command shell programs

## Related

[WindomZ/gitdate.js](https://github.com/WindomZ/gitdate.js) - Written in **Node.js**

## License

The [MIT License](https://github.com/WindomZ/gitdate/blob/dev/LICENSE)
