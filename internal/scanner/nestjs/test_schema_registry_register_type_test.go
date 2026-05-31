//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what registerType enum 스키마 등록 및 top-level 미등록(직접 호출) 테스트
package nestjs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestSchemaRegistryRegisterType(t *testing.T) {
	dir := t.TempDir()
	src := []byte(`export enum Status { A = 'a', B = 'b' }`)
	absPath := filepath.Join(dir, "status.ts")
	if err := os.WriteFile(absPath, src, 0o644); err != nil {
		t.Fatal(err)
	}
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	j := schemaJob{typeName: "Status", imports: map[string]string{}, referrer: absPath, projectRoot: dir}

	// non-top-level enum -> registered
	r := &schemaRegistry{cache: map[string][]scanner.Field{}, schemas: map[string]any{}, processed: map[string]bool{}, topLevel: map[string]bool{}}
	r.registerType(j, absPath, root, src, map[string]string{})
	if _, ok := r.schemas["Status"]; !ok {
		t.Errorf("enum should register: %v", r.schemas)
	}

	// top-level enum -> NOT registered
	r2 := &schemaRegistry{cache: map[string][]scanner.Field{}, schemas: map[string]any{}, processed: map[string]bool{}, topLevel: map[string]bool{"Status": true}}
	r2.registerType(j, absPath, root, src, map[string]string{})
	if _, ok := r2.schemas["Status"]; ok {
		t.Error("top-level enum must not register")
	}
}
