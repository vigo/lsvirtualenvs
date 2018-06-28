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
$ go get github.com/vigo/ls_virtual_env

```

## Usage

@wip

---

## Change Log

**2018-06-28**

* Initial commit

---

## Contributer(s)

* [Uğur "vigo" Özyılmazel](https://github.com/vigo) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/vigo/ls_virtual_env/fork)
1. Create your `branch` (`git checkout -b my-features`)
1. `commit` yours (`git commit -am 'added killer options'`)
1. `push` your `branch` (`git push origin my-features`)
1. Than create a new **Pull Request**!

---

## License

This project is licensed under MIT

---

