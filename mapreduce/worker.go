package mapreduce

func Execute(text string) map[string]int {
	// Split text into chunks (for simplicity, we'll process it as a single chunk)
	chunks := []string{text}

	// Map phase
	var intermediateResults []map[string]int
	for _, chunk := range chunks {
		intermediateResults = append(intermediateResults, Map(chunk))
	}

	// Reduce phase
	finalResult := Reduce(intermediateResults)

	return finalResult
}
