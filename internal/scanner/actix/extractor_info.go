//ff:type feature=scan type=model topic=actix
//ff:what actix extractor 정보(kind/typeName/rawType)
package actix

type extractorInfo struct {
	kind     string // "path", "json", "query", "form"
	typeName string // inner type name (e.g., "i64", "CreateUserRequest")
	rawType  string // full type text
}
