package dataweave

import "reflect"

type FilterObjectFunc[T any] func(value T, key string, index int) bool

func FilterObject[T any](obj interface{}, f FilterObjectFunc[T]) interface{} {

	typ := reflect.TypeOf(obj)

	if typ.Kind() != reflect.Struct {
		return nil
	}

	valueOfStruct := reflect.ValueOf(obj)
	var newStructFields []reflect.StructField
	var newStructFieldValues []reflect.Value

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		valueOfField := valueOfStruct.Field(i)

		if f(valueOfField.Interface().(T), field.Name, i) {
			// The struct field and the field value will be in the same order
			newStructFields = append(newStructFields, field)
			newStructFieldValues = append(newStructFieldValues, valueOfField)
		}
	}

	if len(newStructFields) <= 0 {
		return nil
	}

	newTyp := reflect.StructOf(newStructFields)
	newValue := reflect.New(newTyp).Elem()

	for i := 0; i < len(newStructFields); i++ {
		newValue.Field(i).Set(newStructFieldValues[i])
	}

	return newValue.Interface()

}
