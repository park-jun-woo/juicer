//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what import/require 구문에서 변수명→파일 경로 매핑을 수집한다
package express

import "path/filepath"

func resolveImports(fi *fileInfo, absRoot string, aliases map[string]string) map[string]string {
	imports := make(map[string]string)
	dir := filepath.Dir(fi.Path)
	for _, stmt := range findAllByType(fi.Root, "import_statement") {
		resolveOneImport(stmt, fi.Src, dir, imports, absRoot, aliases)
	}
	for _, decl := range findAllByType(fi.Root, "lexical_declaration") {
		resolveOneRequire(decl, fi.Src, dir, imports, absRoot, aliases)
	}
	return imports
}
