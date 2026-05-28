//ff:func feature=scan type=test control=sequence topic=hono
//ff:what z.object() AST → scanner.Field 변환 테스트
package hono

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner/zod"
)

func TestParseZodSchema(t *testing.T) {
	src := []byte(`
const schema = z.object({
  name: z.string().min(1),
  email: z.string().email(),
  age: z.number().int().min(0).optional()
})
`)
	fi := mustParse(t, src)
	schemas := zod.CollectSchemas(fi.Root, fi.Src)
	node, ok := schemas["schema"]
	if !ok {
		t.Fatal("schema not found")
	}
	fields := zod.ParseSchema(node, fi.Src)
	if len(fields) != 3 {
		t.Fatalf("expected 3 fields, got %d", len(fields))
	}
	// name
	if fields[0].Name != "name" {
		t.Errorf("expected field name, got %s", fields[0].Name)
	}
	if fields[0].Type != "string" {
		t.Errorf("expected type string, got %s", fields[0].Type)
	}
	if fields[0].MinLength == nil || *fields[0].MinLength != 1 {
		t.Errorf("expected minLength 1")
	}
	// email
	if fields[1].Name != "email" {
		t.Errorf("expected field email, got %s", fields[1].Name)
	}
	if fields[1].Validate != "email" {
		t.Errorf("expected validate email, got %s", fields[1].Validate)
	}
	// age
	if fields[2].Name != "age" {
		t.Errorf("expected field age, got %s", fields[2].Name)
	}
	if fields[2].Type != "integer" {
		t.Errorf("expected type integer, got %s", fields[2].Type)
	}
	if !fields[2].Nullable {
		t.Errorf("expected nullable true for optional field")
	}
	if fields[2].Minimum == nil || *fields[2].Minimum != 0 {
		t.Errorf("expected minimum 0")
	}
}
