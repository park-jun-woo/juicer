//ff:func feature=scan type=convert control=iteration dimension=1 topic=flask
//ff:what JSON 바디 존재 시 ep.Request.Body를 get_json 메서드로 설정한다
package flask

import "github.com/park-jun-woo/codistill/internal/scanner"

// applyJSONBody sets an application/json body (with inline fields) when the
// handler accessed request.json / request.get_json(). TypeName stays empty so
// bodySchema renders an inline object schema.
func applyJSONBody(ep *scanner.Endpoint, ri routeInfo) {
	if !ri.hasJSONBody {
		return
	}
	scanner.EnsureRequest(ep)
	var fields []scanner.Field
	for _, k := range ri.jsonFields {
		fields = append(fields, scanner.Field{Name: k, Type: "string"})
	}
	ep.Request.Body = &scanner.Body{Method: "get_json", Fields: fields}
}
