//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 단일 import 구문에서 변수명과 파일 경로를 추출하여 매핑에 추가한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func resolveOneImport(stmt *sitter.Node, src []byte, dir string, imports map[string]string, absRoot string, aliases map[string]string) {
	varName := extractImportVarName(stmt, src)
	importPath := extractImportPath(stmt, src)
	if varName == "" || importPath == "" {
		return
	}
	resolved := resolveRelativePath(dir, importPath)
	if resolved == "" {
		resolved = resolvePathAlias(absRoot, importPath, aliases)
	}
	if resolved != "" {
		imports[varName] = resolved
	}
}
