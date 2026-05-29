//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what Route::group(['prefix'=>..,'middleware'=>..], fn) 옵션 배열에서 prefix/middleware를 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractGroupArrayModifier extracts prefix and middleware from the options
// array of a Route::group(['prefix' => 'x', 'middleware' => [...]], fn) call.
func extractGroupArrayModifier(arr *sitter.Node, fi fileInfo) (string, []string) {
	var prefix string
	var mw []string
	for _, elem := range childrenOfType(arr, "array_element_initializer") {
		applyArrayOption(elem, fi, &prefix, &mw)
	}
	return prefix, mw
}
