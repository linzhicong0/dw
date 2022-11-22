package dataweave

import (
	"reflect"
	"testing"
)

func Test_split(t *testing.T) {
	type args struct {
		total int
		by    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "total equals to by", args: struct {
			total int
			by    int
		}{total: 10, by: 10}, want: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{name: "total greater than by, total=10, by=3", args: struct {
			total int
			by    int
		}{total: 10, by: 3}, want: []int{3, 3, 4}},
		{name: "total greater than by, total=11, by=3", args: struct {
			total int
			by    int
		}{total: 11, by: 3}, want: []int{3, 4, 4}},
		{name: "total greater than by, total=100, by=3", args: struct {
			total int
			by    int
		}{total: 100, by: 3}, want: []int{33, 33, 34}},

		{name: "total greater than by, total=101, by=3", args: struct {
			total int
			by    int
		}{total: 101, by: 3}, want: []int{33, 34, 34}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.total, tt.args.by); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}
