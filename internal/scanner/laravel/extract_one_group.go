//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what Route::prefix('x')->group(fn) 또는 Route::middleware([..])->group(fn) 하나를 처리한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractOneGroup handles a single Route::prefix('x')->group(fn) or
// Route::middleware(['auth'])->group(fn) call.
func extractOneGroup(mc *sitter.Node, fi fileInfo, outerPrefix string, outerMiddleware []string) []routeInfo {
	if lastMemberCallName(mc, fi.src) != "group" {
		return nil
	}

	scopedCall := findChildByType(mc, "scoped_call_expression")
	if scopedCall == nil {
		return extractChainedGroupIfPresent(mc, fi, outerPrefix, outerMiddleware)
	}

	prefix, mw := extractGroupModifier(scopedCall, fi)
	combinedPrefix := joinGroupPrefix(outerPrefix, prefix)
	combinedMW := mergeMiddleware(outerMiddleware, mw)

	closureBody := groupClosureBody(mc, fi)
	if closureBody == nil {
		return nil
	}
	return collectRoutesFromBody(closureBody, fi, combinedPrefix, combinedMW)
}
