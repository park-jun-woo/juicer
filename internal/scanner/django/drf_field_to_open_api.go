//ff:func feature=scan type=convert control=selection topic=django
//ff:what DRF Serializer 필드 타입을 OpenAPI 타입으로 변환한다
package django

// drfFieldToOpenAPI converts a DRF serializer field type to OpenAPI type.
func drfFieldToOpenAPI(fieldType string) openAPIType {
	switch fieldType {
	case "CharField", "SlugField", "SlugRelatedField":
		return openAPIType{Type: "string"}
	case "EmailField":
		return openAPIType{Type: "string", Format: "email"}
	case "URLField":
		return openAPIType{Type: "string", Format: "uri"}
	case "UUIDField":
		return openAPIType{Type: "string", Format: "uuid"}
	case "IntegerField", "PrimaryKeyRelatedField":
		return openAPIType{Type: "integer"}
	case "FloatField":
		return openAPIType{Type: "number", Format: "float"}
	case "DecimalField":
		return openAPIType{Type: "number"}
	case "BooleanField", "NullBooleanField":
		return openAPIType{Type: "boolean"}
	case "DateTimeField":
		return openAPIType{Type: "string", Format: "date-time"}
	case "DateField":
		return openAPIType{Type: "string", Format: "date"}
	case "TimeField":
		return openAPIType{Type: "string", Format: "time"}
	case "FileField", "ImageField":
		return openAPIType{Type: "string", Format: "binary"}
	case "ListField", "ListSerializer", "ManyRelatedField":
		return openAPIType{Type: "array"}
	case "DictField", "JSONField", "HStoreField":
		return openAPIType{Type: "object"}
	case "ChoiceField":
		return openAPIType{Type: "string"}
	case "IPAddressField", "GenericIPAddressField":
		return openAPIType{Type: "string", Format: "ipv4"}
	case "DurationField":
		return openAPIType{Type: "string"}
	case "SerializerMethodField":
		return openAPIType{Type: "string"}
	default:
		return openAPIType{Type: "string"}
	}
}
