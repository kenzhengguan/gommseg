gommseg
=======

mmseg in go

## Install

$ go get github.com/raquelken/gommseg

## Getting Started
```go
package main

import (
	"fmt"
	"github.com/raquelken/gommseg"
)

func main() {
	text := "希望找到一个能发挥你能力的地方"
	words := gommseg.Ana.Cut(text)
	for _, word := range words {
		fmt.Println(word)
	}
}
```
  
