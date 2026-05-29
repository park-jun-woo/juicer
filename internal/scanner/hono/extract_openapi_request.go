//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what createRoute object의 request 필드(body/query/params)에서 zod 스키마 노드를 ValidatorInfo로 추출한다
package hono

import (
	"github.com/park-jun-woo/codistill/internal/scanner/zod"
	sitter "github.com/smacker/go-tree-sitter"
)

func extractOpenAPIRequest(createObj *sitter.Node, src []byte) []zod.ValidatorInfo {
	reqVal := findObjectValueByKey(createObj, "request", src)
	if reqVal == nil || reqVal.Type() != "object" {
		return nil
	}
	sections := []struct {
		key    string
		target string
	}{
		{"body", "json"},
		{"query", "query"},
		{"params", "param"},
	}
	var validators []zod.ValidatorInfo
	for _, s := range sections {
		secVal := findObjectValueByKey(reqVal, s.key, src)
		if secVal == nil {
			continue
		}
		calls := zod.FindObjectCalls(secVal, src)
		if len(calls) == 0 {
			continue
		}
		validators = append(validators, zod.ValidatorInfo{Target: s.target, SchemaNode: calls[0]})
	}
	return validators
}
