//ff:func feature=scan type=extract control=selection topic=hono
//ff:what 단일 call_expression에서 zValidator() 호출을 파싱하여 target + schema를 추출한다
package hono

import (
	"github.com/park-jun-woo/codistill/internal/scanner/zod"
	sitter "github.com/smacker/go-tree-sitter"
)

func extractOneZodValidator(arg *sitter.Node, src []byte) *zod.ValidatorInfo {
	if arg.Type() != "call_expression" {
		return nil
	}
	fn := findChildByType(arg, "identifier")
	if fn == nil {
		return nil
	}
	if nodeText(fn, src) != "zValidator" {
		return nil
	}
	args := findChildByType(arg, "arguments")
	if args == nil {
		return nil
	}
	innerArgs := collectArgNodes(args)
	if len(innerArgs) < 2 {
		return nil
	}
	targetNode := innerArgs[0]
	if targetNode.Type() != "string" {
		return nil
	}
	target := unquoteTS(nodeText(targetNode, src))
	info := zod.ValidatorInfo{Target: target}
	schemaArg := innerArgs[1]
	switch schemaArg.Type() {
	case "identifier":
		info.SchemaName = nodeText(schemaArg, src)
	default:
		info.SchemaNode = schemaArg
	}
	return &info
}
