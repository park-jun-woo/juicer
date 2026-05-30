//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveClassFields_Round5 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"os"
	"path/filepath"
	"testing"
)

func TestResolveClassFields_Round5(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "UserDto.java"), []byte(`package x;
class UserDto { public String name; }
`), 0o644); err != nil {
		t.Fatal(err)
	}
	fields, err := resolveClassFields(filepath.Join(dir, "UserDto.java"), "UserDto", dir, map[string][]scanner.Field{})
	if err != nil {
		t.Fatal(err)
	}
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("fields: %+v", fields)
	}
}
