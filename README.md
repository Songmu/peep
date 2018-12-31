peep
=======

[![Build Status](https://travis-ci.org/Songmu/peep.png?branch=master)][travis]
[![Coverage Status](https://coveralls.io/repos/Songmu/peep/badge.png?branch=master)][coveralls]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/peep?status.svg)][godoc]

[travis]: https://travis-ci.org/Songmu/peep
[coveralls]: https://coveralls.io/r/Songmu/peep?branch=master
[license]: https://github.com/Songmu/peep/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/Songmu/peep

peep a process

## Description

Watch a process and execute specified command for notification when finished.

## Installation

    % go get github.com/Songmu/peep/cmd/peep
    % go get github.com/Songmu/peep/cmd/peepnotify

Built binaries are available on gihub releases.
<https://github.com/Songmu/peep/releases>

## Synopsis

    % peep [-H user@host -p $PORT] $PID -- peepnotify slack
    % peep [-H user@host -p $PORT] $PID -- /path/to/your-notification-script

## Result JSON

The notification script accepts a result JSON via STDIN that reports command result like following.

```json
{
  "user": "Songmu",
  "command": "perl -E say $$; sleep 10",
  "started": "2018-12-31T17:29:56+09:00",
  "ended": "2018-12-31T17:30:07+09:00",
  "host": "localhost"
}
```

## Author

[Songmu](https://github.com/Songmu)
