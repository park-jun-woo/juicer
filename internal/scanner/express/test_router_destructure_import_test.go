//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what destructure import { Router } from 'express' + 3단계 체인 prefix 테스트
package express

import "testing"

func TestScan_RouterDestructureImport(t *testing.T) {
	dir := t.TempDir()

	usersRoutes := `
import { Router } from 'express';
const router = Router();
router.get('/', listUsers);
router.post('/', createUser);
export default router;
`
	writeFile(t, dir, "routes/users.ts", usersRoutes)

	v1Routes := `
import { Router } from 'express';
import usersRouter from './users';
const router = Router();
router.use('/users', usersRouter);
export default router;
`
	writeFile(t, dir, "routes/v1.ts", v1Routes)

	appSrc := `
import express from 'express';
import v1Router from './routes/v1';
const app = express();
app.use('/v1', v1Router);
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	expected := []string{
		"GET /v1/users",
		"POST /v1/users",
	}
	for _, e := range expected {
		if !found[e] {
			t.Errorf("missing endpoint %s, got %v", e, found)
		}
	}
}
