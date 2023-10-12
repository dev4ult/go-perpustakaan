package helpers

func Response(message string, data any) map[string]any {
	var res = map[string]any{
		"message": message,
	}

	if data != nil {
		res["data"] = data
	}

	return res
}