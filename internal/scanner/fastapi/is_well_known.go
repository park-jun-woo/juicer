//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 부모 클래스 이름이 well-known BaseModel 계열인지 확인한다
package fastapi

import "strings"

var wellKnownBases = map[string]bool{
	"BaseModel":    true,
	"SQLModel":     true,
	"BaseSettings": true,
}

// isWellKnown checks if a parent name is a well-known base class.
// Handles both plain names ("BaseModel") and qualified names ("pydantic.BaseModel").
func isWellKnown(name string) bool {
	if wellKnownBases[name] {
		return true
	}
	if idx := strings.LastIndex(name, "."); idx >= 0 {
		return wellKnownBases[name[idx+1:]]
	}
	return false
}
