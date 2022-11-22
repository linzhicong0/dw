package dataweave

type commonTypeConstraints interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string | uint | uint8 | uint16 | uint32 | uint64
}

type DistinctFunc[T any] func(item T, index int) any

func DistinctBy[T any](items []T, f DistinctFunc[T]) []T {
	if len(items) <= 0 {
		return nil
	}

	set := make(map[any]byte)
	var result []T

	for index, item := range items {

		key := f(item, index)
		if _, ok := set[key]; !ok {
			set[key] = 0
			result = append(result, item)
		}
	}
	return result
}
