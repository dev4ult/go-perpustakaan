package helpers

func Response(message string) map[string]any {
	return map[string]any{
		"message": message,
	}
}