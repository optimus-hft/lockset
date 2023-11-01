# Lockset
![pipeline](https://github.com/optimus-hft/lockset/actions/workflows/go-ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/optimus-hft/lockset/branch/main/graph/badge.svg)](#)
[![Go Report Card](https://goreportcard.com/badge/github.com/optimus-hft/lockset)](https://goreportcard.com/report/github.com/optimus-hft/lockset)
[![Go Reference](https://pkg.go.dev/badge/github.com/optimus-hft/lockset.svg)](https://pkg.go.dev/github.com/optimus-hft/lockset)

## GoLang Dynamic Mutexes
Lockset provides dynamic mutexes based on lock keys. Each key is locked and unlocked separately and does not affect other keys.
Instead of protecting everything with a giant mutex, Different parts of code can be protected by a tiny mutex in isolation to provide more throughput and concurrency.

## Getting Started
### Installation
```
go get github.com/optimus-hft/lockset
```

### Usage

```go
package main

import (
	"github.com/optimus-hft/lockset"
)

func main() {

}
```

## Contributing
Pull requests and bug reports are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is licensed under the MIT License.