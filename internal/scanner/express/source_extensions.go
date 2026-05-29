//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what Express 스캐너가 인식하는 소스 확장자 집합(.ts/.tsx/.js/.jsx/.mjs/.cjs)을 정의하고 멤버십을 판정한다 (TS/JS 공통)
package express

import "strings"

// sourceExtensions lists the source file extensions the Express scanner
// collects and resolves. TypeScript and JavaScript share identical Express
// route syntax, so the tree-sitter typescript grammar parses both. `.d.ts`
// declaration files are excluded separately (see isCollectableSource).
var sourceExtensions = []string{".ts", ".tsx", ".js", ".jsx", ".mjs", ".cjs"}

// hasSourceExtension reports whether name ends with one of the known source
// extensions.
func hasSourceExtension(name string) bool {
	for _, ext := range sourceExtensions {
		if strings.HasSuffix(name, ext) {
			return true
		}
	}
	return false
}
