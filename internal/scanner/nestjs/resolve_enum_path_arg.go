//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 데코레이터 경로 인자(Enum.Member/식별자)를 enum 자산으로 해석한다
package nestjs

import (
	"os"
	"path/filepath"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// resolveEnumPathArg resolves a decorator path argument that is an enum member
// expression (e.g. "RouteKey.Asset") into its underlying string value
// ("assets"). It first searches the same file, then follows the import path of
// the enum name. Returns ("", false) when the argument is not a member
// expression or cannot be resolved (caller keeps the raw string as fallback).
func resolveEnumPathArg(arg string, root *sitter.Node, src []byte,
	filePath string, imports map[string]string, projectRoot string) (string, bool) {
	dot := strings.IndexByte(arg, '.')
	if dot <= 0 || dot == len(arg)-1 {
		return "", false
	}
	enumName := arg[:dot]
	memberName := arg[dot+1:]
	if strings.ContainsAny(memberName, ".([") {
		return "", false // not a simple Enum.Member expression
	}
	// 1. same file
	if v, ok := extractEnumMemberValue(root, src, enumName, memberName); ok {
		return v, true
	}
	// 2. follow import path of the enum name
	importPath, ok := imports[enumName]
	if !ok {
		return "", false
	}
	absPath := resolveImportPath(filepath.Dir(filePath), importPath, projectRoot)
	if absPath == "" {
		return "", false
	}
	enumSrc, err := os.ReadFile(absPath)
	if err != nil {
		return "", false
	}
	enumRoot, err := parseTypeScript(enumSrc)
	if err != nil {
		return "", false
	}
	return extractEnumMemberValue(enumRoot, enumSrc, enumName, memberName)
}
