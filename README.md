# lukasmalkmus/interactive
> Simple and easy interactive commandline applications. - by **[Lukas Malkmus](https://github.com/lukasmalkmus)**

[![Travis Status][travis_badge]][travis]
[![Coverage Status][coverage_badge]][coverage]
[![Go Report][report_badge]][report]
[![GoDoc][docs_badge]][docs]
[![Latest Release][release_badge]][release]
[![License][license_badge]][license]

---

## Table of Contents
1. [Introduction](#introduction)
2. [Usage](#usage)
3. [Contributing](#contributing)
4. [License](#license)

### Introduction
This package is build on top of the excellent `golang.org/x/crypto/ssh/terminal`
package and tries to simplify the creation of small and simple applications
which run in shell mode.
It isn't very powerful (yet) but has enough features for basic usage. For
example it shuts down gracefully on Ctrl^C and Ctrl^D.

#### Todo
  - [ ] Tests!

### Usage
#### Installation
Please use a dependency manager like [glide](http://glide.sh) to make sure you
use a tagged release.

Install using `go get`:
```bash
go get -u github.com/lukasmalkmus/interactive
```

#### Minimum setup
A more complete example can be found [here](examples/usage.go).

This small application echos the entered text:
```go
s := interactive.New(">")
s.Action = func(c *interactive.Context) error {
    text, _ := c.Scan()
    c.Println(text)
    return nil
}
s.Run()
```
Returning an error instead of nil will print the error. Calling `context.Close()`
will close the session with the specified exit code.

### Contributing
Feel free to submit PRs or to fill Issues. Every kind of help is appreciated.

### License
© Lukas Malkmus, 2017

Distributed under MIT License (`The MIT License`).

See [LICENSE](LICENSE) for more information.


[travis]: https://travis-ci.org/lukasmalkmus/interactive
[travis_badge]: https://travis-ci.org/lukasmalkmus/interactive.svg
[coverage]: https://coveralls.io/github/lukasmalkmus/interactive?branch=master
[coverage_badge]: https://coveralls.io/repos/github/lukasmalkmus/interactive/badge.svg?branch=master
[report]: https://goreportcard.com/report/github.com/lukasmalkmus/interactive
[report_badge]: https://goreportcard.com/badge/github.com/lukasmalkmus/interactive
[docs]: https://godoc.org/github.com/lukasmalkmus/interactive
[docs_badge]: https://godoc.org/github.com/lukasmalkmus/interactive?status.svg
[release]: https://github.com/lukasmalkmus/interactive/releases
[release_badge]: https://img.shields.io/github/release/lukasmalkmus/interactive.svg
[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/badge/license-MIT-blue.svg