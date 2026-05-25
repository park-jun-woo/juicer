//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what 폼 필드와 파일 업로드에서 multipart/form-data requestBody를 생성한다
package scanner

// buildMultipartBody creates a multipart/form-data requestBody from form fields and files.
func buildMultipartBody(req *Request) map[string]any {
	props := map[string]any{}
	for _, f := range req.FormFields {
		props[f.Name] = map[string]any{"type": "string"}
	}
	for _, f := range req.Files {
		props[f.Name] = map[string]any{"type": "string", "format": "binary"}
	}
	return map[string]any{
		"required": true,
		"content": map[string]any{
			"multipart/form-data": map[string]any{
				"schema": map[string]any{
					"type":       "object",
					"properties": props,
				},
			},
		},
	}
}
