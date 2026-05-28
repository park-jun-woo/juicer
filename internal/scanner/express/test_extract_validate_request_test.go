//ff:func feature=scan type=test control=sequence topic=express
//ff:what validateRequest({body: X, query: Y}) 추출 테스트
package express

import "testing"

func TestExtractValidateRequest_BodyAndQuery(t *testing.T) {
	src := []byte(`
import { Router } from 'express';
import { validateRequest } from 'zod-express-middleware';
const router = Router();
router.post('/users', validateRequest({ body: CreateUserSchema, query: PaginationSchema }), createUser);
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi, collectExpressRouterImports(fi))
	routes := extractRoutes(fi, routers)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if len(r.ZodValidators) != 2 {
		t.Fatalf("expected 2 validators, got %d", len(r.ZodValidators))
	}
	body := r.ZodValidators[0]
	if body.Target != "json" {
		t.Errorf("expected target json, got %s", body.Target)
	}
	if body.SchemaName != "CreateUserSchema" {
		t.Errorf("expected schema CreateUserSchema, got %s", body.SchemaName)
	}
	query := r.ZodValidators[1]
	if query.Target != "query" {
		t.Errorf("expected target query, got %s", query.Target)
	}
	if query.SchemaName != "PaginationSchema" {
		t.Errorf("expected schema PaginationSchema, got %s", query.SchemaName)
	}
}
