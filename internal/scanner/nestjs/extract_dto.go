//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what DTO 소스 파일에서 지정 클래스의 프로퍼티를 추출한다
package nestjs

import (
	"os"
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// extractDTO parses a DTO source file and extracts the class with the given name.
// imports and projectRoot are used for resolving parent DTO references from extends clauses.
// cache prevents infinite recursion for circular references.
func extractDTO(filePath, className string, imports map[string]string, projectRoot string, cache map[string][]scanner.Field) ([]scanner.Field, error) {
	if cached, ok := cache[className]; ok {
		return cached, nil
	}
	src, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	root, err := parseTypeScript(src)
	if err != nil {
		return nil, err
	}
	fileImports := extractImports(root, src)
	merged := mergeImports(imports, fileImports)

	classes := findAllByType(root, "class_declaration")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "type_identifier")
		if nameNode == nil {
			continue
		}
		if nodeText(nameNode, src) != className {
			continue
		}
		cache[className] = nil
		directFields := extractClassProperties(cls, src)
		resolveEnumTypeNames(directFields, root, src, filePath, merged, projectRoot)
		parentFields := resolveDTOExtends(cls, src, filePath, merged, projectRoot, cache)
		combined := mergeFields(parentFields, directFields)
		result := dtoFieldsToScannerFields(combined)
		cache[className] = result
		return result, nil
	}
	// class not found: follow barrel re-exports when the file is index.ts
	if strings.HasSuffix(filePath, "index.ts") {
		realPath := resolveBarrelExport(filePath, className)
		if realPath != "" {
			return extractDTO(realPath, className, imports, projectRoot, cache)
		}
	}
	return nil, nil
}
