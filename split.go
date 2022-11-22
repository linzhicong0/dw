package dataweave

func split(total, by int) []int {

	if total <= 0 || by <= 0 {
		return nil
	}

	var interval int

	var result []int

	if by == total {
		interval = 1
		for i := 0; i < by; i++ {
			result = append(result, interval)
		}
		return result
	}

	// total: 10, by: 3
	// 3, 3, 4
	// total: 11, by: 3
	// interval: 3, mod: 2
	// 3, 4, 4
	if by < total {
		interval = total / by
		mod := total % by
		for i := 0; i < by; i++ {
			if i >= by-mod {
				result = append(result, interval+1)
				continue
			}
			result = append(result, interval)
		}
	}

	return result

}
