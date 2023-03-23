package main

import (
	"reflect"
	"testing"
)

func Test_Reverse(t *testing.T) {
	type testType struct {
		id int
	}
	type testCase struct {
		name    string
		arg     any
		want    any
		wantErr bool
	}
	tests := []testCase{
		{
			name: "Reverse []int",
			arg:  []int{1, 2, 3, 4, 5, 6},
			want: []int{6, 5, 4, 3, 2, 1},
		},
		{
			name: "Reverse []testType",
			arg:  []testType{{id: 1}, {id: 2}, {id: 3}, {id: 4}, {id: 5}, {id: 6}},
			want: []testType{{id: 6}, {id: 5}, {id: 4}, {id: 3}, {id: 2}, {id: 1}},
		},
		{
			name:    "Reverse nil",
			arg:     nil,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Reverse int",
			arg:     1,
			want:    1,
			wantErr: true,
		},
		{
			name:    "Reverse *[]int",
			arg:     new([]int),
			want:    new([]int),
			wantErr: true,
		},
		{
			name: "Reverse empty []int",
			arg:  []int{},
			want: []int{},
		},
		{
			name:    "Reverse *[]int",
			arg:     &[]int{1, 2},
			want:    &[]int{1, 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Reverse(tt.arg)
			if err != nil && !tt.wantErr {
				t.Fail()
			}
			if !reflect.DeepEqual(tt.arg, tt.want) {
				t.Fail()
			}
		})
	}
}

func Test_ReverseUnsafe(t *testing.T) {
	type testType struct {
		id int
	}
	type testCase struct {
		name      string
		arg       any
		want      any
		wantPanic bool
	}
	tests := []testCase{
		{
			name: "Reverse []int",
			arg:  []int{1, 2, 3, 4, 5, 6},
			want: []int{6, 5, 4, 3, 2, 1},
		},
		{
			name: "Reverse []testType",
			arg:  []testType{{id: 1}, {id: 2}, {id: 3}, {id: 4}, {id: 5}, {id: 6}},
			want: []testType{{id: 6}, {id: 5}, {id: 4}, {id: 3}, {id: 2}, {id: 1}},
		},
		{
			name:      "Reverse nil",
			arg:       nil,
			want:      nil,
			wantPanic: true,
		},
		{
			name:      "Reverse int",
			arg:       1,
			want:      1,
			wantPanic: true,
		},
		{
			name:      "Reverse *[]int",
			arg:       new([]int),
			want:      new([]int),
			wantPanic: true,
		},
		{
			name: "Reverse empty []int",
			arg:  []int{},
			want: []int{},
		},
		{
			name:      "Reverse *[]int",
			arg:       &[]int{1, 2},
			want:      &[]int{1, 2},
			wantPanic: true,
		},
	}

	var resch = make(chan bool)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				defer func() {
					if recover() != nil {
						resch <- true
					}
				}()
				ReverseUnsafe(tt.arg)
				resch <- false
			}()

			if <-resch != tt.wantPanic {
				t.Fail()
			}
			if !reflect.DeepEqual(tt.arg, tt.want) {
				t.Fail()
			}
		})
	}
}
