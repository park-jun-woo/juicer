//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what pathParams + Zod validator로 Request를 생성한다
package express

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"github.com/park-jun-woo/codistill/internal/scanner/zod"
)

func buildRequest(r routeInfo, pathParams []string, ctx *scanContext, fi *fileInfo) *scanner.Request {
	var req scanner.Request
	hasContent := false

	if len(pathParams) > 0 {
		for _, p := range pathParams {
			req.PathParams = append(req.PathParams, scanner.Param{Name: p, Type: "string"})
		}
		hasContent = true
	}

	for _, v := range r.ZodValidators {
		if zod.ApplyValidator(&req, v, ctx.schemas, fi.Src, ctx.schemaSrc) {
			hasContent = true
		}
	}

	if !hasContent {
		return nil
	}
	return &req
}
