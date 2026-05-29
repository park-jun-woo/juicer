//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 호출의 arguments에서 첫 문자열 리터럴 내용을 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func extractFirstStringArg(callExpr *sitter.Node, src []byte) string {
	args := findChildByType(callExpr, "arguments")
	if args == nil {
		return ""
	}
	strLit := findChildByType(args, "string_literal")
	if strLit == nil {
		return ""
	}
	strContent := findChildByType(strLit, "string_content")
	if strContent == nil {
		return ""
	}
	return nodeText(strContent, src)
}
