package bitmap

import (
	"math"
	"testing"
)

func TestBit1Count(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 10",
			args{
				num: 10,
			},
			2,
		},
		{
			"test -10",
			args{
				num: -10,
			},
			62,
		},
		{
			"test -MinInt64",
			args{
				num: math.MinInt64,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bit1Count(tt.args.num); got != tt.want {
				t.Errorf("Bit1Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
