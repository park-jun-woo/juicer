//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what Pass 2: 각 파일에서 라우트를 추출하고 JSON Schema를 해석하여 Endpoint를 생성한다
package fastify

import (
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func scanPass2(ctx *scanContext) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for path, fi := range ctx.parsed {
		instances := ctx.instances[path]
		if len(instances) == 0 {
			continue
		}
		prefix := ctx.prefixMap[path]
		relPath := path
		if rel, err := filepath.Rel(ctx.absRoot, path); err == nil {
			relPath = rel
		}
		routes := extractRoutes(fi, instances)
		for _, r := range routes {
			eps := buildEndpointsFromRoute(r, prefix, relPath, fi.Src)
			endpoints = append(endpoints, eps...)
		}
	}
	return endpoints
}
