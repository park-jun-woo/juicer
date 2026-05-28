//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what E2E 테스트: Express 스캔 결과에서 인증/역할이 올바르게 매핑되는지 검증한다
package express

import (
	"reflect"
	"testing"
)

func TestExtractAuthMiddleware_E2E(t *testing.T) {
	dir := t.TempDir()

	routeSrc := `
import { Router } from 'express';
import { authenticate } from '../middleware/authenticate';
import { authorize } from '../middleware/authorize';

const router = Router();

router.get('/health', healthCheck);
router.get('/profile', authenticate, getProfile);
router.delete('/users/:id', authenticate, authorize('admin'), deleteUser);
router.put('/settings', authenticate, authorize('admin', 'manager'), updateSettings);

export default router;
`
	writeFile(t, dir, "routes/api.ts", routeSrc)

	appSrc := `
import express from "express";
import apiRouter from "./routes/api";
const app = express();
app.use("/api", apiRouter);
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	type expected struct {
		method    string
		path      string
		authLevel string
		roles     []string
	}
	expectations := []expected{
		{"GET", "/api/health", "public", nil},
		{"GET", "/api/profile", "auth_required", nil},
		{"DELETE", "/api/users/{id}", "auth_required", []string{"admin"}},
		{"PUT", "/api/settings", "auth_required", []string{"admin", "manager"}},
	}

	epMap := map[string]struct {
		authLevel string
		roles     []string
	}{}
	for _, ep := range result.Endpoints {
		epMap[ep.Method+" "+ep.Path] = struct {
			authLevel string
			roles     []string
		}{ep.AuthLevel, ep.Roles}
	}

	for _, exp := range expectations {
		key := exp.method + " " + exp.path
		got, ok := epMap[key]
		if !ok {
			t.Errorf("missing endpoint %s", key)
			continue
		}
		if got.authLevel != exp.authLevel {
			t.Errorf("%s AuthLevel: want %s, got %s", key, exp.authLevel, got.authLevel)
		}
		if !reflect.DeepEqual(got.roles, exp.roles) {
			t.Errorf("%s Roles: want %v, got %v", key, exp.roles, got.roles)
		}
	}
}
