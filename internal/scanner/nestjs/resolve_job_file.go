//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what schemaJob의 타입명을 import 경로로 절대 파일 경로로 해석한다
package nestjs

import "path/filepath"

// resolveJobFile resolves the absolute source file path that defines the job's
// type, via its import path. Returns "" when the type is not imported or the
// path cannot be resolved.
func resolveJobFile(j schemaJob) string {
	importPath := j.imports[j.typeName]
	if importPath == "" {
		return ""
	}
	return resolveImportPath(filepath.Dir(j.referrer), importPath, j.projectRoot)
}
