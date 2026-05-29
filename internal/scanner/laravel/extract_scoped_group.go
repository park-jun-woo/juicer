//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what Route::group(['prefix'=>..], fn) scoped 호출 하나를 그룹 라우트로 펼친다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractScopedGroup handles a single Route::group(['prefix' => 'x', ...], fn)
// scoped_call_expression whose options are given as an array literal.
func extractScopedGroup(call *sitter.Node, fi fileInfo, outerPrefix string, outerMiddleware []string) []routeInfo {
	scope := findChildByType(call, "name")
	if scope == nil || nodeText(scope, fi.src) != "Route" {
		return nil
	}
	if secondScopedName(call, fi.src) != "group" {
		return nil
	}

	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	argNodes := childrenOfType(args, "argument")
	if len(argNodes) < 2 {
		return nil
	}
	arr := findChildByType(argNodes[0], "array_creation_expression")
	if arr == nil {
		return nil
	}

	prefix, mw := extractGroupArrayModifier(arr, fi)
	combinedPrefix := joinGroupPrefix(outerPrefix, prefix)
	combinedMW := mergeMiddleware(outerMiddleware, mw)

	closureBody := groupClosureBody(call, fi)
	if closureBody == nil {
		return nil
	}
	return collectRoutesFromBody(closureBody, fi, combinedPrefix, combinedMW)
}
