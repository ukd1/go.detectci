# Go Detect Ci

Package ``detectci`` allows you to detect if, and what kind of CI service your code is running on, allowing you to behave differently when needed.

## Installation

Use the `go` command:

	$ go get github.com/ukd1/go.detectci

## Requirements

detectci package is tested against Go >= 1.16.

## Example

```go
package main

import (
	"github.com/ukd1/go.detectci"
	"log"
)

func main() {
	// Detecting any CI
	if detectci.IsCI() {
		log.Println("Looks like we are running in CI")
	} else {
		log.Println("Looks like we are not running in CI")
	}

	// Getting which CI by name
	found, ci_name := detectci.WhichCI()
	if found {
		log.Printf("Looks like we are running in %v", ci_name)
	} else {
		log.Println("Looks like we are not running in CI")
	}
}
```

```bash
% CIRCLECI=true go run example.go 
2021/05/11 18:38:10 Looks like we are running in CI
2021/05/11 18:38:10 Looks like we are running in circle-ci
```