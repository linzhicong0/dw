package dataweave

type FilterFunc[T any] func(item T, index int) bool

func Filter[T any](items []T, f FilterFunc[T]) []T {

	var result []T

	for i, item := range items {

		if f(item, i) {
			result = append(result, item)
		}
	}
	return result
}
