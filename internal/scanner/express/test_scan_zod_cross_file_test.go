//ff:func feature=scan type=test control=sequence topic=express
//ff:what E2E 스캔 테스트: 스키마가 별도 파일에 정의된 cross-file import
package express

import "testing"

func TestScan_ZodCrossFile(t *testing.T) {
	dir := t.TempDir()

	schemaSrc := `
import { z } from 'zod';
export const PaginationSchema = z.object({
  page: z.number().int().min(1),
  limit: z.number().int().min(1).max(100),
});
`
	writeFile(t, dir, "schemas/pagination.ts", schemaSrc)

	routeSrc := `
import { Router } from 'express';
import { validateRequest } from 'zod-express-middleware';
import { PaginationSchema } from '../schemas/pagination';
const router = Router();
router.get('/users', validateRequest({ query: PaginationSchema }), listUsers);
export default router;
`
	writeFile(t, dir, "routes/users.ts", routeSrc)

	appSrc := `
import express from 'express';
import usersRouter from './routes/users';
const app = express();
app.use('/api', usersRouter);
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	ep := findEndpoint(result.Endpoints, "GET", "/api/users")
	if ep == nil {
		t.Fatal("missing endpoint GET /api/users")
	}
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	if len(ep.Request.Query) != 2 {
		t.Fatalf("expected 2 query params, got %d", len(ep.Request.Query))
	}
	if ep.Request.Query[0].Name != "page" {
		t.Errorf("expected page, got %s", ep.Request.Query[0].Name)
	}
	if ep.Request.Query[1].Name != "limit" {
		t.Errorf("expected limit, got %s", ep.Request.Query[1].Name)
	}
}
