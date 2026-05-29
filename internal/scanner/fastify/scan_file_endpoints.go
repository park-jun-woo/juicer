//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 단일 파일의 라우트마다 적용 prefix 목록을 구해 prefix별로 Endpoint를 생성한다
package fastify

import (
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func scanFileEndpoints(ctx *scanContext, path string, fi *fileInfo, instances map[string]bool) []scanner.Endpoint {
	relPath := path
	if rel, err := filepath.Rel(ctx.absRoot, path); err == nil {
		relPath = rel
	}
	filePfx := ctx.prefixMap[path]
	scopes := ctx.wrappers[path]
	typeBoxVars := extractTypeBoxVars(fi)
	var endpoints []scanner.Endpoint
	for _, r := range extractRoutes(fi, instances) {
		for _, prefix := range routePrefixes(r, filePfx, scopes) {
			endpoints = append(endpoints, buildEndpointsFromRoute(r, prefix, relPath, fi.Src, typeBoxVars)...)
		}
	}
	return endpoints
}
