//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what DTO 클래스를 추적하여 필드와 타입 파라미터를 해석한다
package quarkus

import (
	"os"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveClassFieldsWithParams(filePath, className string, projectRoot string, cache map[string][]scanner.Field) (classFieldsResult, error) {
	if cached, ok := cache[className]; ok {
		return classFieldsResult{fields: cached}, nil
	}
	src, err := os.ReadFile(filePath)
	if err != nil {
		return classFieldsResult{}, err
	}
	root, err := parseJava(src)
	if err != nil {
		return classFieldsResult{}, err
	}
	cache[className] = nil
	fields, typeParams := resolveClassFromASTWithParams(root, src, className, filePath, projectRoot, cache)
	if fields != nil {
		return classFieldsResult{fields: fields, typeParams: typeParams}, nil
	}
	enumFields := resolveEnumFromAST(root, src, className, cache)
	if enumFields != nil {
		return classFieldsResult{fields: enumFields}, nil
	}
	return classFieldsResult{}, nil
}
