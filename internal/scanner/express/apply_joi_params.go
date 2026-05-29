//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what Joi params 필드를 중복 없이 Request.PathParams에 추가한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

func applyJoiParams(req *scanner.Request, fields []scanner.Field) bool {
	changed := false
	for _, f := range fields {
		if pathParamExists(req.PathParams, f.Name) {
			continue
		}
		req.PathParams = append(req.PathParams, scanner.Param{Name: f.Name, Type: f.Type})
		changed = true
	}
	return changed
}
