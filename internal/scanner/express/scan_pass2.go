//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what Pass 2: 각 파일에서 라우트를 추출하고 prefix를 적용하여 Endpoint를 생성한다
package express

import (
	"path/filepath"
	"sort"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func scanPass2(ctx *scanContext, absRoot string) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	// ctx.parsed는 map이므로 키를 정렬해 결정적 순서로 순회한다.
	paths := make([]string, 0, len(ctx.parsed))
	for path := range ctx.parsed {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fi := ctx.parsed[path]
		routers := ctx.allRouters[path]
		if len(routers) == 0 {
			continue
		}
		relPath := path
		if rel, err := filepath.Rel(absRoot, path); err == nil {
			relPath = rel
		}
		eps := buildEndpointsFromFile(fi, routers, path, relPath, ctx)
		endpoints = append(endpoints, eps...)
	}
	return endpoints
}
