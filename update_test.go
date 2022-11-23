package dataweave

import (
	"reflect"
	"strings"
	"testing"
)

func TestUpdate(t *testing.T) {

	type people struct {
		Id   string
		Name string
		Age  int
	}

	f := func(key string, value any, index int) any {

		switch key {
		case "Name":
			return strings.ToUpper(value.(string))
		case "Age":
			return value.(int) + 1
		default:
			return value
		}
	}

	type args struct {
		obj interface{}
		f   UpdateFunc[any]
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "TestUpdate", args: struct {
			obj interface{}
			f   UpdateFunc[any]
		}{obj: people{Id: "1", Name: "test", Age: 25}, f: f}, want: people{
			Id: "1", Name: "TEST", Age: 26,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Update(tt.args.obj, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
