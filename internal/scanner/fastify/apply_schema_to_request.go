//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what schemaInfo의 params/querystring/body를 Request에 적용한다 (TypeBox 변수 참조 해석 포함)
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func applySchemaToRequest(si *schemaInfo, src []byte, req *scanner.Request, hasReq *bool, vars map[string]*sitter.Node) {
	if si.Params != nil {
		if p := schemaNodeToParams(si.Params, src, vars); len(p) > 0 {
			req.PathParams = p
			*hasReq = true
		}
	}
	if si.Querystring != nil {
		if q := schemaNodeToParams(si.Querystring, src, vars); len(q) > 0 {
			req.Query = q
			*hasReq = true
		}
	}
	if si.Body != nil {
		if fields := schemaNodeToFields(si.Body, src, vars); len(fields) > 0 {
			req.Body = &scanner.Body{VarName: "body", Method: "json", Fields: fields}
			*hasReq = true
		}
	}
}
