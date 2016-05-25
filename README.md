# GoRadix [![Build Status](https://travis-ci.org/falmar/goradix.svg?branch=master)](https://travis-ci.org/falmar/goradix)

Radix Tree implementation written in Golang. **Still under development**


#### Radix Tree
> In computer science, a radix tree (also radix trie or compact prefix tree) is a data structure that represents a space-optimized trie in which each node that is the only child is merged with its parent. - [Wikipedia](https://en.wikipedia.org/wiki/Radix_tree)

![](https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Patricia_trie.svg/400px-Patricia_trie.svg.png)

## Content
 - [Usage](#usage)
 - [Benchmarks](#benchmarks)
 - [License](#license)

## Todo


| Code Task | Implementation | Test | Benchmark
|---|:---:|:---:|:---:|
| Insert | x | x |  |
| LookUp | x | x |  |
| AutoComplete | x | x |  |
| Match |  |  |  |
| Remove |  |  |  | |

- Usage Examples:
    * [x] Insert
    * [x] LookUp
    * [x] AutoComplete
    * [ ] Match
    * [ ] Remove


## Usage:

Download: `go get github.com/falmar/goradix`

Import `import "github.com/falmar/goradix"`

### Insert | InsertBytes

```go
// string required, value (optional)
func (r *Radix) Insert(s string, value ...interface{}){...}
// or
// slice of bytes required, value (optional)
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
// string required
func (r *Radix) LookUp(s string) (interface{}, error) {...}
// or
// slice of bytes required
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

	if err != nil { // No Match Found
		return
	}

	// Output: Found node, Value: 100
	fmt.Println("Found node, Value: ", value)
}
```

### Remove
Under development

### AutoComplete
```go
// string required, bool required
func (r Radix) AutoComplete(s string, wholeWord bool) ([]string, error) {...}
// or
// slice of bytes required,  bool required
func (r Radix) AutoComplete(bs []byte, wholeWord bool) ([][]byte, error) {...}
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

	// Return remaining text
	words, err := radix.AutoComplete("ro", false)

	if err != nil { // No Match Found
		return
	}

	// Output: AutoComplete 'rom'; Words: [manus mane mulus]
	fmt.Printf("AutoComplete: '%s'; Words: %v\n", "ro", words)

	// Return whole words
	words, _ = radix.AutoComplete("ro", true)

	// Output: AutoComplete 'rom'; Words: [romanus romane romulus]
	fmt.Printf("AutoComplete: '%s'; Words: %v\n", "ro", words)
}
```

## Benchmarks
Under development

## License

Copyright 2016 David Lavieri. All rights reserved.
Use of this source code is governed by a MIT License that can be found in the LICENSE file.
