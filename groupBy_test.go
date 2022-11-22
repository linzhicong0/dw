package dataweave

import (
	"reflect"
	"strconv"
	"testing"
)

func TestGroupByInt(t *testing.T) {

	f := func(item int, index int) string {
		return strconv.Itoa(item)
	}

	type args struct {
		items []int
		f     GroupByFunc[int]
	}
	tests := []struct {
		name string
		args args
		want map[string][]int
	}{
		// TODO: Add test cases.
		{name: "GroupByInt", args: struct {
			items []int
			f     GroupByFunc[int]
		}{items: []int{1, 2, 3, 4, 1}, f: f}, want: map[string][]int{"1": {1, 1}, "2": {2}, "3": {3}, "4": {4}}},

		{name: "GroupByInt", args: struct {
			items []int
			f     GroupByFunc[int]
		}{items: []int{1, 1, 1, 1, 1}, f: f}, want: map[string][]int{"1": {1, 1, 1, 1, 1}}},

		{name: "GroupByInt", args: struct {
			items []int
			f     GroupByFunc[int]
		}{items: []int{1, 2, 2, 4, 4}, f: f}, want: map[string][]int{"1": {1}, "2": {2, 2}, "4": {4, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.items, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupByBool(t *testing.T) {

	f := func(item bool, index int) string {
		return strconv.FormatBool(item)
	}

	type args struct {
		items []bool
		f     GroupByFunc[bool]
	}
	tests := []struct {
		name string
		args args
		want map[string][]bool
	}{
		// TODO: Add test cases.
		{name: "GroupByBool", args: struct {
			items []bool
			f     GroupByFunc[bool]
		}{items: []bool{true, true, true}, f: f}, want: map[string][]bool{"true": {true, true, true}}},

		{name: "GroupByBool", args: struct {
			items []bool
			f     GroupByFunc[bool]
		}{items: []bool{true, true, false}, f: f}, want: map[string][]bool{"true": {true, true}, "false": {false}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.items, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupByStruct(t *testing.T) {

	type people struct {
		id   string
		name string
	}

	f := func(item people, index int) string {
		return item.name
	}

	peoples1 := []people{
		{id: "1", name: "P"},
		{id: "2", name: "P"},
		{id: "3", name: "P1"},
		{id: "4", name: "P2"},
		{id: "5", name: "P3"},
		{id: "6", name: "P4"},
	}
	wantPeoples1 := map[string][]people{
		"P": {
			{id: "1", name: "P"},
			{id: "2", name: "P"},
		},
		"P1": {
			{id: "3", name: "P1"},
		},
		"P2": {
			{id: "4", name: "P2"},
		},
		"P3": {
			{id: "5", name: "P3"},
		},
		"P4": {
			{id: "6", name: "P4"},
		},
	}

	peoples2 := []people{
		{id: "1", name: "P"},
		{id: "1", name: "P"},
		{id: "1", name: "P"},
		{id: "1", name: "P"},
		{id: "1", name: "P"},
		{id: "1", name: "P"},
	}

	wantPeoples2 := map[string][]people{
		"P": {
			{id: "1", name: "P"},
			{id: "1", name: "P"},
			{id: "1", name: "P"},
			{id: "1", name: "P"},
			{id: "1", name: "P"},
			{id: "1", name: "P"},
		},
	}

	type args struct {
		items []people
		f     GroupByFunc[people]
	}
	tests := []struct {
		name string
		args args
		want map[string][]people
	}{
		// TODO: Add test cases.
		{name: "GroupByStruct", args: struct {
			items []people
			f     GroupByFunc[people]
		}{items: peoples1, f: f}, want: wantPeoples1},
		{name: "GroupByStruct", args: struct {
			items []people
			f     GroupByFunc[people]
		}{items: peoples2, f: f}, want: wantPeoples2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.items, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
