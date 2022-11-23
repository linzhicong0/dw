package dataweave

import "reflect"

type UpdateFunc[T any] func(key string, value T, index int) T

func Update[T any](obj interface{}, f UpdateFunc[T]) interface{} {

	typ := reflect.TypeOf(obj)

	if typ.Kind() != reflect.Struct || obj == nil {
		return nil
	}

	valueOfStruct := reflect.ValueOf(obj)

	newValueOfStruct := reflect.New(typ).Elem()
	for i := 0; i < typ.NumField(); i++ {

		field := typ.Field(i)
		valueOfField := valueOfStruct.Field(i)
		newValue := f(field.Name, valueOfField.Interface().(T), i)
		newValueOfStruct.Field(i).Set(reflect.ValueOf(newValue))

	}

	return newValueOfStruct.Interface()
}
