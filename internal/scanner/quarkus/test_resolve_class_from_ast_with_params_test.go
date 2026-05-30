//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveClassFromASTWithParams 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestResolveClassFromASTWithParams(t *testing.T) {
	src := []byte(`class UserDto { private String name; }`)
	root, _ := parseJava(src)
	cache := map[string][]scanner.Field{}
	fields, _ := resolveClassFromASTWithParams(root, src, "UserDto", "/abs/UserDto.java", "/abs", cache)
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("got %+v", fields)
	}
}
