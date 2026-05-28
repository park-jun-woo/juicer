//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 핸들러 함수의 본문(statement_block)을 찾는다 — 인라인 또는 같은 파일 내 이름 참조
package express

import sitter "github.com/smacker/go-tree-sitter"

func findHandlerBody(fi *fileInfo, ri routeInfo) *sitter.Node {
	if ri.HandlerNode != nil {
		body := extractFunctionBody(ri.HandlerNode)
		if body != nil {
			return body
		}
	}
	if ri.Handler == "" || ri.Handler == "(anonymous)" {
		return nil
	}
	return findNamedFunctionBody(fi, ri.Handler)
}
