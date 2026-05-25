//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 응답 목록에서 OpenAPI responses 객체를 생성한다
package scanner

func buildResponses(responses []Response, schemas map[string]any) map[string]any {
	if len(responses) == 0 {
		return map[string]any{
			"200": map[string]any{"description": "OK"},
		}
	}

	grouped := map[string][]Response{}
	for _, r := range responses {
		grouped[r.Status] = append(grouped[r.Status], r)
	}

	result := map[string]any{}
	for status, resps := range grouped {
		resp := pickBestResponse(resps)
		oaResp := map[string]any{
			"description": statusDescription(status),
		}

		if resp.Kind == "json" && (len(resp.Fields) > 0 || resp.TypeName != "") {
			schema := responseSchema(resp, schemas)
			if resp.Confidence == "partial" {
				schema["x-schema-confidence"] = "partial"
			}
			oaResp["content"] = map[string]any{
				"application/json": map[string]any{
					"schema": schema,
				},
			}
		}

		result[status] = oaResp
	}

	return result
}

