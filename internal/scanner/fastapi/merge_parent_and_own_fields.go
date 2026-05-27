//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 자식 필드가 부모 필드를 override하여 병합한다
package fastapi

// mergeParentAndOwnFields merges parent fields with own fields.
// If a field name exists in both, the own (child) field takes precedence.
func mergeParentAndOwnFields(parent, own []pydanticField) []pydanticField {
	nameSet := make(map[string]bool)
	for _, f := range own {
		nameSet[f.name] = true
	}
	var merged []pydanticField
	for _, f := range parent {
		if !nameSet[f.name] {
			merged = append(merged, f)
		}
	}
	merged = append(merged, own...)
	return merged
}
