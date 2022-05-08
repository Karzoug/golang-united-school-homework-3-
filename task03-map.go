package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	result = make([]string, len(input))
	for index, value := range keys {
		result[index] = input[value]
	}
	return
}
