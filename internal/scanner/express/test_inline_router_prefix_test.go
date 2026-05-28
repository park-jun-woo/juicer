//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what E2E 테스트: 인라인 express.Router() prefix가 자식 라우터에 전파되는지 검증한다
package express

import "testing"

func TestInlineRouterPrefix(t *testing.T) {
	dir := t.TempDir()

	writeFile(t, dir, "routes/orgs.ts", `
import { Router } from 'express';
const router = Router();
router.get('/', (req, res) => { res.json({}); });
router.get('/:orgId', (req, res) => { res.json({}); });
export default router;
`)
	writeFile(t, dir, "routes/projects.ts", `
import { Router } from 'express';
const router = Router({ mergeParams: true });
router.get('/', (req, res) => { res.json({}); });
export default router;
`)
	writeFile(t, dir, "app.ts", `
import express from 'express';
import orgRouter from './routes/orgs';
import projectRouter from './routes/projects';

const app = express();
const v1Router = express.Router();
v1Router.use('/orgs', orgRouter);
v1Router.use('/orgs/:orgId/projects', projectRouter);
app.use('/v1', v1Router);
`)
	writeFile(t, dir, "package.json", `{"dependencies":{"express":"^4.18.0"}}`)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	want := []string{
		"GET /v1/orgs",
		"GET /v1/orgs/{orgId}",
		"GET /v1/orgs/{orgId}/projects",
	}
	got := endpointSummary(result.Endpoints)
	gotSet := make(map[string]bool)
	for _, s := range got {
		gotSet[s] = true
	}
	for _, w := range want {
		if !gotSet[w] {
			t.Errorf("missing endpoint %s; got %v", w, got)
		}
	}
}
