//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 상위 클래스의 필드를 해석한다
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveParentFields(cls *sitter.Node, src []byte, referrerPath, projectRoot string, imports map[string]string, cache map[string][]scanner.Field) []scanner.Field {
	superclass := findChildByType(cls, "superclass")
	if superclass == nil {
		return nil
	}
	parentName := extractSuperclassName(superclass, src)
	if parentName == "" {
		return nil
	}
	parentPath := resolveSameFileClass(referrerPath, parentName, projectRoot)
	if parentPath == "" {
		parentPath = resolveSamePackageClass(referrerPath, parentName)
	}
	if parentPath == "" {
		if fqcn, ok := imports[parentName]; ok {
			parentPath = resolveImportPath(projectRoot, fqcn)
		}
	}
	if parentPath == "" {
		return nil
	}
	fields, _ := resolveClassFields(parentPath, parentName, projectRoot, cache)
	return fields
}
