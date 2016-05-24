# GoRadix [![Build Status](https://travis-ci.org/falmar/goradix.svg?branch=master)](https://travis-ci.org/falmar/goradix)

Radix Tree implementation written in Golang. **Still under development**


#### Radix Tree
> In computer science, a radix tree (also radix trie or compact prefix tree) is a data structure that represents a space-optimized trie in which each node that is the only child is merged with its parent. - [Wikipedia](https://en.wikipedia.org/wiki/Radix_tree)

![](https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Patricia_trie.svg/400px-Patricia_trie.svg.png)

## Content
 - [Todo](#todo)
 - [Usage](#usage)
 - [Benchmarks](#benchmarks)
 - [License](#license)

## Todo


| Code Task | Implementation | Test | Benchmark
|---|:---:|:---:|:---:|
| Insert | x | x |  |
| LookUp | x | x |  |
| Remove |  |  |  |
| AutoComplete | x | x |  |

- Usage Examples:
    * [x] Set
    * [x] Get
    * [x] Insert
    * [x] LookUp
    * [ ] Remove
    * [ ] AutoComplete


## Usage:

Download: `go get github.com/falmar/goradix`

Import `import "github.com/falmar/goradix"`

### Set
```go
// Set a value to the Radix Tree node
func (r *Radix) Set(v interface{}) {...}
```
```go
package main

import "github.com/falmar/goradix"

func main() {
  radix := goradix.New()
  radix.Set(100)
}
```
### Get
```go
// Get a value from Radix Tree node
func (r *Radix) Get() interface{} {...}
```
```go
package main

import "github.com/falmar/goradix"

func main() {
  radix := goradix.New()
  radix.Set(100)
  fmt.Println(radix.Get()) // output: 100

  radix.Set("something")
  fmt.Println(radix.Get()) // output: something
}
```

### Insert | InsertBytes

```go
// s string required, value (optional)
func (r *Radix) Insert(s string, value ...interface{}){...}
// or
// bs slice of bytes required, value (optional)
func (r *Radix) InsertBytes(bs []byte, val ...interface{}) bool {...}
```
```go
package main

import "github.com/falmar/goradix"

func main() {
	radix := goradix.New()

  // Simple string insert
	radix.Insert("romanus")
	radix.Insert("romane")
	radix.Insert("romulus")

  // You can also insert slice of bytes
  radix.InsertBytes([]byte("rubens"))
}
```

### LookUp | LookUpBytes

```go
// bs string required
func (r *Radix) LookUp(s string) (interface{}, error) {...}
// or
// bs slice of bytes required
func (r *Radix) LookUpBytes(bs []byte) (interface{}, error) {...}
```
```go
package main

import (
	"fmt"

	"github.com/falmar/goradix"
)

func main() {
	radix := goradix.New()

	radix.Insert("romanus", 1)
	radix.Insert("romane", 100)
	radix.Insert("romulus", 1000)

	value, err := radix.LookUp("romane")

	if err != nil {
    // No Match Found
		fmt.Println(err)
	} else {
    // Found node, Value: 100
		fmt.Println("Found node, Value: ", value)
	}
}
```

### Remove
Under development

### AutoComplete
Under development

## Benchmarks
Under development

## License

Copyright 2016 David Lavieri. All rights reserved.
Use of this source code is governed by a MIT License that can be found in the LICENSE file.
