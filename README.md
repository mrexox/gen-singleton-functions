# gen-singleton-functions

## Usage

Make methods `Bark` and `Validate` public functions of a package `dog`.

```go
package dog

import "fmt"

//go:generate gen-singleton-functions github.com/me/dog *Dog
type Dog struct {
    numTails int
    barkingSound string
}

func (d *Dog) Bark() string {
    return d.barkingSound
}

func (d *Dog) Validate() error {
    if d.numTails != 1 {
        return fmt.Errorf("unexpected number of tails: %d", d.numTails)
    }

    return nil
}
```

After running `go generate ./...` the resulting file will contain the following:

```go
// Code generated by generator, DO NOT EDIT.
//
// This file provides the wrappers for exported methods of a global var `globalDog`.
package dog

var globalDog *Dog

func SetGlobalDog(v *Dog) {
	globalDog = v
}

// Bark calls `globalDog.Bark`.
func Bark() string {
	if globalDog == nil {
		panic("globalDog instance is not set. Call SetGlobalDog(var) before calling Bark")
	}

	return globalDog.Bark()
}

// Validate calls `globalDog.Validate`.
func Validate() error {
	if globalDog == nil {
		panic("globalDog instance is not set. Call SetGlobalDog(var) before calling Validate")
	}

	return globalDog.Validate()
}
```
