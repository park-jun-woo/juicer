//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what applyNamedRef DTO/배열 $ref 설정, 인라인 enum/빌트인 무시 테스트
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyNamedRef(t *testing.T) {
	// scalar DTO -> Ref set
	var f scanner.Field
	applyNamedRef(&f, "UserDto")
	if f.Ref != "UserDto" || f.Type == "array" {
		t.Errorf("scalar: %+v", f)
	}
	// array DTO -> Ref + Type array
	var a scanner.Field
	applyNamedRef(&a, "UserDto[]")
	if a.Ref != "UserDto" || a.Type != "array" {
		t.Errorf("array: %+v", a)
	}
	// builtin -> untouched
	var b scanner.Field
	applyNamedRef(&b, "string")
	if b.Ref != "" {
		t.Errorf("builtin: %+v", b)
	}
	// inline enum -> untouched
	e := scanner.Field{Enum: []string{"A"}}
	applyNamedRef(&e, "SomeDto")
	if e.Ref != "" {
		t.Errorf("inline enum: %+v", e)
	}
}
