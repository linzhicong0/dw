package dataweave

import (
	"reflect"
	"testing"
)

func TestDistinctByInt(t *testing.T) {

	var f DistinctFunc[int]

	f = func(item int, index int) any {
		return item
	}
	type args struct {
		items []int
		f     DistinctFunc[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Distinct int", args: struct {
			items []int
			f     DistinctFunc[int]
		}{items: []int{1, 1, 3, 4}, f: f}, want: []int{1, 3, 4}},
		{name: "Distinct int", args: struct {
			items []int
			f     DistinctFunc[int]
		}{items: []int{1, 1, 1, 1}, f: f}, want: []int{1}},
		{name: "Distinct int", args: struct {
			items []int
			f     DistinctFunc[int]
		}{items: []int{1, 2, 3, 4}, f: f}, want: []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctBy(tt.args.items, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctByString(t *testing.T) {

	var f DistinctFunc[string]

	f = func(item string, index int) any {
		return item
	}
	type args struct {
		items []string
		f     DistinctFunc[string]
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Distinct string", args: struct {
			items []string
			f     DistinctFunc[string]
		}{items: []string{"1", "1", "3", "4"}, f: f}, want: []string{"1", "3", "4"}},
		{name: "Distinct string", args: struct {
			items []string
			f     DistinctFunc[string]
		}{items: []string{"1", "1", "1", "1"}, f: f}, want: []string{"1"}},
		{name: "Distinct string", args: struct {
			items []string
			f     DistinctFunc[string]
		}{items: []string{"1", "2", "3", "4"}, f: f}, want: []string{"1", "2", "3", "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctBy(tt.args.items, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctByStruct(t *testing.T) {

	type people struct {
		id   string
		name string
	}

	var f DistinctFunc[people]

	f = func(item people, index int) any {
		return item.id
	}

	peoples1 := []people{
		{id: "1", name: "1"},
		{id: "2", name: "2"},
		{id: "3", name: "3"},
		{id: "4", name: "4"},
		{id: "5", name: "5"},
	}
	peoplesWant1 := []people{
		{id: "1", name: "1"},
		{id: "2", name: "2"},
		{id: "3", name: "3"},
		{id: "4", name: "4"},
		{id: "5", name: "5"},
	}

	peoples2 := []people{
		{id: "1", name: "1"},
		{id: "2", name: "2"},
		{id: "2", name: "3"},
		{id: "4", name: "4"},
		{id: "5", name: "5"},
	}
	peoplesWant2 := []people{
		{id: "1", name: "1"},
		{id: "2", name: "2"},
		{id: "4", name: "4"},
		{id: "5", name: "5"},
	}

	peoples3 := []people{
		{id: "1", name: "1"},
		{id: "1", name: "2"},
		{id: "1", name: "3"},
		{id: "1", name: "4"},
		{id: "1", name: "5"},
	}
	peoplesWant3 := []people{
		{id: "1", name: "1"},
	}

	type args struct {
		items []people
		f     DistinctFunc[people]
	}
	tests := []struct {
		name string
		args args
		want []people
	}{
		// TODO: Add test cases.
		{name: "Distinct people", args: struct {
			items []people
			f     DistinctFunc[people]
		}{items: peoples1, f: f}, want: peoplesWant1},

		{name: "Distinct people", args: struct {
			items []people
			f     DistinctFunc[people]
		}{items: peoples2, f: f}, want: peoplesWant2},

		{name: "Distinct people", args: struct {
			items []people
			f     DistinctFunc[people]
		}{items: peoples3, f: f}, want: peoplesWant3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctBy(tt.args.items, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
			}
		})
	}

}
