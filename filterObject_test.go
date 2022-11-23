package dataweave

import (
	"reflect"
	"testing"
)

func TestFilterObject(t *testing.T) {

	type people struct {
		Id   string
		Name string
	}

	f1 := func(value string, key string, index int) bool {
		return key == "Name"
	}
	f2 := func(value string, key string, index int) bool {
		return value == "test2"
	}

	f3 := func(value string, key string, index int) bool {
		return true
	}

	type args struct {
		obj people
		f   FilterObjectFunc[string]
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{name: "TestFilterObject_Key", args: struct {
			obj people
			f   FilterObjectFunc[string]
		}{obj: people{Id: "1", Name: "test"}, f: f1}, want: struct {
			Name string
		}{Name: "test"}},
		{name: "TestFilterObject_Value", args: struct {
			obj people
			f   FilterObjectFunc[string]
		}{obj: people{Id: "1", Name: "test2"}, f: f2}, want: struct {
			Name string
		}{Name: "test2"}},
		{name: "TestFilterObject_f_Always_Return_True", args: struct {
			obj people
			f   FilterObjectFunc[string]
		}{obj: people{Id: "1", Name: "test2"}, f: f3}, want: struct {
			Id   string
			Name string
		}{Id: "1", Name: "test2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterObject(tt.args.obj, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
