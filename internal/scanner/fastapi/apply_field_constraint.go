//ff:func feature=scan type=extract control=selection topic=fastapi
//ff:what 키워드 이름에 따라 pydanticField에 제약 조건 값을 설정한다
package fastapi

// applyFieldConstraint sets a constraint value on pydanticField based on the keyword name.
func applyFieldConstraint(key, valStr string, f *pydanticField) {
	switch key {
	case "default", "default_factory":
		f.hasDefault = true
	case "ge":
		v := parseIntDefault(valStr, -1)
		if v >= 0 {
			f.ge = &v
		}
	case "le":
		v := parseIntDefault(valStr, -1)
		if v >= 0 {
			f.le = &v
		}
	case "min_length":
		v := parseIntDefault(valStr, -1)
		if v >= 0 {
			f.minLength = &v
		}
	case "max_length":
		v := parseIntDefault(valStr, -1)
		if v >= 0 {
			f.maxLength = &v
		}
	}
}
