//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what Joi query 필드를 Request.Query에 추가한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

func applyJoiQuery(req *scanner.Request, fields []scanner.Field) bool {
	for _, f := range fields {
		req.Query = append(req.Query, scanner.Param{Name: f.Name, Type: f.Type})
	}
	return len(fields) > 0
}
