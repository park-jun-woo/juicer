//ff:func feature=scan type=convert control=sequence
//ff:what Request에서 OpenAPI requestBody를 생성한다
package scanner

func buildRequestBody(req *Request, schemas map[string]any) map[string]any {
	if len(req.Files) > 0 || len(req.FormFields) > 0 {
		return buildMultipartBody(req)
	}

	if req.Body != nil {
		schema := bodySchema(req.Body, schemas)
		return map[string]any{
			"required": true,
			"content": map[string]any{
				"application/json": map[string]any{
					"schema": schema,
				},
			},
		}
	}

	if req.RawBody {
		return map[string]any{
			"required": true,
			"content": map[string]any{
				"application/octet-stream": map[string]any{
					"schema": map[string]any{"type": "string", "format": "binary"},
				},
			},
		}
	}

	return nil
}
