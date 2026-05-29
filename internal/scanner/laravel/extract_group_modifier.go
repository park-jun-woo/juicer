//ff:func feature=scan type=extract control=selection topic=laravel
//ff:what Route::prefix('x') 또는 Route::middleware([..])에서 prefix/middleware를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractGroupModifier extracts prefix or middleware from Route::prefix('x') or Route::middleware([...]).
func extractGroupModifier(scopedCall *sitter.Node, fi fileInfo) (string, []string) {
	args := findChildByType(scopedCall, "arguments")
	if args == nil {
		return "", nil
	}
	switch secondScopedName(scopedCall, fi.src) {
	case "prefix":
		if p, ok := firstArgString(args, fi.src); ok {
			return p, nil
		}
	case "middleware":
		return "", extractMiddlewareArgs(args, fi)
	}
	return "", nil
}
