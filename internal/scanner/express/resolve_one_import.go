//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 단일 import 구문에서 변수명과 파일 경로를 추출하여 매핑에 추가한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func resolveOneImport(stmt *sitter.Node, src []byte, dir string, imports map[string]string, absRoot string, aliases map[string]string) {
	importPath := extractImportPath(stmt, src)
	if importPath == "" {
		return
	}
	resolved := resolveRelativePath(dir, importPath)
	if resolved == "" {
		resolved = resolvePathAlias(absRoot, importPath, aliases)
	}
	if resolved == "" {
		return
	}
	// default + named import의 로컬 바인딩명을 모두 같은 파일로 매핑한다.
	// (import { aRouter, bRouter } from "..." 처럼 한 구문에 여러 라우터가 있어도 누락 없음)
	for _, name := range collectImportedNames(stmt, src) {
		if name != "" {
			imports[name] = resolved
		}
	}
}
