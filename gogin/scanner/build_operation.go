//ff:func feature=scan type=convert control=sequence
//ff:what 단일 Endpoint에서 OpenAPI operation 객체를 생성한다
package scanner

func buildOperation(ep Endpoint, schemas map[string]any) map[string]any {
	op := map[string]any{
		"operationId": generateOperationID(ep),
	}

	if params := buildOperationParams(ep.Request); len(params) > 0 {
		op["parameters"] = params
	}

	if ep.Request != nil {
		if rb := buildRequestBody(ep.Request, schemas); rb != nil {
			op["requestBody"] = rb
		}
	}

	op["responses"] = buildResponses(ep.Responses, schemas)

	return op
}
