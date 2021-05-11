![Version](https://img.shields.io/badge/version-0.1.0-yellow.svg)
![Go](https://img.shields.io/badge/golang-1.16.3-black.svg)
[![Documentation](https://godoc.org/github.com/vigo/lsvirtualenvs?status.svg)](https://pkg.go.dev/github.com/vigo/lsvirtualenvs)
[![Go Report Card](https://goreportcard.com/badge/github.com/vigo/lsvirtualenvs)](https://goreportcard.com/report/github.com/vigo/lsvirtualenvs)
[![Build Status](https://travis-ci.org/vigo/lsvirtualenvs.svg?branch=main)](https://travis-ci.org/vigo/lsvirtualenvs)
![Go Build Status](https://github.com/vigo/lsvirtualenvs/actions/workflows/go.yml/badge.svg)

# List Virtual Environments for `virtualenvwrapper`

If you use `virtualenvwrapper` you’ll love this :)

Due to `virtualenvwrapper`’s `lsvirtualenv`’s super slow speed and lack of
information, I made this simple cli-tool with `golang`.


## Requirements

I’m assuming that you are already using [virtualenvwrapper][virtualenvwrapper]
and you have `WORKON_HOME` environment variable is already exists in your
shell environment.

---

## Installation

```bash
$ go get github.com/vigo/lsvirtualenvs
```

This will build and install binary of `lsvirtualenvs` under `$GOPATH/bin` path.

### Build from source

Check your `go env GOPATH` then check sources;

```bash
$ ls "$(go env GOPATH)/src/"
cloud.google.com  github.com  go.opencensus.io  golang.org  google.golang.org

# if github.com does not exists, create the folder via
# $ mkdir "$(go env GOPATH)/src/github.com"

$ mkdir "$(go env GOPATH)/src/github.com/vigo" # need for run/build operations
$ cd "$(go env GOPATH)/src/github.com/vigo"
$ git clone git@github.com:vigo/lsvirtualenvs.git
$ cd lsvirtualenvs/
$ go build
$ ls "$(go env GOPATH)/bin" # you should see `lsvirtualenvs` binary
$ lsvirtualenvs -h
```

## Usage

```bash
$ lsvirtualenvs -h

usage: lsvirtualenvs [-flags]

lists existing virtualenvs which are created via "mkvirtualenv" command.

  flags:

  -c, -color          enable colored output
  -s, -simple         just list environment names, overrides -c, -i
  -i, -index          add index number to output
      -version        display version information (X.X.X)
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

$ lsvirtualenvs -c -i # colored output with index
```

Run tests via;

```bash
$ go test -v ./...
```

---

## Docker

Note that, `lsvirtualenvs` requires `WORKON_HOME` environment variable which is
not available inside of the docker container :) This is just an example / concept
of dockerized version of the application :)

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

## Rake Tasks

```bash
$ rake -T

rake default            # show avaliable tasks (default task)
rake release[revision]  # Release new version major,minor,patch, default: patch
rake test[verbose]      # run tests
```

---

## Change Log

**2021-05-09**

- Add github action for go build status
- Changed `master` branch to `main`
- Add missing information to `-h` help
- Add Rake tasks

**2021-01-06**

- Complete make-over from scratch, removed `sync.Map()`, used channels
- Fix information on README

**2018-07-05**

- Due to @fatih’s warning, removed `Lock()` and used `sync.Map()`
- Version 2.1.1

**2018-07-04**

- App refactored
- Unit tests are completed
- Version 2.0.1

**2018-07-02**

- Basic unit testing

**2018-07-01**

- Code refactor

**2018-06-29**

- First release
- Addition: `--version`

**2018-06-28**

- Initial commit

---

## Contributer(s)

* [Uğur "vigo" Özyılmazel](https://github.com/vigo) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/vigo/lsvirtualenvs/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'Add garlic and yogurt'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

---

## License

This project is licensed under MIT

---


[virtualenvwrapper]: https://virtualenvwrapper.readthedocs.io/en/latest/