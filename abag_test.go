package abag

import (
	"reflect"
	"testing"
)

func Test_isAestheticallyPleasing(t *testing.T) {
	tc := []struct {
		name   string
		trees  []int
		expect bool
	}{
		{
			"valid row with 4 trees",
			[]int{4, 5, 3, 7},
			true,
		}, {
			"invalid row with 5 trees at random",
			[]int{3, 4, 5, 3, 7},
			false,
		}, {
			"invalid row with 5 trees progressing linearly",
			[]int{1, 2, 3, 4, 5},
			false,
		}, {
			"invalid row with 4 trees not aesthetically pleasing",
			[]int{4, 7, 2, 1},
			false,
		},
	}

	for _, value := range tc {
		t.Run(value.name, func(t *testing.T) {
			result := isAestheticallyPleasing(value.trees)
			if result != value.expect {
				t.Errorf("expected %v,\n got %v", value.expect, result)
			}
		})
	}
}

func TestSolution(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			"row with 5 trees at random",
			[]int{3, 4, 5, 3, 7},
			3,
		}, {
			"row with 5 trees progressing linearly",
			[]int{1, 2, 3, 4, 5},
			-1,
		}, {
			"row with 4 trees aesthetically pleasing",
			[]int{1, 3, 1, 2},
			0,
		}, {
			"row with 4 trees not aesthetically pleasing",
			[]int{4, 7, 2, 1},
			2,
		}, {
			"row with 17 trees aesthetically pleasing",
			[]int{4, 7, 2, 4, 3, 100, 5, 700795, 23509, 640964096, 35053095, 39509358095, 6, 7, 2, 4, 2},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeIndex(t *testing.T) {
	type args struct {
		s []int
		i int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"remove element at zero index",
			args{[]int{4, 7, 2, 1}, 0},
			[]int{7, 2, 1},
		}, {
			"remove element at third index",
			args{[]int{4, 7, 2, 1}, 3},
			[]int{4, 7, 2},
		}, {
			"remove element at second index",
			args{[]int{4, 7, 2, 1}, 2},
			[]int{4, 7, 1},
		}, {
			"remove element at first index",
			args{[]int{4, 7, 2, 1}, 1},
			[]int{4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeIndex(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
