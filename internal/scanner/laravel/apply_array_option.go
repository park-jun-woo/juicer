//ff:func feature=scan type=extract control=selection topic=laravel
//ff:what 'prefix'/'middleware' 배열 항목 하나를 prefix/middleware 누적기에 적용한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// applyArrayOption applies a single 'key' => value array option (prefix or
// middleware) from a Route::group([...], fn) options array.
func applyArrayOption(elem *sitter.Node, fi fileInfo, prefix *string, mw *[]string) {
	keyNode := findChildByType(elem, "string")
	if keyNode == nil {
		return
	}
	value := groupArrayValue(elem)
	if value == nil {
		return
	}
	switch extractStringContent(keyNode, fi.src) {
	case "prefix":
		*prefix = joinGroupPrefix(*prefix, extractStringContent(value, fi.src))
	case "middleware":
		*mw = mergeMiddleware(*mw, middlewareValues(value, fi))
	}
}
