//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what Endpoint의 요청에서 OpenAPI parameters 배열을 생성한다
package scanner

// buildOperationParams creates the OpenAPI parameters array from path params and query params.
func buildOperationParams(req *Request) []map[string]any {
	if req == nil {
		return nil
	}
	var params []map[string]any
	for _, p := range req.PathParams {
		params = append(params, map[string]any{
			"name":     p.Name,
			"in":       "path",
			"required": true,
			"schema":   buildParamSchema(p.Type),
		})
	}
	for _, q := range req.Query {
		qp := map[string]any{
			"name":   q.Name,
			"in":     "query",
			"schema": buildParamSchema(q.Type),
		}
		if q.DefaultIsNull {
			qp["schema"].(map[string]any)["default"] = nil
			qp["schema"].(map[string]any)["nullable"] = true
		} else if q.Default != "" {
			qp["schema"].(map[string]any)["default"] = q.Default
		}
		params = append(params, qp)
	}
	return params
}
