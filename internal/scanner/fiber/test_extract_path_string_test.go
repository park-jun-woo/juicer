//ff:func feature=scan type=test topic=fiber control=sequence
//ff:what extractPathString 리터럴/문자열연결 경로 추출 및 비문자열 false 테스트
package fiber

import (
	"go/parser"
	"testing"
)

func TestExtractPathString(t *testing.T) {
	lit, _ := parser.ParseExpr(`"/users"`)
	if s, ok := extractPathString(lit); !ok || s != "/users" {
		t.Errorf("literal: (%q,%v)", s, ok)
	}
	bin, _ := parser.ParseExpr(`"/users/" + "list"`)
	if s, ok := extractPathString(bin); !ok || s == "" {
		t.Errorf("concat: (%q,%v)", s, ok)
	}
	num, _ := parser.ParseExpr(`42`)
	if _, ok := extractPathString(num); ok {
		t.Error("non-string should be false")
	}
}
