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
### Contributing
We welcome contributions from the community! If you would like to contribute to Xerus, please follow these steps:
1.) Fork the repository on GitHub.

2.) Clone your forked repository to your local machine.

3.) Create a new branch for your feature or bug fix.
```bash
$ git checkout -b my-feature-branch
```

4.) Make your changes and commit them with a descriptive commit message.
```bash
$ git commit -a -m "Added some feature for Xerus (description of feature) (issue number if applicable)"
```

5.) Push your branch to your forked repository on GitHub.
```bash
$ git push origin my-feature-branch
```

6.) Open a pull request against the main branch of this repository.

7.) Wait for your pull request to be reviewed and merged. We will review your code, test it, and merge it into the main branch if it meets our guidelines and quality standards.

Thank you for contributing to Xerus! If you have any questions or need any help with your contribution, please don't hesitate to reach out to us.