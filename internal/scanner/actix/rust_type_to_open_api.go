//ff:func feature=scan type=convert control=selection topic=actix
//ff:what Rust 타입을 OpenAPI 타입으로 변환한다
package actix

import "strings"

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

	return rustComplexTypeToOpenAPI(rtype)
}
