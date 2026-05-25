//ff:func feature=scan type=extract control=selection
//ff:what Go 타입과 validate 태그에서 OpenAPI format을 추론한다
package scanner

import (
	"strings"
)

func goTypeFormat(goType string, f Field) string {
	switch goType {
	case "int64", "uint64":
		return "int64"
	case "int32", "uint32":
		return "int32"
	case "float64":
		return "double"
	case "float32":
		return "float"
	case "time.Time":
		return "date-time"
	}

	// validate 태그에서 format 추론
	if f.Validate != "" {
		if strings.Contains(f.Validate, "email") {
			return "email"
		}
		if strings.Contains(f.Validate, "url") || strings.Contains(f.Validate, "uri") {
			return "uri"
		}
	}

	return ""
}

