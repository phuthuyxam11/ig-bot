package helpers

func MapToStringSlice(m map[int32]string) []string {
	result := make([]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func StringSliceToMap(m []string) map[int32]string {
	result := make(map[int32]string, len(m))
	count := 0
	for _, v := range m {
		result[int32(count)] = v
		count++
	}
	return result
}
