//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what barrel 파일(index.ts)에서 export * from 문을 추적하여 className을 포함하는 실제 파일 경로를 반환한다
package nestjs

import (
	"os"
	"path/filepath"
)

// resolveBarrelExport parses a barrel file (index.ts) and follows
// "export * from './...'" statements to find the file that actually
// declares the given className.
func resolveBarrelExport(barrelPath, className string) string {
	src, err := os.ReadFile(barrelPath)
	if err != nil {
		return ""
	}
	root, err := parseTypeScript(src)
	if err != nil {
		return ""
	}
	barrelDir := filepath.Dir(barrelPath)
	for _, stmt := range findAllByType(root, "export_statement") {
		source := findChildByType(stmt, "string")
		if source == nil {
			continue
		}
		relPath := unquoteTS(nodeText(source, src))
		absPath := tryResolveTS(filepath.Join(barrelDir, relPath))
		if absPath == "" {
			continue
		}
		if hasClassDeclaration(absPath, className) {
			return absPath
		}
	}
	return ""
}
