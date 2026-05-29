//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 타입명이 FormRequest 서브클래스(내장 타입 아님 + Request 접미사)인지 보고한다
package laravel

import "strings"

func isFormRequestType(typeName string) bool {
	if typeName == "" || typeName == "Request" ||
		typeName == "int" || typeName == "string" || typeName == "float" ||
		typeName == "bool" || typeName == "array" {
		return false
	}
	return strings.HasSuffix(typeName, "Request")
}
