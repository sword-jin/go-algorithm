package backpack

import "testing"

func TestMaxValue(t *testing.T) {
	type args struct {
		total  int
		weight []int
		value  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test basic case",
			args{
				total:  4,
				weight: []int{1, 3, 4},
				value:  []int{15, 20, 30},
			},
			35,
		},
		{
			"test complex case",
			args{
				total:  5,
				weight: []int{1, 2, 2, 3, 4},
				value:  []int{15, 20, 30, 35, 60},
			},
			75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxValue2(tt.args.total, tt.args.weight, tt.args.value); got != tt.want {
				t.Errorf("MaxValue() = %v, want %v", got, tt.want)
			}
			if got := MaxValue(tt.args.total, tt.args.weight, tt.args.value); got != tt.want {
				t.Errorf("MaxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
