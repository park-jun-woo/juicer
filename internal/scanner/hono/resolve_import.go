//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what import 구문에서 변수명→파일 경로 매핑을 수집한다
package hono

import "path/filepath"

func resolveImports(fi *fileInfo, absRoot string) map[string]string {
	imports := make(map[string]string)
	dir := filepath.Dir(fi.Path)
	for _, stmt := range findAllByType(fi.Root, "import_statement") {
		resolveOneImport(stmt, fi.Src, dir, imports, absRoot)
	}
	return imports
}
