//ff:func feature=scan type=extract control=sequence topic=flask
//ff:what register_blueprint 호출에서 Blueprint 변수명과 prefix 오버라이드를 파싱한다
package flask

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// tryParseRegisterBlueprint tries to parse a register_blueprint call.
// Returns (blueprintVarName, overridePrefix).
// e.g., app.register_blueprint(users_bp) -> ("users_bp", "")
// e.g., app.register_blueprint(users_bp, url_prefix="/v2") -> ("users_bp", "/v2")
func tryParseRegisterBlueprint(call *sitter.Node, src []byte) (string, string) {
	attrNode := findChildByType(call, "attribute")
	if attrNode == nil {
		return "", ""
	}
	attrText := nodeText(attrNode, src)
	if !strings.HasSuffix(attrText, ".register_blueprint") {
		return "", ""
	}

	args := findChildByType(call, "argument_list")
	if args == nil {
		return "", ""
	}

	// First positional argument is the blueprint variable
	bpVarName := firstIdentArg(args, src)
	if bpVarName == "" {
		return "", ""
	}

	// Check for url_prefix override
	overridePrefix := extractKeywordArg(args, "url_prefix", src)

	return bpVarName, overridePrefix
}
