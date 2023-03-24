package main

import (
	"errors"
	"reflect"
)

var ErrWrongType = errors.New("wrong type")

// Reverse Safe way to reverse an array
func Reverse(arr interface{}) (err error) {
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
// NOTE: Panics if it is not an array or a *[]interface{}
func ReverseUnsafe(arr interface{}) {
	n := reflect.ValueOf(arr).Len()
	swap := reflect.Swapper(arr)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
