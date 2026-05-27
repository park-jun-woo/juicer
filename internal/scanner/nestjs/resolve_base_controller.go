//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what extends 절의 팩토리 호출을 추적하여 부모 클래스 메서드를 추출한다
package nestjs

import (
	"os"
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"
)

// resolveBaseController detects a factory-call extends clause on a controller
// class, resolves the import to an absolute file, parses that file, finds the
// factory function, and returns the inherited endpoints extracted from the
// inner class declared inside the factory.
//
// absFile is the absolute path of the file containing cls.
// imports maps type names to their relative import paths.
// file is the relative path used for endpoint info.
func resolveBaseController(cls *sitter.Node, src []byte, absFile string, imports map[string]string, file string) []endpointInfo {
	funcName := extendsFactoryName(cls, src)
	if funcName == "" {
		return nil
	}
	relImport, ok := imports[funcName]
	if !ok {
		return nil
	}
	referrerDir := filepath.Dir(absFile)
	resolved := resolveImportPath(referrerDir, relImport)
	if resolved == "" {
		return nil
	}
	baseSrc, err := os.ReadFile(resolved)
	if err != nil {
		return nil
	}
	baseRoot, err := parseTypeScript(baseSrc)
	if err != nil {
		return nil
	}
	innerCls := findFactoryInnerClass(baseRoot, baseSrc, funcName)
	if innerCls == nil {
		return nil
	}
	relBase, _ := filepath.Rel(filepath.Dir(absFile), resolved)
	if relBase == "" {
		relBase = resolved
	}
	endpoints := extractMethods(innerCls, baseSrc, relBase)

	// 제네릭 치환: 자식 extends 절의 타입 인자 → 내부 클래스의 타입 파라미터 매핑
	typeArgs := extractTypeArgs(cls, src)
	typeParams := extractTypeParams(innerCls, baseSrc)
	if len(typeArgs) > 0 && len(typeParams) > 0 {
		n := len(typeParams)
		if len(typeArgs) < n {
			n = len(typeArgs)
		}
		typeMap := make(map[string]string, n)
		for i := 0; i < n; i++ {
			typeMap[typeParams[i]] = typeArgs[i]
		}
		substituteTypes(endpoints, typeMap)
	}

	return endpoints
}
