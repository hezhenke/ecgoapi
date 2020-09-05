package dtos

const (
	UNKNOWN_ERROR   = 10000
	NOT_FOUND       = 404
)

func FormatBody(m map[string]interface{}) map[string]interface{}{
	m["error_code"] = 0
	return m
}

func FormatError(code int,message string)map[string]interface{} {
	switch code {
	case UNKNOWN_ERROR:
		message = "message.error.unknown"
		break
	case NOT_FOUND:
		message = "message.error.404"
	}

	return map[string]interface{}{
		"error":true,
		"error_code":code,
		"error_desc":message,
	}
}
