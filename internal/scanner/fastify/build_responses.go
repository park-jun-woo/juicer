//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 라우트 스키마에서 응답 정보를 생성한다
package fastify

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildResponses(r routeInfo, src []byte) []scanner.Response {
	if r.Schema == nil {
		return nil
	}
	si := extractJSONSchema(r.Schema, src)
	if si == nil {
		return nil
	}
	var responses []scanner.Response
	for status, respNode := range si.Response {
		resp := scanner.Response{Status: status, Kind: "json"}
		fields := jsonSchemaToFields(respNode, src)
		if len(fields) > 0 {
			resp.Fields = fields
		}
		responses = append(responses, resp)
	}
	return responses
}
