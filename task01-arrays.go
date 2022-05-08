package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32
	for _, val := range input {
		sum += val
	}

	return sum / float32(len(input))
}
