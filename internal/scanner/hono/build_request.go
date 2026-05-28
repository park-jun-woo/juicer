//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what 경로 파라미터 + Zod validator로 Request를 생성한다
package hono

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

	schemaSrc := buildSchemaSrcMap(ctx)
	for _, v := range r.ZodValidators {
		if zod.ApplyValidator(&req, v, ctx.schemas, fi.Src, schemaSrc) {
			hasContent = true
		}
	}

	if !hasContent {
		return nil
	}
	return &req
}

