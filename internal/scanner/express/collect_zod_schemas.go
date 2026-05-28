//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 파일 내 z.object() Zod 스키마를 수집한다 (zod 패키지 위임)
package express

import (
	"github.com/park-jun-woo/codistill/internal/scanner/zod"
	sitter "github.com/smacker/go-tree-sitter"
)

func collectZodSchemas(fi *fileInfo) map[string]*sitter.Node {
	return zod.CollectSchemas(fi.Root, fi.Src)
}
