//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what zValidator("json", schema) 호출에서 target과 스키마 변수명을 추출한다
package hono

import (
	"github.com/park-jun-woo/codistill/internal/scanner/zod"
	sitter "github.com/smacker/go-tree-sitter"
)

func extractZodValidators(argNodes []*sitter.Node, src []byte) []zod.ValidatorInfo {
	var validators []zod.ValidatorInfo
	for _, arg := range argNodes {
		v := extractOneZodValidator(arg, src)
		if v != nil {
			validators = append(validators, *v)
		}
	}
	return validators
}
