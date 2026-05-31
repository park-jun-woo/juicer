//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what schemaRegistry.process 파일 해석→enum 스키마 등록 및 부재/중복 가드 테스트
package nestjs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestSchemaRegistryProcess(t *testing.T) {
	dir := t.TempDir()
	enumFile := filepath.Join(dir, "status.enum.ts")
	if err := os.WriteFile(enumFile, []byte(`export enum Status { Active = 'active', Done = 'done' }`), 0o644); err != nil {
		t.Fatal(err)
	}
	r := &schemaRegistry{
		cache:     map[string][]scanner.Field{},
		schemas:   map[string]any{},
		processed: map[string]bool{},
		topLevel:  map[string]bool{},
	}
	j := schemaJob{
		typeName:    "Status",
		imports:     map[string]string{"Status": "./status.enum"},
		referrer:    filepath.Join(dir, "c.ts"),
		projectRoot: dir,
	}
	r.process(j)
	if _, ok := r.schemas["Status"]; !ok {
		t.Errorf("Status enum schema not registered: %v", r.schemas)
	}
	if !r.processed["Status"] {
		t.Error("not marked processed")
	}

	// duplicate process is a no-op (already processed)
	before := len(r.schemas)
	r.process(j)
	if len(r.schemas) != before {
		t.Error("duplicate process should be no-op")
	}

	// unresolvable type -> no schema added
	r.process(schemaJob{typeName: "Ghost", imports: map[string]string{}, referrer: j.referrer, projectRoot: dir})
	if _, ok := r.schemas["Ghost"]; ok {
		t.Error("ghost should not register")
	}
}
