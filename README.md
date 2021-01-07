![Version](https://img.shields.io/badge/version-2.1.2-yellow.svg)
![Go](https://img.shields.io/badge/golang-1.15.6-black.svg)
[![Documentation](https://godoc.org/github.com/vigo/lsvirtualenvs?status.svg)](https://pkg.go.dev/github.com/vigo/lsvirtualenvs)
[![Go Report Card](https://goreportcard.com/badge/github.com/vigo/lsvirtualenvs)](https://goreportcard.com/report/github.com/vigo/lsvirtualenvs)
[![Build Status](https://travis-ci.org/vigo/lsvirtualenvs.svg?branch=master)](https://travis-ci.org/vigo/lsvirtualenvs)

# List Virtual Environments for `virtualenvwrapper`

If you use `virtualenvwrapper` you’ll love this :)

Due `virtualenvwrapper`’s `lsvirtualenv` super slow speed and lack of information,
I made this simple command in `golang`.


## Requirements

I’m assuming that you are already using [virtualenvwrapper](https://virtualenvwrapper.readthedocs.io/en/latest/)
and you have `WORKON_HOME` variable exists in your command-line environment.

- `Go 1.15.6` or higher

---

## Installation

```bash
$ go get -u github.com/vigo/lsvirtualenvs
```

This will build and install binary of `lsvirtualenvs` under `$GOPATH/bin` path.

## Usage

```bash
$ lsvirtualenvs -h

usage: lsvirtualenvs [-flags]

  flags:

  -c, -color          enable colored output
  -s, -simple         just list environment names, overrides -c, -i
  -i, -index          add index number to output
      -version        display version information (3.0.0)
```

Usage examples:

```bash
$ lsvirtualenvs -h
$ lsvirtualenvs -c
$ lsvirtualenvs -color
$ lsvirtualenvs -c -i
$ lsvirtualenvs -color -index
$ lsvirtualenvs -s
$ lsvirtualenvs -simple
```

Example output:

```bash
$ lsvirtualenvs
you have 2 environments available

textmate................... 3.8.0
trash...................... 3.8.0

$ lsvirtualenvs -i
you have 2 environments available

[0001] textmate................... 3.8.0
[0002] trash...................... 3.8.0
```

Run tests via;

```bash
$ go test -v ./...
```

## Docker

Note that, app checks for `WORKON_HOME` environment variable which is
not available inside of the docker container :) This is just an
example :)

Build:

```bash
$ docker build -t lsvirtualenvs .
```

Run:

```bash
$ docker run -i -t lsvirtualenvs lsvirtualenvs
WORKON_HOME environment variable doesn't exists in your environment

$ docker run -i -t lsvirtualenvs lsvirtualenvs -h
```

---

## Change Log

**2021-01-06**

* Complete make-over from scratch, removed `sync.Map()`, used channels

**2018-07-05**

* Due to @fatih’s warning, removed `Lock()` and used `sync.Map()`
* Version 2.1.1

**2018-07-04**

* App refactored
* Unit tests are completed
* Version 2.0.1

**2018-07-02**

* Basic unit testing

**2018-07-01**

* Code refactor

**2018-06-29**

* First release
* Addition: `--version`

**2018-06-28**

* Initial commit

---

## Contributer(s)

* [Uğur "vigo" Özyılmazel](https://github.com/vigo) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/vigo/lsvirtualenvs/fork)
1. Create your `branch` (`git checkout -b my-features`)
1. `commit` yours (`git commit -am 'added killer options'`)
1. `push` your `branch` (`git push origin my-features`)
1. Than create a new **Pull Request**!

---

## License

This project is licensed under MIT

---

