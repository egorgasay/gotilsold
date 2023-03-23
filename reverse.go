package main

import (
	"errors"
	"reflect"
)

var ErrWrongType = errors.New("wrong type")

// Reverse Safe way to reverse an array
func Reverse(arr any) (err error) {
	defer func() {
		if recover() != nil {
			err = ErrWrongType
		}
	}()
	n := reflect.ValueOf(arr).Len()
	swap := reflect.Swapper(arr)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return nil
}

// ReverseUnsafe Safe way to reverse an array
// NOTE: Panics if it is not an array or an *[]any
func ReverseUnsafe(arr any) {
	n := reflect.ValueOf(arr).Len()
	swap := reflect.Swapper(arr)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
