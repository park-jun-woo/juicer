//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what validateRequest({body: X, query: Y}) 호출을 파싱하여 []zod.ValidatorInfo를 반환한다
package express

import (
	"github.com/park-jun-woo/codistill/internal/scanner/zod"
	sitter "github.com/smacker/go-tree-sitter"
)

// validateRequestFunctions — validateRequest 미들웨어 함수명 목록
var validateRequestFunctions = map[string]bool{
	"validateRequest": true,
	"validate":        true,
	"zodValidate":     true,
}

// validateRequestTargetMap — pair key → ValidatorInfo target 매핑
var validateRequestTargetMap = map[string]string{
	"body":   "json",
	"query":  "query",
	"param":  "param",
	"params": "param",
}

func extractValidateRequest(node *sitter.Node, src []byte) []zod.ValidatorInfo {
	if node.Type() != "call_expression" {
		return nil
	}
	fn := findChildByType(node, "identifier")
	if fn == nil {
		return nil
	}
	if !validateRequestFunctions[nodeText(fn, src)] {
		return nil
	}
	args := findChildByType(node, "arguments")
	if args == nil {
		return nil
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) < 1 {
		return nil
	}
	obj := argNodes[0]
	if obj.Type() != "object" {
		return nil
	}
	var validators []zod.ValidatorInfo
	for _, pair := range childrenOfType(obj, "pair") {
		keyNode := pair.ChildByFieldName("key")
		if keyNode == nil {
			continue
		}
		key := nodeText(keyNode, src)
		target, ok := validateRequestTargetMap[key]
		if !ok {
			continue
		}
		valueNode := pair.ChildByFieldName("value")
		if valueNode == nil {
			continue
		}
		info := zod.ValidatorInfo{Target: target}
		switch valueNode.Type() {
		case "identifier":
			info.SchemaName = nodeText(valueNode, src)
		default:
			info.SchemaNode = valueNode
		}
		validators = append(validators, info)
	}
	return validators
}
