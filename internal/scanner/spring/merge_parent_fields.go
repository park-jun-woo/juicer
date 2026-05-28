//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 부모 필드와 자체 필드를 병합한다
package spring

import "github.com/park-jun-woo/codistill/internal/scanner"

func mergeParentFields(parent, own []scanner.Field) []scanner.Field {
	if len(parent) == 0 {
		return own
	}
	ownNames := make(map[string]bool)
	for _, f := range own {
		ownNames[f.Name] = true
	}
	var result []scanner.Field
	for _, f := range parent {
		if !ownNames[f.Name] {
			result = append(result, f)
		}
	}
	result = append(result, own...)
	return result
}
