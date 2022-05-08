package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	res := make([]int64, len(input))
	lastIndex := len(input) - 1
	for i := lastIndex; i >= 0; i-- {
		res[lastIndex-i] = input[i]
	}

	return res
}
