package dataweave

type GroupByFunc[T any] func(item T, index int) string

func GroupBy[T any](items []T, f GroupByFunc[T]) map[string][]T {

	if len(items) <= 0 {
		return nil
	}
	result := make(map[string][]T)

	for index, item := range items {
		key := f(item, index)
		list, ok := result[key]
		if ok {
			list = append(list, item)
			result[key] = list
		} else {
			newList := []T{item}
			result[key] = newList
		}
	}
	return result
}
