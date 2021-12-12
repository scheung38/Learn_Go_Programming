package main

import (
	"reflect"
	"testing"
)

func Test_swap(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swap() = %v, want %v", got, tt.want)
			}
		})
	}
}
