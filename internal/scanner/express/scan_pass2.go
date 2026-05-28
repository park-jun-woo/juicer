//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what Pass 2: 각 파일에서 라우트를 추출하고 prefix를 적용하여 Endpoint를 생성한다
package express

import (
	"path/filepath"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func scanPass2(ctx *scanContext, absRoot string) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for path, fi := range ctx.parsed {
		routers := ctx.allRouters[path]
		if len(routers) == 0 {
			continue
		}
		prefix := ctx.prefixMap[path]
		relPath := path
		if rel, err := filepath.Rel(absRoot, path); err == nil {
			relPath = rel
		}
		eps := buildEndpointsFromFile(fi, routers, prefix, relPath, ctx)
		endpoints = append(endpoints, eps...)
	}
	return endpoints
}
