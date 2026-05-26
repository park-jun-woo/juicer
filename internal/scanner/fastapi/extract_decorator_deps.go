//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 데코레이터 call 노드에서 dependencies 키워드 인자의 Depends 목록을 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractDecoratorDeps extracts Depends function names from the dependencies=
// keyword argument in a route decorator's argument list.
// e.g., @router.get("/", dependencies=[Depends(verify_token), Depends(log_request)])
// returns ["verify_token", "log_request"].
func extractDecoratorDeps(callNode *sitter.Node, src []byte) []string {
	if callNode == nil {
		return nil
	}
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return nil
	}
	return findDependenciesKeyword(args, src)
}
