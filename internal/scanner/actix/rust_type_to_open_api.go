//ff:func feature=scan type=convert control=selection topic=actix
//ff:what Rust 타입을 OpenAPI 타입으로 변환한다
package actix

import "strings"

type openAPIType struct {
	Type   string
	Format string
	Items  string
}

func rustTypeToOpenAPI(rtype string) openAPIType {
	rtype = strings.TrimSpace(rtype)

	switch rtype {
	case "String", "&str", "str":
		return openAPIType{Type: "string"}
	case "i8", "i16", "i32":
		return openAPIType{Type: "integer", Format: "int32"}
	case "i64":
		return openAPIType{Type: "integer", Format: "int64"}
	case "u8", "u16", "u32":
		return openAPIType{Type: "integer", Format: "int32"}
	case "u64":
		return openAPIType{Type: "integer", Format: "int64"}
	case "f32":
		return openAPIType{Type: "number", Format: "float"}
	case "f64":
		return openAPIType{Type: "number", Format: "double"}
	case "bool":
		return openAPIType{Type: "boolean"}
	case "Uuid":
		return openAPIType{Type: "string", Format: "uuid"}
	case "NaiveDateTime", "DateTime", "DateTime<Utc>":
		return openAPIType{Type: "string", Format: "date-time"}
	case "NaiveDate":
		return openAPIType{Type: "string", Format: "date"}
	}

	if strings.HasPrefix(rtype, "Vec<") {
		inner := extractGenericInner(rtype)
		return openAPIType{Type: "array", Items: inner}
	}

	if strings.HasPrefix(rtype, "HashMap<") || strings.HasPrefix(rtype, "BTreeMap<") {
		return openAPIType{Type: "object"}
	}

	if strings.HasPrefix(rtype, "Option<") {
		inner := extractGenericInner(rtype)
		return rustTypeToOpenAPI(inner)
	}

	return openAPIType{Type: "object"}
}

func extractGenericInner(t string) string {
	idx := strings.Index(t, "<")
	if idx < 0 {
		return t
	}
	inner := t[idx+1:]
	if len(inner) > 0 && inner[len(inner)-1] == '>' {
		inner = inner[:len(inner)-1]
	}
	return strings.TrimSpace(inner)
}

func isOptionType(t string) bool {
	return strings.HasPrefix(t, "Option<")
}
