//ff:func feature=scan type=test control=sequence topic=express
//ff:what E2E 스캔 테스트: validateRequest + z.object() → Request.Body.Fields 추출
package express

import "testing"

func TestScan_ZodBody(t *testing.T) {
	dir := t.TempDir()

	appSrc := `
import { Router } from 'express';
import { validateRequest } from 'zod-express-middleware';
import { z } from 'zod';

const CreateUserSchema = z.object({
  name: z.string().min(1).max(100),
  email: z.string().email(),
  age: z.number().int().min(0).optional(),
});

const router = Router();
router.post('/users', validateRequest({ body: CreateUserSchema }), createUser);
export default router;
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if ep.Method != "POST" || ep.Path != "/users" {
		t.Fatalf("unexpected endpoint: %s %s", ep.Method, ep.Path)
	}
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	if ep.Request.Body == nil {
		t.Fatal("expected body")
	}
	if ep.Request.Body.TypeName != "CreateUserSchema" {
		t.Errorf("expected typeName CreateUserSchema, got %s", ep.Request.Body.TypeName)
	}
	if len(ep.Request.Body.Fields) != 3 {
		t.Fatalf("expected 3 body fields, got %d", len(ep.Request.Body.Fields))
	}
	f0 := ep.Request.Body.Fields[0]
	if f0.Name != "name" || f0.Type != "string" {
		t.Errorf("field 0: expected name/string, got %s/%s", f0.Name, f0.Type)
	}
	if f0.MinLength == nil || *f0.MinLength != 1 {
		t.Error("expected minLength 1")
	}
	if f0.MaxLength == nil || *f0.MaxLength != 100 {
		t.Error("expected maxLength 100")
	}
	f1 := ep.Request.Body.Fields[1]
	if f1.Name != "email" || f1.Type != "string" {
		t.Errorf("field 1: expected email/string, got %s/%s", f1.Name, f1.Type)
	}
	if f1.Validate != "email" {
		t.Errorf("expected validate email, got %s", f1.Validate)
	}
	f2 := ep.Request.Body.Fields[2]
	if f2.Name != "age" || f2.Type != "integer" {
		t.Errorf("field 2: expected age/integer, got %s/%s", f2.Name, f2.Type)
	}
	if !f2.Nullable {
		t.Error("expected nullable for optional field")
	}
	if f2.Minimum == nil || *f2.Minimum != 0 {
		t.Error("expected minimum 0")
	}
}
