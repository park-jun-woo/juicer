//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what import 경로가 외부 패키지(node_modules)인지 판별한다
package nestjs

import "strings"

// isExternalPackage returns true if the import path refers to an external npm package
// (e.g. @nestjs/common, class-validator) rather than a project-internal path.
// Relative paths (./foo) and project-root paths (src/foo) return false.
func isExternalPackage(importPath string) bool {
	if strings.HasPrefix(importPath, ".") {
		return false
	}
	if strings.HasPrefix(importPath, "@/") {
		return false
	}
	if strings.HasPrefix(importPath, "@") {
		return true
	}
	if strings.HasPrefix(importPath, "src/") || strings.HasPrefix(importPath, "src\\") {
		return false
	}
	if !strings.Contains(importPath, "/") {
		return true
	}
	return false
}
