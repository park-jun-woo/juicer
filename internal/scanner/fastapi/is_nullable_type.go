//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what Python 타입이 nullable인지 확인한다
package fastapi

import "strings"

// isNullableType checks if a Python type is nullable.
func isNullableType(typeName string) bool {
	if strings.HasPrefix(typeName, "Optional[") {
		return true
	}
	if strings.Contains(typeName, "| None") || strings.Contains(typeName, "None |") {
		return true
	}
	if strings.HasPrefix(typeName, "Union[") && strings.Contains(typeName, "None") {
		return true
	}
	return false
}
