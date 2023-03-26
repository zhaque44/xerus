# xerus
![Image](images/logo/xerus.png "Xerus")
## Getting started
First, make sure that you have Go installed on your computer: [golang](https://go.dev/)

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