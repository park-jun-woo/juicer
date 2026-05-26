//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 데코레이터 인자에서 path, status_code, response_model, response_class를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractDecoratorArgs extracts path, status_code, response_model, and response_class
// from decorator arguments.
func extractDecoratorArgs(callNode *sitter.Node, src []byte) (string, int, string, string) {
	if callNode == nil {
		return "", 0, "", ""
	}
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return "", 0, "", ""
	}

	path := firstStringArg(args, src)
	statusCode := parseIntDefault(extractKeywordArg(args, "status_code", src), 0)
	responseModel := extractKeywordArg(args, "response_model", src)
	responseClass := extractKeywordArg(args, "response_class", src)

	return path, statusCode, responseModel, responseClass
}
