//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what E2E 스캔 테스트: 선두 검증자 체인 .route().post(validateRequest({body: z.object()}), h).get(h) → POST에 zod requestBody 생성(미들웨어 선행 없음)
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestScan_ChainZodLeadingValidator(t *testing.T) {
	dir := t.TempDir()

	appSrc := `
import { Router } from 'express';
import { validateRequest } from 'zod-express-middleware';
import { z } from 'zod';

const router = Router();
router.route('/items')
  .post(validateRequest({ body: z.object({ name: z.string() }) }), createItem)
  .get(listItems);
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
		t.Fatal("expected POST request body from leading chain zod validator")
	}
	if len(post.Request.Body.Fields) != 1 {
		t.Fatalf("expected 1 body field, got %d", len(post.Request.Body.Fields))
	}
	f0 := post.Request.Body.Fields[0]
	if f0.Name != "name" || f0.Type != "string" {
		t.Errorf("field 0: want name/string, got %s/%s", f0.Name, f0.Type)
	}

	if get == nil {
		t.Fatal("expected GET endpoint")
	}
	if get.Request != nil && get.Request.Body != nil {
		t.Error("GET should not have a request body")
	}
}
