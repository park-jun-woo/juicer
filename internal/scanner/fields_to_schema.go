//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what Field 목록을 OpenAPI schema object로 변환한다
package scanner

import (
	"sort"
)

func fieldsToSchema(fields []Field) map[string]any {
	props := map[string]any{}
	var required []string

	for _, f := range fields {
		propName := f.JSON
		if propName == "" {
			propName = f.Name
		}

		prop := fieldToProperty(f)
		props[propName] = prop

		if isRequired(f) {
			required = append(required, propName)
		}
	}

	schema := map[string]any{
		"type":       "object",
		"properties": props,
	}
	if len(required) > 0 {
		sort.Strings(required)
		schema["required"] = required
	}
	return schema
}

