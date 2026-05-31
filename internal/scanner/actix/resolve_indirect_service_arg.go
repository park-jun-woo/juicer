//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what .service(fn()) 일반 함수 호출 인자를 함수 인덱스에서 되짚어 반환 scope 체인으로 재귀한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// resolveIndirectServiceArg handles `.service(xxx_scope())`. If callExpr is a
// plain function call (head is a bare identifier) whose name is registered in
// the cross-file handlerFuncs index, it extracts the function's returned
// web::scope(...) chain and re-runs processServiceArg on it, so the existing
// web::scope branch synthesizes the scope prefix. A visited-name set guards
// against cyclic (mutually recursive) scope definitions.
func resolveIndirectServiceArg(callExpr *sitter.Node, src []byte, prefix string, routes *[]builderRoute, handlerFuncs map[string]*handlerInfo, visited map[string]bool) {
	if callExpr.Type() != "call_expression" {
		return
	}
	head := callExpr.Child(0)
	if head == nil || head.Type() != "identifier" {
		return
	}
	name := nodeText(head, src)
	info := handlerFuncs[name]
	if info == nil {
		return
	}
	if visited[name] {
		return
	}
	chain := funcReturnScopeChain(info.funcNode, info.file.src)
	if chain == nil {
		return
	}
	visited[name] = true
	processServiceArg(chain, info.file.src, prefix, routes, handlerFuncs, visited)
	delete(visited, name)
}
