package dataweave

import "reflect"

type MapObjectFunc[T any] func(value T, key string, index int) (newValue T, newKey string)

func MapObject[T any](obj interface{}, f MapObjectFunc[T]) interface{} {

	typ := reflect.TypeOf(obj)

	if typ.Kind() != reflect.Struct || obj == nil {
		return nil
	}

	valueOfStruct := reflect.ValueOf(obj)
	var newStructFields []reflect.StructField
	var newStructFieldValues []reflect.Value

	for i := 0; i < typ.NumField(); i++ {

		field := typ.Field(i)
		valueOfField := valueOfStruct.Field(i)

		newValue, newKey := f(valueOfField.Interface().(T), field.Name, i)

		field.Name = newKey

		newStructFields = append(newStructFields, field)
		newStructFieldValues = append(newStructFieldValues, reflect.ValueOf(newValue))

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
