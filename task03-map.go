package homework

func sortMapValues(input map[int]string) (result []string) {
	//Place your code
	var keys []int
	var values []string

	for key, val := range input {
		keys = append(keys, key)
		values = append(values, val)
	}

	for i := 0; i < len(keys)-1; i++ {
		r := 0
		for r < len(keys)-1 {
			if keys[r] > keys[r+1] {
				keys[r], keys[r+1] = keys[r+1], keys[r]
				values[r], values[r+1] = values[r+1], values[r]
			}
			r++
		}
	}
	return values
}
