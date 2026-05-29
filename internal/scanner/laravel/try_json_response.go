//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what response()->json($data, code) 형태의 JSON 응답을 추출한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func tryJSONResponse(retNode *sitter.Node, src []byte, text string) *scanner.Response {
	if !strings.Contains(text, "response()->json(") {
		return nil
	}
	return &scanner.Response{
		Status: jsonResponseStatus(retNode, src),
		Kind:   "json",
	}
}
