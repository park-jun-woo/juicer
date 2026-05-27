//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what dtoField 배열에서 enumTypeName이 있는 필드의 enum 멤버를 해석한다
package nestjs

import (
	"os"
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"
)

// resolveEnumTypeNames resolves enum member values for fields that have an enumTypeName set.
// It first searches the same file, then follows import paths to find cross-file enum definitions.
func resolveEnumTypeNames(fields []dtoField, root *sitter.Node, src []byte,
	filePath string, imports map[string]string, projectRoot string) {
	for i := range fields {
		if fields[i].enumTypeName == "" || len(fields[i].enum) > 0 {
			continue
		}
		// 1. 같은 파일에서 탐색
		vals := extractEnumMembers(root, src, fields[i].enumTypeName)
		if vals != nil {
			fields[i].enum = vals
			continue
		}
		// 2. import 경로에서 탐색
		importPath, ok := imports[fields[i].enumTypeName]
		if !ok {
			continue
		}
		referrerDir := filepath.Dir(filePath)
		absPath := resolveImportPath(referrerDir, importPath, projectRoot)
		if absPath == "" {
			continue
		}
		enumSrc, err := os.ReadFile(absPath)
		if err != nil {
			continue
		}
		enumRoot, err := parseTypeScript(enumSrc)
		if err != nil {
			continue
		}
		vals = extractEnumMembers(enumRoot, enumSrc, fields[i].enumTypeName)
		if vals != nil {
			fields[i].enum = vals
		}
	}
}
