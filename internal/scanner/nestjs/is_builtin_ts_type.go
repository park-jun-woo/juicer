//ff:func feature=scan type=convert control=selection topic=nestjs
//ff:what TS/JS 빌트인·프리미티브 타입명인지 검사한다
package nestjs

// isBuiltinTSType reports whether s is a TS/JS builtin or primitive that must
// not be registered as a component schema.
func isBuiltinTSType(s string) bool {
	switch s {
	case "string", "number", "boolean", "Date", "Uuid", "ObjectId",
		"any", "void", "object", "Object", "unknown", "never", "null",
		"undefined", "bigint", "symbol", "Buffer", "Record", "Array",
		"Promise", "Map", "Set", "true", "false":
		return true
	}
	return false
}
