//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what E2E 스캔 테스트: 체인 라우트 .route().post(validateRequest({body: z.object()})).get() → POST에 zod requestBody 생성, GET 정상
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestScan_ChainZodBody(t *testing.T) {
	dir := t.TempDir()

	appSrc := `
import { Router } from 'express';
import { validateRequest } from 'zod-express-middleware';
import { z } from 'zod';

const CreateItemSchema = z.object({
  name: z.string().min(1).max(100),
  qty: z.number().int().min(0),
});

const router = Router();
router.route('/items')
  .post(auth('manageItems'), validateRequest({ body: CreateItemSchema }), createItem)
  .get(auth('getItems'), listItems);
export default router;
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	var post, get *scanner.Endpoint
	for i := range result.Endpoints {
		ep := &result.Endpoints[i]
		switch ep.Method {
		case "POST":
			post = ep
		case "GET":
			get = ep
		}
	}

	if post == nil {
		t.Fatal("expected POST endpoint")
	}
	if post.Path != "/items" {
		t.Fatalf("POST path: want /items, got %s", post.Path)
	}
	if post.Request == nil || post.Request.Body == nil {
		t.Fatal("expected POST request body from chain zod validator")
	}
	if post.Request.Body.TypeName != "CreateItemSchema" {
		t.Errorf("POST body typeName: want CreateItemSchema, got %s", post.Request.Body.TypeName)
	}
	if len(post.Request.Body.Fields) != 2 {
		t.Fatalf("expected 2 body fields, got %d", len(post.Request.Body.Fields))
	}
	f0 := post.Request.Body.Fields[0]
	if f0.Name != "name" || f0.Type != "string" {
		t.Errorf("field 0: want name/string, got %s/%s", f0.Name, f0.Type)
	}
	f1 := post.Request.Body.Fields[1]
	if f1.Name != "qty" || f1.Type != "integer" {
		t.Errorf("field 1: want qty/integer, got %s/%s", f1.Name, f1.Type)
	}

	if get == nil {
		t.Fatal("expected GET endpoint")
	}
	if get.Path != "/items" {
		t.Fatalf("GET path: want /items, got %s", get.Path)
	}
	if get.Request != nil && get.Request.Body != nil {
		t.Error("GET should not have a request body")
	}
}
