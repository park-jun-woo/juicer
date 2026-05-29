//ff:func feature=scan type=extract control=selection topic=zod
//ff:what 단일 Zod validator 결과를 Request에 반영한다 (json→body, query→query, param→pathParams, form→formFields)
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// ApplyValidator — ValidatorInfo → Request 반영
func ApplyValidator(req *scanner.Request, v ValidatorInfo, schemas map[string]*sitter.Node, inlineSrc []byte, schemaSrc map[string][]byte) bool {
	fields := ResolveValidatorFields(v, schemas, inlineSrc, schemaSrc)
	switch v.Target {
	case "json":
		if len(fields) > 0 {
			req.Body = &scanner.Body{Method: "json", TypeName: v.SchemaName, Fields: fields}
			return true
		}
	case "query":
		for _, f := range fields {
			req.Query = append(req.Query, scanner.Param{Name: f.Name, Type: f.Type})
		}
		return len(fields) > 0
	case "param":
		return ApplyParamValidator(req, fields)
	case "form":
		for _, f := range fields {
			req.FormFields = append(req.FormFields, scanner.Param{Name: f.Name, Type: f.Type})
		}
		return len(fields) > 0
	}
	return false
}
