//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 파일에서 @fastify/autoload 에 바인딩된 식별자 집합을 수집한다
package fastify

const autoloadModule = "@fastify/autoload"

func autoloadVarNames(fi *fileInfo) map[string]bool {
	names := make(map[string]bool)
	for _, stmt := range findAllByType(fi.Root, "import_statement") {
		if extractImportPath(stmt, fi.Src) != autoloadModule {
			continue
		}
		if v := extractImportVarName(stmt, fi.Src); v != "" {
			names[v] = true
		}
	}
	collectAutoloadRequireNames(fi, names)
	return names
}
