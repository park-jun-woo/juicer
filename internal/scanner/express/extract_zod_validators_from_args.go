//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 미들웨어 인자(argNodes[start..n-1])에서 validateRequest 호출을 찾는다
package express

import (
	"github.com/park-jun-woo/codistill/internal/scanner/zod"
	sitter "github.com/smacker/go-tree-sitter"
)

func extractZodValidatorsFromArgs(argNodes []*sitter.Node, src []byte, start int) []zod.ValidatorInfo {
	if len(argNodes) <= start {
		return nil
	}
	var validators []zod.ValidatorInfo
	for i := start; i < len(argNodes); i++ {
		node := argNodes[i]
		if node.Type() == "call_expression" {
			vr := extractValidateRequest(node, src)
			validators = append(validators, vr...)
		}
	}
	return validators
}
