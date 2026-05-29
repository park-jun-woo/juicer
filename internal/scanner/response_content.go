//ff:func feature=scan type=convert control=sequence
//ff:what 단일 응답의 OpenAPI content 객체를 생성한다(empty Kind는 nil)
package scanner

func responseContent(resp Response, schemas map[string]any) map[string]any {
	if resp.Kind == "empty" {
		return nil
	}

	mediaType := responseMediaType(resp)

	if resp.Kind == "json" && (len(resp.Fields) > 0 || resp.TypeName != "") {
		schema := responseSchema(resp, schemas)
		if resp.Confidence == "partial" {
			schema["x-schema-confidence"] = "partial"
		}
		return map[string]any{mediaType: map[string]any{"schema": schema}}
	}

	if resp.Kind == "text" || resp.Kind == "string" {
		return map[string]any{mediaType: map[string]any{"schema": map[string]any{"type": "string"}}}
	}

	return nil
}
