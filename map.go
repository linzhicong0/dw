package dataweave

// A Mapper can convert itself to type R
type Mapper[R any] interface {
	To(index int) R
}

// A function Map type T to R
type MapFunc[T any, R any] func(item T, index int) R

func Map[T any, R any](items []T, f MapFunc[T, R]) []R {

	var result []R
	for i, item := range items {
		result = append(result, f(item, i))
	}
	return result
}

// using the mapper to convert from type T to R
func MapWithMapper[T Mapper[R], R any](items []T, mapper Mapper[R]) []R {

	var result []R
	for i, item := range items {
		result = append(result, item.To(i))
	}
	return result
}

// TODO: Try build the go routine one
