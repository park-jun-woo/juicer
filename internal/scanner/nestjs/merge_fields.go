//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 부모와 자식 DTO 필드를 병합한다
package nestjs

// mergeFields combines parent and child fields. Child fields override parent fields with the same name.
func mergeFields(parent, child []dtoField) []dtoField {
	if len(parent) == 0 {
		return child
	}
	childNames := make(map[string]struct{}, len(child))
	for _, f := range child {
		childNames[f.name] = struct{}{}
	}
	var result []dtoField
	for _, f := range parent {
		if _, overridden := childNames[f.name]; !overridden {
			result = append(result, f)
		}
	}
	result = append(result, child...)
	return result
}
