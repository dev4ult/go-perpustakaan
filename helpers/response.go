package helpers

func Response(message string, data any, err ...any) map[string]any {

	var res = map[string]any{
		"message": message,
	}

	if data != nil {
		res["data"] = data
	}

	if err != nil {
		res["error"] = err
	}

	return res
}