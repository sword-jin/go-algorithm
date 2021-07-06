package kmp

import (
	"reflect"
	"testing"
)

func Test_buildNext(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test [1, 2, 3, 1, 2, 5]",
			args{
				"123125",
			},
			[]int{-1, 0, 0, 0, 1, 2},
		},
		{
			"test [0, 1, 3, 4, 4, 0, 3, 4, 5]",
			args{
				"013440345",
			},
			[]int{-1, 0, 0, 0, 0, 0, 1, 0, 0},
		},
		{
			"test [a-z]",
			args{
				"abcdefghijklmnopqrstuvwxyz",
			},
			[]int{-1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildNext(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMPSearch(t *testing.T) {
	type args struct {
		s      string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	"test abc a",
		//	args{
		//		s:      "abc",
		//		substr: "a",
		//	},
		//	0,
		//},
		{
			"test abc d",
			args{
				s:      "abc",
				substr: "d",
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KMPSearch(tt.args.s, tt.args.substr); got != tt.want {
				t.Errorf("KMPSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
