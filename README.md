# xerus
[![Go Report Card](https://goreportcard.com/badge/github.com/zhaque44/xerus)](https://goreportcard.com/report/github.com/zhaque44/xerus)
[![GoDoc](https://godoc.org/github.com/zhaque44/xerus?status.svg)](https://godoc.org/github.com/zhaque44/xerus)
[![Build Status](https://travis-ci.org/zhaque44/xerus.svg?branch=master)](https://travis-ci.org/zhaque44/xerus)
[![codecov](https://codecov.io/gh/zhaque44/xerus/branch/master/graph/badge.svg)](https://codecov.io/gh/zhaque44/xerus)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub release](https://img.shields.io/github/release/zhaque44/xerus.svg)](
![Image](images/logo/xerus.png "Xerus")
## Getting started
First, make sure that you have Go installed on your computer: [golang](https://go.dev/).

```bash
$ brew install go
```
### Install this library
```
go get github.com/zhaque44/xerus
``` 
### Verify that the project is installed:
```
$ go list -m github.com/zhaque44/xerus
```
### Update your go.mod file
```
$ go mod tidy
```
### Usage
```go
import "github.com/zhaque44/xerus"

x := []float64{1, 2, 3, 4, 5}

percentile _ := xerus.Percentile(x, 0.5)

fmt.Println(percentile)
```