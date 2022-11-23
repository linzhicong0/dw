package dataweave

import "reflect"

type ReduceFunc[T any, R any] func(item T, acc R) R

func Reduce[T any, R any](items []T, f ReduceFunc[T, R]) R {

	var acc R
	if len(items) <= 0 {
		return acc
	}

	accType := reflect.TypeOf(acc)
	if accType.Kind() == reflect.Map {
		acc = reflect.MakeMap(accType).Interface().(R)
	}

	for _, item := range items {
		acc = f(item, acc)
	}

	return acc
}
