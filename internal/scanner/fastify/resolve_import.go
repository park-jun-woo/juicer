//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what import/require 구문에서 변수명 -> 파일 경로 매핑을 수집한다
package fastify

import "path/filepath"

func resolveImports(fi *fileInfo, absRoot string) map[string]string {
	imports := make(map[string]string)
	dir := filepath.Dir(fi.Path)
	for _, stmt := range findAllByType(fi.Root, "import_statement") {
		resolveOneImport(stmt, fi.Src, dir, imports)
	}
	for _, decl := range findAllByType(fi.Root, "lexical_declaration") {
		resolveRequireDecl(decl, fi.Src, dir, imports)
	}
	return imports
}
