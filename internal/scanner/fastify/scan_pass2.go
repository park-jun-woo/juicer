//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what Pass 2: 각 파일에서 라우트를 추출하고 prefix 모델을 적용하여 Endpoint를 생성한다
package fastify

import "github.com/park-jun-woo/codistill/internal/scanner"

func scanPass2(ctx *scanContext) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for path, fi := range ctx.parsed {
		instances := ctx.instances[path]
		if len(instances) == 0 {
			continue
		}
		endpoints = append(endpoints, scanFileEndpoints(ctx, path, fi, instances)...)
	}
	return endpoints
}
