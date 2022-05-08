package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))
	for index, value := range input {
		result[len(input)-index-1] = value
	}
	return
}
