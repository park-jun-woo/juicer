//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what assignment 노드를 Blueprint 생성자 호출로 파싱을 시도한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// tryParseBlueprintAssignment tries to parse an assignment as a Blueprint constructor call.
func tryParseBlueprintAssignment(assign *sitter.Node, src []byte) *blueprintInfo {
	// Left side: identifier (variable name)
	leftNodes := childrenOfType(assign, "identifier")
	if len(leftNodes) == 0 {
		return nil
	}
	varName := nodeText(leftNodes[0], src)

	// Right side: call node
	callNode := findChildByType(assign, "call")
	if callNode == nil {
		return nil
	}

	// Check if it's Blueprint(...)
	funcNode := findChildByType(callNode, "identifier")
	if funcNode == nil || nodeText(funcNode, src) != "Blueprint" {
		return nil
	}

	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return nil
	}

	// First positional string arg is the blueprint name
	name := firstStringArg(args, src)

	// Extract url_prefix keyword argument
	urlPrefix := extractKeywordArg(args, "url_prefix", src)

	return &blueprintInfo{
		varName:   varName,
		name:      name,
		urlPrefix: urlPrefix,
	}
}
