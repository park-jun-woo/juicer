//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what main.ts에서 setGlobalPrefix 호출을 찾아 전역 접두사를 반환한다
package nestjs

import (
	"os"
	"path/filepath"
)

// detectGlobalPrefix searches for setGlobalPrefix('prefix') in main.ts.
// When the argument is not a string literal (e.g. configService.getOrThrow(...)),
// it falls back to .env.example or config file defaults.
func detectGlobalPrefix(root string) string {
	mainPath := filepath.Join(root, "src", "main.ts")
	src, err := os.ReadFile(mainPath)
	if err != nil {
		return ""
	}
	astRoot, err := parseTypeScript(src)
	if err != nil {
		return ""
	}
	calls := findAllByType(astRoot, "call_expression")
	found := false
	for _, call := range calls {
		if prefix, ok := trySetGlobalPrefix(call, src); ok {
			return prefix
		}
		if hasSetGlobalPrefix(call, src) {
			found = true
		}
	}
	if found {
		return fallbackGlobalPrefix(root)
	}
	return ""
}
