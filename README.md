# ffp - Frivolous Functional Programming

**ffp** is a Go library that adds functional types and functions to the Go language. It is designed to make functional programming in Go more fun. I called it Frivolous because I like idiomatic go and don't thing functional concepts should be used most of the time. I made this library primarily for the fun of it.

## Installation

To install **ffp**, use `go get`:

```sh
go get github.com/Solidsilver/ffp
```

## Usage

Here's an example of how to use **ffp**:

```go
package main

import (
    "fmt"
    "github.com/Solidsilver/ffp"
)

func main() {
    // Create a new list.
    xs := []int {1, 2, 3, 4, 5}

    xsm := ffp.Map(xs, func(val int) int {
        return val * 2
    })

    // Print the result.
    fmt.Println(xsm)
    // []int{2, 4, 6, 8, 10}

    // Filter over the list.
    ys := xs.Filter(func(x int) bool {
        return x > 3
    })

    // Print the result.
    fmt.Println(ys)
    // []int{4, 5}
}
```

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue on GitHub. If you want to contribute code, please fork the repository and submit a pull request.

## License

This library is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
