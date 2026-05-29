//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what Route::prefix('x')->middleware(['y'])->group(fn) 체인을 처리한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractChainedGroup handles Route::prefix('x')->middleware(['y'])->group(fn) chains.
func extractChainedGroup(outerMC, innerMC *sitter.Node, fi fileInfo, outerPrefix string, outerMiddleware []string) []routeInfo {
	prefix := outerPrefix
	mw := copyMiddleware(outerMiddleware)

	walkChain(innerMC, fi, &prefix, &mw)

	closureBody := groupClosureBody(outerMC, fi)
	if closureBody == nil {
		return nil
	}
	return collectRoutesFromBody(closureBody, fi, prefix, mw)
}
