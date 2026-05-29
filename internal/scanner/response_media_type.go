//ff:func feature=scan type=convert control=selection
//ff:what 응답의 ContentType/Kind로부터 OpenAPI 미디어 타입 문자열을 결정한다
package scanner

func responseMediaType(resp Response) string {
	if resp.ContentType != "" {
		return resp.ContentType
	}
	switch resp.Kind {
	case "text", "string":
		return "text/plain"
	default:
		return "application/json"
	}
}
