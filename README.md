# goLib
Some useful libraries for GoLang

## Install
~~~
go get github.com/kilfu0701/goLib
~~~

## Use goLib Modules

#### Zip
~~~ go
package main

import "github.com/kilfu0701/goLib"

func main() {
    zf := goLib.Zip("path/to/myfile.zip")
    zf.Extract("path/to/dest")
}
~~~

