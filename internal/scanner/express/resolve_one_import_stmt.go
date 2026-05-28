//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 하나의 import 문에서 미해결 스키마가 있으면 해당 파일을 추가 파싱한다
package express

import (
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"
)

func resolveOneImportStmt(ctx *scanContext, fi *fileInfo, stmt *sitter.Node, unresolvedSet map[string]bool) {
	importedNames := collectImportedNames(stmt, fi.Src)
	if !hasNeededName(importedNames, unresolvedSet) {
		return
	}
	importPath := extractImportPath(stmt, fi.Src)
	if importPath == "" {
		return
	}
	dir := filepath.Dir(fi.Path)
	resolved := resolveRelativePath(dir, importPath)
	if resolved == "" {
		resolved = resolvePathAlias(ctx.absRoot, importPath, ctx.pathAliases)
	}
	if resolved == "" || ctx.parsed[resolved] != nil {
		return
	}
	extraFi, err := parseFile(resolved)
	if err != nil {
		return
	}
	ctx.parsed[resolved] = extraFi
	mergeFileSchemas(ctx, extraFi)
}
