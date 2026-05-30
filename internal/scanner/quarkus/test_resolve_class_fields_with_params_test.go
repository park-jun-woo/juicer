//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveClassFieldsWithParams 테스트
package quarkus

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"path/filepath"
	"testing"
)

func TestResolveClassFieldsWithParams(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "UserDto.java")
	writeFile(t, dir, "UserDto.java", `class UserDto { private String name; private int age; }`)
	cache := map[string][]scanner.Field{}
	r, err := resolveClassFieldsWithParams(p, "UserDto", dir, cache)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.fields) != 2 {
		t.Fatalf("got %+v", r.fields)
	}
}
