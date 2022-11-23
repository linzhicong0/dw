package dataweave

import (
	"reflect"
	"testing"
)

func TestReduceInt(t *testing.T) {

	f := func(item int, acc int) int {
		return acc + item
	}

	type args struct {
		items []int
		f     ReduceFunc[int, int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "ReduceInt", args: struct {
			items []int
			f     ReduceFunc[int, int]
		}{items: []int{1, 2, 3}, f: f}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.items, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduceStruct(t *testing.T) {
	type env struct {
		id   string
		name string
	}

	f := func(item env, acc map[string]string) map[string]string {

		if _, ok := acc[item.name]; !ok {
			acc[item.name] = item.id
		}
		return acc
	}

	envs := []env{
		{id: "1", name: "dev"},
		{id: "2", name: "test"},
		{id: "3", name: "uat"},
		{id: "4", name: "prod"},
	}

	wants := map[string]string{
		"dev":  "1",
		"test": "2",
		"uat":  "3",
		"prod": "4",
	}

	type args struct {
		items []env
		f     ReduceFunc[env, map[string]string]
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{

		{name: "ReduceStruct", args: struct {
			items []env
			f     ReduceFunc[env, map[string]string]
		}{items: envs, f: f}, want: wants},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.items, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
