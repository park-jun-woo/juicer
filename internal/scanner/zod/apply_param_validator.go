//ff:func feature=scan type=extract control=iteration dimension=1 topic=zod
//ff:what Zod param validator 필드로 path parameter 타입을 갱신한다
package zod

import "github.com/park-jun-woo/codistill/internal/scanner"

// ApplyParamValidator — param validator → pathParams 타입 갱신
func ApplyParamValidator(req *scanner.Request, fields []scanner.Field) bool {
	changed := false
	for _, f := range fields {
		idx := findPathParamIndex(req.PathParams, f.Name)
		if idx >= 0 {
			req.PathParams[idx].Type = f.Type
		} else {
			req.PathParams = append(req.PathParams, scanner.Param{Name: f.Name, Type: f.Type})
			changed = true
		}
	}
	return changed
}
