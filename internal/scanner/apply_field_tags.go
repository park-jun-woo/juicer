//ff:func feature=scan type=extract control=sequence
//ff:what struct 태그에서 json/validate/binding 정보를 Field에 적용한다. json:"-"이면 true 반환.
package scanner

import (
	"reflect"
	"strings"
)

// ApplyFieldTags applies json/validate/binding tags to a Field. Returns true if the field should be excluded.
func ApplyFieldTags(field *Field, tag string) bool {
	st := reflect.StructTag(tag)
	if jsonTag := st.Get("json"); jsonTag != "" {
		if idx := strings.Index(jsonTag, ","); idx >= 0 {
			field.JSON = jsonTag[:idx]
		} else {
			field.JSON = jsonTag
		}
		if field.JSON == "-" {
			return true
		}
	}
	if valTag := st.Get("validate"); valTag != "" {
		field.Validate = valTag
	}
	if bindTag := st.Get("binding"); bindTag != "" && field.Validate == "" {
		field.Validate = bindTag
	}
	return false
}
