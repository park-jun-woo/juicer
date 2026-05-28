//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 상수 참조(예: AppConstants.DEFAULT_PAGE_NUMBER)를 실제 값으로 해석한다
package spring

import (
	"os"
	"strings"
)

func resolveConstantValue(constRef string, imports map[string]string, referrerPath, projectRoot string) string {
	parts := strings.SplitN(constRef, ".", 2)
	if len(parts) != 2 {
		return constRef
	}
	className := parts[0]
	fieldName := parts[1]

	filePath := ""
	if fqcn, ok := imports[className]; ok {
		filePath = resolveImportPath(projectRoot, fqcn)
	}
	if filePath == "" {
		filePath = resolveSamePackageClass(referrerPath, className)
	}
	if filePath == "" {
		filePath = resolveSameFileClass(referrerPath, className, projectRoot)
	}
	if filePath == "" {
		return constRef
	}

	src, err := os.ReadFile(filePath)
	if err != nil {
		return constRef
	}
	root, err := parseJava(src)
	if err != nil {
		return constRef
	}

	classes := findAllByType(root, "class_declaration")
	for _, cls := range classes {
		nameNode := findChildByType(cls, "identifier")
		if nameNode == nil || nodeText(nameNode, src) != className {
			continue
		}
		val := findStaticFinalField(cls, src, fieldName)
		if val != "" {
			return val
		}
	}
	return constRef
}
