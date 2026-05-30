//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveParentFields 테스트
package spring

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"path/filepath"
	"testing"
)

func TestResolveParentFields(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Base.java", `class Base { private Long id; }`)
	childSrc := `class Child extends Base { private String name; }`
	writeFile(t, dir, "Child.java", childSrc)
	root, _ := parseJava([]byte(childSrc))
	cls := findAllByType(root, "class_declaration")[0]
	fields := resolveParentFields(cls, []byte(childSrc), filepath.Join(dir, "Child.java"), dir, map[string]string{}, map[string][]scanner.Field{})
	if len(fields) != 1 || fields[0].Name != "id" {
		t.Fatalf("got %+v", fields)
	}
}
