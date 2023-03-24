![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/egorgasay/gotils)
![GitHub issues](https://img.shields.io/github/issues/egorgasay/gotils)
![GitHub](https://img.shields.io/github/license/egorgasay/gotils)
# gotils

## Collection of implementations of useful functions for the Go language

## Quick start
```
go get github.com/egorgasay/gotils
```

# Reverse (Unsafe)
```go
arr := []int{1, 2, 3, 4, 5, 6}
gotils.ReverseUnsafe(arr)
fmt.Println(arr) // [6 5 4 3 2 1]
```
## Note: Panics if it is not an array or 
## To avoid panic use gotils.Reverse(obj any)

# Reverse (Safe)
```go
arr := []int{1, 2, 3, 4, 5, 6}
err := gotils.Reverse(arr)
if err != nil { 
    fmt.Printf("Oops, an error was occured: %v\n", err)
} else {
    fmt.Println(arr) 
}

// Output: [6 5 4 3 2 1]

var parr *[]int
err := gotils.Reverse(arr)
if err != nil { 
    fmt.Printf("Oops, an error was occured: %v\n", err)
} else {
    fmt.Println(arr) // [6 5 4 3 2 1]
}

// Output: Oops, an error was occured: wrong type
```
