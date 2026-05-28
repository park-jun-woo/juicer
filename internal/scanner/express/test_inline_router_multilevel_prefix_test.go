//ff:func feature=scan type=test control=sequence topic=express
//ff:what E2E 테스트: 다단계 인라인 라우터 prefix가 올바르게 합산되는지 검증한다
package express

import "testing"

func TestInlineRouterMultilevelPrefix(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "routes/users.ts", `
import { Router } from 'express';
const router = Router();
router.get('/', (req, res) => { res.json({}); });
export default router;
`)
	writeFile(t, dir, "app.ts", `
import express from 'express';
import usersRouter from './routes/users';

const app = express();
const apiRouter = express.Router();
const v1Router = express.Router();
v1Router.use('/users', usersRouter);
apiRouter.use('/v1', v1Router);
app.use('/api', apiRouter);
`)
	writeFile(t, dir, "package.json", `{"dependencies":{"express":"^4.18.0"}}`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	ep := findEndpoint(result.Endpoints, "GET", "/api/v1/users")
	if ep == nil {
		t.Errorf("missing GET /api/v1/users; got endpoints:")
		for _, e := range result.Endpoints {
			t.Errorf("  %s %s", e.Method, e.Path)
		}
	}
}
