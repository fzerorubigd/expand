# expand

[![Build Status](https://travis-ci.org/fzerorubigd/expand.svg)](https://travis-ci.org/fzerorubigd/expand)
[![Coverage Status](https://coveralls.io/repos/fzerorubigd/expand/badge.svg?branch=master&service=github)](https://coveralls.io/github/fzerorubigd/expand?branch=master)
[![GoDoc](https://godoc.org/github.com/fzerorubigd/expand?status.svg)](https://godoc.org/github.com/fzerorubigd/expand)

--
    import "github.com/fzerorubigd/expand"

Package expand is a simple utility for expanding some variables in path

its base on https://github.com/mitchellh/go-homedir , with some more functionality
also there is a compiler build tag for getting home dir, not checking the
GOOS variable.

*WARNING* : I TEST IT ONLY ON Linux!

## Usage

#### func  HomeDir

```go
func HomeDir() (string, error)
```
HomeDir return the home directory of he current user without need to use CGO so
the cross compiling is not a pain just for getting a home directory.

#### func  Path

```go
func Path(in string) (string, error)
```
Path is a function to expand path with variables into real address. supported
variables are $HOME, $PWD and ~ if and only if its at the begin of string.

#### func  Pwd

```go
func Pwd() (string, error)
```
Pwd return the path to the current executable directory
