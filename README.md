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

Watch a process using $PID and execute specified command for notification when finished.
It can also watch a remote command.

## Installation

### go get

    % go get github.com/Songmu/peep/cmd/peep
    % go get github.com/Songmu/peep/cmd/peep-notify

### homebrew

    % brew install Songmu/tap/peep

### using [ghg](https://github.com/Songmu/ghg)

    % ghg get Songmu/peep

Built binaries are available on gihub releases.
<https://github.com/Songmu/peep/releases>

## Synopsis

### Easy Usage

    % peep [-H user@host -p $PORT] $PID -- peep-notify slack

![](https://user-images.githubusercontent.com/177122/50564546-05b53680-0d69-11e9-8d58-748c7469c325.png)

### Custom Notification Script

    % peep [-H user@host -p $PORT] $PID -- /path/to/your-notification-script

## Commands

- peep
  - main program
- peep-notify
  - bundled command to notify easily which supports [slack](https://slack.com), [pushpullet](https://www.pushbullet.com/) and mac desktop notification.

## Result JSON

The custom notification script accepts a result JSON via STDIN that reports command result like following.

```json
{
  "user": "Songmu",
  "command": "perl -E say $$; sleep 10",
  "startAt": "2018-12-31T17:29:56+09:00",
  "endAt": "2018-12-31T17:30:07+09:00",
  "host": "localhost",
  "pid":20028
}
```

## See Also

- [horenso](https://github.com/Songmu/horenso)

## Author

[Songmu](https://github.com/Songmu)
