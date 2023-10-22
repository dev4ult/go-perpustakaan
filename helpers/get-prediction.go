package helpers

func GetPrediction(label string) string {
	var status = ""

	switch label {
	case "Positive":
		status = "medium"
	case "Negative":
		status = "high"
	default:
		status = "low"
	}

	return status
}