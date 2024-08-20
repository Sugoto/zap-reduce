package mapreduce

func Reduce(intermediate []map[string]int) map[string]int {
	finalCounts := make(map[string]int)
	for _, counts := range intermediate {
		for word, count := range counts {
			finalCounts[word] += count
		}
	}
	return finalCounts
}
