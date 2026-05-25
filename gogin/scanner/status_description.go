//ff:func feature=scan type=extract control=sequence
//ff:what HTTP 상태 코드에서 표준 설명 문자열을 반환한다
package scanner

func statusDescription(status string) string {
	descs := map[string]string{
		"200": "OK",
		"201": "Created",
		"204": "No Content",
		"301": "Moved Permanently",
		"302": "Found",
		"304": "Not Modified",
		"400": "Bad Request",
		"401": "Unauthorized",
		"403": "Forbidden",
		"404": "Not Found",
		"409": "Conflict",
		"422": "Unprocessable Entity",
		"500": "Internal Server Error",
	}
	if d, ok := descs[status]; ok {
		return d
	}
	if status == "(unknown)" {
		return "Error"
	}
	return "Response"
}

