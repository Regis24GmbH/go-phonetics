[![Build Status][travis-image]][travis-url]
[![Github Tag][githubtag-image]][githubtag-url]

[![Coverage Status][coveralls-image]][coveralls-url]
[![Maintainability][codeclimate-image]][codeclimate-url]

[![Go Report Card][goreport-image]][goreport-url]
[![GoDoc][godoc-image]][godoc-url]
[![License][license-image]][license-url]

***

# go-phonetics

> A package that contain an implementation of cologne phonetics (Kölner Phonetik)

Gonetic implements the Kölner Phonetik (Cologne Phonetic) algorithm in Go. It 
is a translation of the php implementation of 
[deezaster](https://github.com/deezaster/germanphonetic) to Go.

## Usage

To get the lastest tagged version of package, execute:

```
go get gopkg.in/Regis24GmbH/go-phonetics.v2
```

To import this package, add the following line to your code:

```
import "gopkg.in/Regis24GmbH/go-phonetics.v2"
```

This is a code example:

```
func main() {
  code := gophonetics.NewPhoneticCode("Müller-Lüdenscheidt")
  println(code) // prints "65752682"
}
``` 

***

[travis-image]: https://travis-ci.org/Regis24GmbH/go-phonetics.svg?branch=master
[travis-url]: https://travis-ci.org/Regis24GmbH/go-phonetics

[githubtag-image]: https://img.shields.io/github/tag/Regis24GmbH/go-phonetics.svg?style=flat
[githubtag-url]: https://github.com/Regis24GmbH/go-phonetics

[coveralls-image]: https://coveralls.io/repos/github/Regis24GmbH/go-phonetics/badge.svg?branch=master
[coveralls-url]: https://coveralls.io/github/Regis24GmbH/go-phonetics?branch=master

[codeclimate-image]: https://api.codeclimate.com/v1/badges/beaa980f05b75209c940/maintainability
[codeclimate-url]: https://codeclimate.com/github/Regis24GmbH/go-phonetics/maintainability

[goreport-image]: https://goreportcard.com/badge/github.com/Regis24GmbH/go-phonetics
[goreport-url]: https://goreportcard.com/report/github.com/Regis24GmbH/go-phonetics

[godoc-image]: https://godoc.org/github.com/Regis24GmbH/go-phonetics?status.svg
[godoc-url]: https://godoc.org/github.com/Regis24GmbH/go-phonetics

[license-image]: https://img.shields.io/github/license/Regis24GmbH/go-phonetics.svg?style=flat
[license-url]: https://github.com/Regis24GmbH/go-phonetics/blob/master/LICENSE
