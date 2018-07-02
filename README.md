![Version](https://img.shields.io/badge/version-0.1.1-yellow.svg)

# List Virtual Environments for `virtualenvwrapper`

If you use `virtualenvwrapper` you’ll love this :)

Due `virtualenvwrapper`’s `lsvirtualenv` super slow speed and lack of information,
I made this simple command in `golang`. This is also my newbie work in Golang.
Thanks to [Ahmet Aygün](https://github.com/ahmet) for showing me `sync.WaitGroup` :)

## Requirements

I’m assuming that you are already using [virtualenvwrapper](https://virtualenvwrapper.readthedocs.io/en/latest/)
and you have `WORKON_HOME` variable exists in your command-line environment.

- `Go 1.10.3` or higher

---

## Installation

```bash
$ go get -u github.com/vigo/lsvirtualenvs

```

This will build and install binary of `lsvirtualenvs` under `$GOPATH/bin` path.

## Usage

```bash
$ lsvirtualenvs -h

Usage: lsvirtualenvs [options...]

List available virtualenvironments created by virtualenvwrapper.

Options:

  -h, --help      Display help! :)
  -c, --color     Enable color output
  -s, --simple    Just list environment names
  -i, --index     Add index number to output
      --version   Version information

Examples:

  lsvirtualenvs -h
  lsvirtualenvs -c
  lsvirtualenvs --color
  lsvirtualenvs -c -i
  lsvirtualenvs --color --index
  lsvirtualenvs -s
  lsvirtualenvs -simple

```

## Benchmark

Let’s try with `virtualenvwrapper`’s `lsvirtualenv`:

```bash
$ time lsvirtualenv 
biges-training-django-cbv
=========================


bilgi.edu.tr
============


demo-project
============


django-project-template
=======================


django2-project-template
========================


put_io_cmd
==========


py27-bilgi.edu.tr
=================


py2712-gsc.bilgi.edu.tr
=======================


py3-http-server
===============


py360-bilgi
===========


py360-pt
========


py360-velisiyim.com
===================


py363-personal-bsp
==================


py363-redis-pubsub
==================


py363-websockets
================


py365-bcp-backend
=================


py365-bcp-py-grpc
=================


py365-biges.com
===============


sil
===


sil2
====


training-django-101-py3-blog-app
================================


training-django-cbv-example
===========================



real     0m5.892s
user     0m4.451s
sys      0m1.256s

```

Let’s try with `lsvirtualenvs`:

```bash
$ time lsvirtualenvs
You have 22 virtualenvs available

[biges-training-django-cbv] ............ 3.6.3
[bilgi.edu.tr] ......................... 2.7.13
[demo-project] ......................... 3.6.3
[django-project-template] .............. 3.6.0
[django2-project-template] ............. 3.6.4
[put_io_cmd] ........................... 3.6.3
[py27-bilgi.edu.tr] .................... 2.7.13
[py2712-gsc.bilgi.edu.tr] .............. 2.7.12
[py3-http-server] ...................... 3.6.3
[py360-bilgi] .......................... 3.6.0
[py360-pt] ............................. 3.6.0
[py360-velisiyim.com] .................. 3.6.0
[py363-personal-bsp] ................... 3.6.3
[py363-redis-pubsub] ................... 3.6.3
[py363-websockets] ..................... 3.6.3
[py365-bcp-backend] .................... 3.6.5
[py365-bcp-py-grpc] .................... 3.6.5
[py365-biges.com] ...................... 3.6.5
[sil] .................................. 3.6.3
[sil2] ................................. 3.6.3
[training-django-101-py3-blog-app] ..... 3.6.0
[training-django-cbv-example] .......... 3.6.0

real    0m0.266s
user    0m0.091s
sys     0m0.110s

```

Blazing fast :)

Run tests via;

```bash
$ go test -v

```

---

## Change Log

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

