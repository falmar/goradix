# GoRadix [![Build Status](https://travis-ci.org/falmar/goradix.svg?branch=master)](https://travis-ci.org/falmar/goradix) [![Codecov](https://img.shields.io/codecov/c/github/falmar/goradix.svg)]()

Radix Tree implementation written in Golang.

Thread Safety is optional. You will need to pass a boolean value when creating the Radix Tree
- true: thread safe
- false: non thread safe

#### Radix Tree
> In computer science, a radix tree (also radix trie or compact prefix tree) is a data structure that represents a space-optimized trie in which each node that is the only child is merged with its parent. - [Wikipedia](https://en.wikipedia.org/wiki/Radix_tree)

![](https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Patricia_trie.svg/400px-Patricia_trie.svg.png)

## Content
 - [Usage](#usage)
 - [Benchmarks](#benchmarks)
 - [License](#license)

## Usage:

Download: `go get github.com/falmar/goradix`

Import `import "github.com/falmar/goradix"`

### Insert

```go
// string required, value (optional)
func (r *Radix) Insert(s string, value ...interface{}){...}
```
```go
package main

import "github.com/falmar/goradix"

func main() {
	radix := goradix.New(false)

  // Simple string insert
	radix.Insert("romanus")
	radix.Insert("romane")
	radix.Insert("romulus")
}
```

### LookUp

```go
// string required
func (r *Radix) LookUp(s string) (interface{}, error) {...} {...}
```
```go
package main

import (
	"fmt"

	"github.com/falmar/goradix"
)

func main() {
	radix := goradix.New(false)
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
```go
// string
func (r Radix) Remove(s string) (bool) {...}
```
```go
package main

import (
	"fmt"

	"github.com/falmar/goradix"
)

func main() {
	radix := goradix.New(false)

	radix.Insert("romanus")
	radix.Insert("romane")
	radix.Insert("romulus")

	succeed := radix.Remove("romulus")

	fmt.Println(succeed) // true
}
```

### AutoComplete
```go
// string, bool required
func (r Radix) AutoComplete(s string, wholeWord bool) ([]string, error) {...}
```
```go
package main

import (
	"fmt"

	"github.com/falmar/goradix"
)

func main() {
	radix := goradix.New(false)
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

##### Non Thread Safe

```text
Go 1.6

BenchmarkInsertNTS-2	 1000000	      1736 ns/op	     279 B/op	       8 allocs/op
BenchmarkLookUpNTS-2	 2000000	       628 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveNTS-2	 2000000	       562 ns/op	       0 B/op	       0 allocs/op
```

```text
Go 1.7

BenchmarkInsertNTS-2   	 2000000	      1056 ns/op	     279 B/op	       8 allocs/op
BenchmarkLookUpNTS-2   	 3000000	       435 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveNTS-2   	 3000000	       539 ns/op	       0 B/op	       0 allocs/op
```

##### Thread Safe

```text
Go 1.6

BenchmarkInsertTS-2	 1000000	      1619 ns/op	     279 B/op	       8 allocs/op
BenchmarkLookUpTS-2	 1000000	      1018 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveTS-2	 1000000	      1181 ns/op	       0 B/op	       0 allocs/op
```

```text
Go 1.7

BenchmarkInsertTS-2   	 1000000	      1018 ns/op	     279 B/op	       8 allocs/op
BenchmarkLookUpTS-2   	 3000000	       553 ns/op	       0 B/op	       0 allocs/op
BenchmarkRemoveTS-2   	 2000000	       651 ns/op	       0 B/op	       0 allocs/op
```

**Benchmark machine**

```text
Architecture:          x86_64
CPU op-mode(s):        32-bit, 64-bit
Byte Order:            Little Endian
CPU(s):                2
On-line CPU(s) list:   0,1
Thread(s) per core:    1
Core(s) per socket:    1
Socket(s):             2
NUMA node(s):          1
Vendor ID:             GenuineIntel
CPU family:            6
Model:                 62
Model name:            Intel(R) Xeon(R) CPU E5-2630L v2 @ 2.40GHz
Stepping:              4
CPU MHz:               2399.998
BogoMIPS:              4799.99
Virtualization:        VT-x
Hypervisor vendor:     KVM
Virtualization type:   full
L1d cache:             32K
L1i cache:             32K
L2 cache:              256K
L3 cache:              15360K
NUMA node0 CPU(s):     0,1
```

## Contribution

Any contribution is more than welcome.

## License

MIT.
