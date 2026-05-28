//ff:func feature=scan type=test control=sequence topic=express
//ff:what forEach 배열 라우터 마운트 테스트: 배열 리터럴에서 path/route 추출
package express

import "testing"

func TestExtractArrayRouteMounts(t *testing.T) {
	src := []byte(`
import authRoute from './routes/auth';
import userRoute from './routes/users';
import express from 'express';

const router = express.Router();

const defaultRoutes: IRoute[] = [
  { path: '/auth', route: authRoute },
  { path: '/users', route: userRoute },
];

defaultRoutes.forEach((r) => {
  router.use(r.path, r.route);
});
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi, nil)
	imports := map[string]string{
		"authRoute": "/abs/routes/auth.ts",
		"userRoute": "/abs/routes/users.ts",
	}

	entries := extractArrayRouteMounts(fi, routers, imports, "test.ts")
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}

	if entries[0].prefix != "/auth" {
		t.Errorf("entries[0].prefix = %q, want /auth", entries[0].prefix)
	}
	if entries[0].varName != "authRoute" {
		t.Errorf("entries[0].varName = %q, want authRoute", entries[0].varName)
	}
	if entries[0].filePath != "/abs/routes/auth.ts" {
		t.Errorf("entries[0].filePath = %q, want /abs/routes/auth.ts", entries[0].filePath)
	}

	if entries[1].prefix != "/users" {
		t.Errorf("entries[1].prefix = %q, want /users", entries[1].prefix)
	}
	if entries[1].varName != "userRoute" {
		t.Errorf("entries[1].varName = %q, want userRoute", entries[1].varName)
	}
	if entries[1].filePath != "/abs/routes/users.ts" {
		t.Errorf("entries[1].filePath = %q, want /abs/routes/users.ts", entries[1].filePath)
	}
}
