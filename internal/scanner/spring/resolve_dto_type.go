//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what DTO 타입을 해석하여 필드를 반환한다
package spring

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveDTOType(dr dtoRequest, projectRoot string, cache map[string][]scanner.Field) []scanner.Field {
	if cached, ok := cache[dr.typeName]; ok {
		return cached
	}

	lookupName := stripGeneric(dr.typeName)
	genericArg := extractGenericArgs(dr.typeName)

	filePath := resolveSameFileClass(dr.referrer, lookupName, projectRoot)
	if filePath == "" {
		if fqcn, ok := dr.imports[lookupName]; ok {
			filePath = resolveImportPath(projectRoot, fqcn)
		}
	}
	if filePath == "" {
		filePath = resolveSamePackageClass(dr.referrer, lookupName)
	}
	if filePath == "" {
		return nil
	}
	result, err := resolveClassFieldsWithParams(filePath, lookupName, projectRoot, cache)
	if err != nil {
		return nil
	}
	fields := result.fields
	if genericArg != "" && len(result.typeParams) > 0 {
		typeArgs := strings.Split(genericArg, ",")
		for i := range typeArgs {
			typeArgs[i] = strings.TrimSpace(typeArgs[i])
		}
		fields = substituteTypeParams(fields, result.typeParams, typeArgs)
	}
	cache[dr.typeName] = fields
	return fields
}
