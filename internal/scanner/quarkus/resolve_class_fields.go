//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what DTO 클래스를 추적하여 필드를 해석한다
package quarkus

import "github.com/park-jun-woo/codistill/internal/scanner"

func resolveClassFields(filePath, className string, projectRoot string, cache map[string][]scanner.Field) ([]scanner.Field, error) {
	r, err := resolveClassFieldsWithParams(filePath, className, projectRoot, cache)
	if err != nil {
		return nil, err
	}
	return r.fields, nil
}
