//ff:func feature=scan type=convert control=sequence
//ff:what 단일 Endpoint에서 OpenAPI operation 객체를 생성한다
package scanner

func buildOperation(ep Endpoint, schemas map[string]any) map[string]any {
	op := map[string]any{
		"operationId": generateOperationID(ep),
	}

	if params := ensurePathParams(buildOperationParams(ep.Request), ep.Path); len(params) > 0 {
		op["parameters"] = params
	}

	if ep.Request != nil {
		if rb := buildRequestBody(ep.Request, schemas); rb != nil {
			op["requestBody"] = rb
		}
	}

	op["responses"] = buildResponses(ep.Responses, schemas)

	if isAuthEndpoint(ep) {
		op["security"] = []any{map[string]any{"bearerAuth": []any{}}}
	}

	return op
}
