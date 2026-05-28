//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what PHP 타입을 OpenAPI 타입으로 변환한다
package laravel

// phpTypeToOpenAPI converts a PHP type name to an OpenAPI type string.
func phpTypeToOpenAPI(phpType string) string {
	switch phpType {
	case "int", "integer":
		return "integer"
	case "float", "double":
		return "number"
	case "string":
		return "string"
	case "bool", "boolean":
		return "boolean"
	case "array":
		return "array"
	default:
		return ""
	}
}
