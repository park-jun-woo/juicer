//ff:func feature=scan type=test control=sequence topic=express
//ff:what E2E 스캔 테스트: validateRequest + z.object() → Request.Query 추출
package express

import "testing"

func TestScan_ZodQuery(t *testing.T) {
	dir := t.TempDir()

	appSrc := `
import { Router } from 'express';
import { validateRequest } from 'zod-express-middleware';
import { z } from 'zod';

const PaginationSchema = z.object({
  page: z.number().int().min(1),
  limit: z.number().int().min(1).max(100),
});

const router = Router();
router.get('/users', validateRequest({ query: PaginationSchema }), listUsers);
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
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	if len(ep.Request.Query) != 2 {
		t.Fatalf("expected 2 query params, got %d", len(ep.Request.Query))
	}
	if ep.Request.Query[0].Name != "page" {
		t.Errorf("expected query param page, got %s", ep.Request.Query[0].Name)
	}
	if ep.Request.Query[1].Name != "limit" {
		t.Errorf("expected query param limit, got %s", ep.Request.Query[1].Name)
	}
}
