//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what main.ts에서 enableVersioning(URI) 호출을 감지한다
package nestjs

import (
	"os"
	"path/filepath"
)

// detectURIVersioning checks if main.ts enables URI versioning.
// It returns true when enableVersioning({ type: VersioningType.URI }) is found.
func detectURIVersioning(root string) bool {
	mainPath := filepath.Join(root, "src", "main.ts")
	src, err := os.ReadFile(mainPath)
	if err != nil {
		return false
	}
	astRoot, err := parseTypeScript(src)
	if err != nil {
		return false
	}
	calls := findAllByType(astRoot, "call_expression")
	for _, call := range calls {
		if isEnableURIVersioning(call, src) {
			return true
		}
	}
	return false
}
