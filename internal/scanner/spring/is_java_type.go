//ff:func feature=scan type=extract control=selection topic=spring
//ff:what 문자열이 Java 기본/내장 타입인지 확인한다
package spring

import "strings"

func isJavaType(t string) bool {
	switch t {
	case "String", "int", "Integer", "long", "Long", "float", "Float",
		"double", "Double", "boolean", "Boolean", "BigDecimal",
		"LocalDateTime", "ZonedDateTime", "Instant", "OffsetDateTime",
		"LocalDate", "UUID", "MultipartFile", "byte[]":
		return true
	}
	if isCollectionType(t) || strings.HasPrefix(t, "Map<") || strings.HasSuffix(t, "[]") {
		return true
	}
	return false
}
