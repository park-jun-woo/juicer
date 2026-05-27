//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 부모 DTO 이름으로 필드를 해석하여 dtoField 목록으로 반환한다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// resolveParentDTO resolves a parent DTO by name and returns its fields as dtoField slice.
func resolveParentDTO(parentName, referrerFile string, imports map[string]string, projectRoot string, cache map[string][]scanner.Field) []dtoField {
	fields := resolveParentDTOFields(parentName, referrerFile, imports, projectRoot, cache)
	return scannerFieldsToDTOFields(fields)
}
