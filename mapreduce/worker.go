package mapreduce

func Execute(text string) map[string]int {
	chunks := []string{text}

	var intermediateResults []map[string]int
	for _, chunk := range chunks {
		intermediateResults = append(intermediateResults, Map(chunk))
	}

	finalResult := Reduce(intermediateResults)

	return finalResult
}
