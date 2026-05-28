//ff:func feature=scan type=extract control=iteration dimension=2 topic=hono
//ff:what 파싱된 파일들에서 스키마 이름과 소스 바이트의 매핑을 생성한다
package hono

import "github.com/park-jun-woo/codistill/internal/scanner/zod"

func buildSchemaSrcMap(ctx *scanContext) map[string][]byte {
	srcMap := make(map[string][]byte)
	for _, pfi := range ctx.parsed {
		schemas := zod.CollectSchemas(pfi.Root, pfi.Src)
		for name := range schemas {
			srcMap[name] = pfi.Src
		}
	}
	return srcMap
}
