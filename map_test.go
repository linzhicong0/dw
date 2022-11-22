package dataweave

import (
	"reflect"
	"testing"
)

type input struct {
	First_Name string
	Last_Name  string
}
type output struct {
	FirstName string
	LastName  string
}

func TestMap(t *testing.T) {

	var f MapFunc[input, output]

	f = func(item input, index int) output {

		return output{
			FirstName: item.First_Name,
			LastName:  item.Last_Name,
		}

	}

	inputs := []input{
		{First_Name: "Jack1", Last_Name: "Lin1"},
		{First_Name: "Jack2", Last_Name: "Lin2"},
		{First_Name: "Jack3", Last_Name: "Lin3"},
		{First_Name: "Jack4", Last_Name: "Lin4"},
	}
	wants := []output{
		{FirstName: "Jack1", LastName: "Lin1"},
		{FirstName: "Jack2", LastName: "Lin2"},
		{FirstName: "Jack3", LastName: "Lin3"},
		{FirstName: "Jack4", LastName: "Lin4"},
	}

	type args struct {
		input []input
		f     MapFunc[input, output]
	}
	tests := []struct {
		name string
		args args
		want []output
	}{
		// TODO: Add test cases.
		{name: "input list 1", args: struct {
			input []input
			f     MapFunc[input, output]
		}{input: inputs, f: f}, want: wants},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map[input, output](tt.args.input, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap2(t *testing.T) {

}
