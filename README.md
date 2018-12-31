peep
=======

[![Build Status](https://travis-ci.org/Songmu/peep.png?branch=master)][travis]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/Songmu/peep?status.svg)][godoc]

[travis]: https://travis-ci.org/Songmu/peep
[license]: https://github.com/Songmu/peep/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/Songmu/peep

peep a process

## Description

Watch a process and execute specified command for notification when finished.

## Installation

    % go get github.com/Songmu/peep/cmd/peep
    % go get github.com/Songmu/peep/cmd/peep-notify

Built binaries are available on gihub releases.
<https://github.com/Songmu/peep/releases>

## Synopsis

### Easy Usage

    % peep [-H user@host -p $PORT] $PID -- peep-notify slack

### Custom Script

    % peep [-H user@host -p $PORT] $PID -- /path/to/your-notification-script

## Commands

- peep
  - main program
- peep-notify
  - bundled command to notify easily which supports slack and pushpullet

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
