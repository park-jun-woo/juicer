//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveClassFieldsWithParams 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"path/filepath"
	"testing"
)

func TestResolveClassFieldsWithParams(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "UserDto.java")
	writeFile(t, dir, "UserDto.java", `class UserDto { private String name; }`)
	r, err := resolveClassFieldsWithParams(p, "UserDto", dir, map[string][]scanner.Field{})
	if err != nil || len(r.fields) != 1 {
		t.Fatalf("got %+v err %v", r, err)
	}
}
