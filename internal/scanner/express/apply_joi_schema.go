//ff:func feature=scan type=extract control=sequence topic=express
//ff:what joi.RequestSchema를 Request에 반영한다 (body→Body, query→Query, params→PathParams)
package express

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"github.com/park-jun-woo/codistill/internal/scanner/joi"
)

// applyJoiSchema — RequestSchema를 Request에 반영하고, 무언가 채웠으면 true를 반환한다.
func applyJoiSchema(req *scanner.Request, rs joi.RequestSchema) bool {
	changed := false
	if len(rs.Body) > 0 {
		req.Body = &scanner.Body{Method: "json", Fields: rs.Body}
		changed = true
	}
	if applyJoiQuery(req, rs.Query) {
		changed = true
	}
	if applyJoiParams(req, rs.Params) {
		changed = true
	}
	return changed
}
