//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 배럴 파일에서 `export { <name|default> as <bind> } from '<source>'`를 스캔해 importName과 일치하는 specifier의 source를 실파일 경로로 해석한다
package express

import (
	"path/filepath"
)

// resolveReexport — 배럴(re-export) 한 홉을 따라간다.
// targetFi AST에서 export 구문의 specifier 바인딩명(alias 우선, 없으면 name)이
// importName과 일치하면, 그 구문의 source 경로를 실제 소스 파일로 해석해 반환한다.
// 매칭 specifier가 없으면 빈 문자열을 반환한다.
func resolveReexport(targetFi *fileInfo, importName, absRoot string, aliases map[string]string) string {
	dir := filepath.Dir(targetFi.Path)
	for _, stmt := range findAllByType(targetFi.Root, "export_statement") {
		if !reexportHasBinding(stmt, targetFi.Src, importName) {
			continue
		}
		src := stmt.ChildByFieldName("source")
		if src == nil {
			continue
		}
		importPath := unquoteTS(nodeText(src, targetFi.Src))
		resolved := resolveRelativePath(dir, importPath)
		if resolved == "" {
			resolved = resolvePathAlias(absRoot, importPath, aliases)
		}
		if resolved != "" {
			return resolved
		}
	}
	return ""
}
