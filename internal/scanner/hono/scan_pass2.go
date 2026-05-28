//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what Pass 2: 각 파일에서 라우트를 추출하고 prefix + Zod 스키마를 적용하여 Endpoint를 생성한다
package hono

import (
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func scanPass2(ctx *scanContext) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for path, fi := range ctx.parsed {
		vars := ctx.honoVars[path]
		if len(vars) == 0 {
			continue
		}
		relPath := path
		if rel, err := filepath.Rel(ctx.absRoot, path); err == nil {
			relPath = rel
		}
		routes := collectRoutes(fi, vars)
		for _, r := range routes {
			eps := buildEndpointsFromRoute(r, vars, ctx, fi, relPath)
			endpoints = append(endpoints, eps...)
		}
	}
	return endpoints
}
