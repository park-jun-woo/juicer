//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 단일 call 노드를 include_router 호출로 파싱 시도한다
package fastapi

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// tryParseIncludeRouter tries to parse a single call as an include_router invocation.
func tryParseIncludeRouter(call *sitter.Node, src []byte) *includeCall {
	attr := findChildByType(call, "attribute")
	if attr == nil {
		return nil
	}
	attrText := nodeText(attr, src)
	if !strings.HasSuffix(attrText, ".include_router") {
		return nil
	}
	parentVar := strings.TrimSuffix(attrText, ".include_router")

	args := findChildByType(call, "argument_list")
	if args == nil {
		return nil
	}
	childVar := firstIdentArg(args, src)
	if childVar == "" {
		return nil
	}
	return &includeCall{
		parentVar:   parentVar,
		childVar:    childVar,
		extraPrefix: extractKeywordArg(args, "prefix", src),
	}
}
