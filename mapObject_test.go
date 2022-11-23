package dataweave

import (
	"reflect"
	"strings"
	"testing"
)

func TestMapObject(t *testing.T) {

	type people struct {
		Id   string
		Name string
	}

	f1 := func(value string, key string, index int) (newValue string, newKey string) {
		return strings.ToUpper(value), strings.ToUpper(key)
	}
	type args struct {
		obj interface{}
		f   MapObjectFunc[string]
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{name: "TestMapObject_String", args: struct {
			obj interface{}
			f   MapObjectFunc[string]
		}{obj: people{
			Id:   "1",
			Name: "test",
		}, f: f1}, want: struct {
			ID   string
			NAME string
		}{ID: "1", NAME: "TEST"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapObject(tt.args.obj, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapObjectAnyType(t *testing.T) {

	type people struct {
		Id   string
		Name string
		Age  int
	}

	f1 := func(value any, key string, index int) (newValue any, newKey string) {
		return value, strings.ToUpper(key)
	}
	type args struct {
		obj interface{}
		f   MapObjectFunc[any]
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{name: "TestMapObject_Any", args: struct {
			obj interface{}
			f   MapObjectFunc[any]
		}{obj: people{
			Id:   "1",
			Name: "test",
			Age:  25,
		}, f: f1}, want: struct {
			ID   string
			NAME string
			AGE  int
		}{ID: "1", NAME: "test", AGE: 25}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapObject(tt.args.obj, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
