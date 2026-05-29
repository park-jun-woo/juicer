//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what prefix가 파일이 아닌 라우터 인스턴스(file,var) 단위로 해석됨을 검증 (Phase097)
package express

import "testing"

func epSet(eps []endpointLike) map[string]bool {
	m := map[string]bool{}
	for _, e := range eps {
		m[e.method+" "+e.path] = true
	}
	return m
}

type endpointLike struct{ method, path string }

func scanSet(t *testing.T, dir string) map[string]bool {
	t.Helper()
	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	var eps []endpointLike
	for _, e := range result.Endpoints {
		eps = append(eps, endpointLike{e.Method, e.Path})
	}
	return epSet(eps)
}

func assertHas(t *testing.T, got map[string]bool, want ...string) {
	t.Helper()
	for _, w := range want {
		if !got[w] {
			t.Errorf("missing endpoint %q; got %v", w, got)
		}
	}
}

// 한 파일에 라우터가 둘이고 서로 다른 prefix로 마운트되면, 각 라우터의
// 라우트는 자기 prefix를 받아야 한다 (파일 단위로 붕괴하면 안 됨).
func TestMultiRouterPerFilePrefix(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/multi.ts", `
import express from "express";
export const aRouter = express.Router();
aRouter.get("/", listA);
export const bRouter = express.Router();
bRouter.get("/", listB);
`)
	writeFile(t, dir, "app.ts", `
import express from "express";
import { aRouter, bRouter } from "./routes/multi";
const app = express();
app.use("/v1/aaa", aRouter);
app.use("/v1/bbb", bRouter);
`)
	writeFile(t, dir, "package.json", `{"dependencies":{"express":"^4.0.0"}}`)
	got := scanSet(t, dir)
	assertHas(t, got, "GET /v1/aaa", "GET /v1/bbb")
}

// 같은 라우터를 두 prefix에 마운트하면 두 경로 모두 산출되어야 한다.
func TestSharedRouterMountedTwice(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/shared.ts", `
import express from "express";
const router = express.Router();
router.get("/", listItems);
export default router;
`)
	writeFile(t, dir, "app.ts", `
import express from "express";
import shared from "./routes/shared";
const app = express();
app.use("/v1/orgs", shared);
app.use("/v1/users", shared);
`)
	writeFile(t, dir, "package.json", `{"dependencies":{"express":"^4.0.0"}}`)
	got := scanSet(t, dir)
	assertHas(t, got, "GET /v1/orgs", "GET /v1/users")
}

// 4단계 + 파라미터 prefix 체인이 끝까지 합성되어야 한다.
func TestDeepParamPrefixChain(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "routes/project.ts", `
import express from "express";
const router = express.Router({mergeParams:true});
router.get("/:projectId/racks", listRacks);
export default router;
`)
	writeFile(t, dir, "routes/orgs.ts", `
import express from "express";
import projects from "./project";
const router = express.Router();
router.use("/:orgId/projects", projects);
export default router;
`)
	writeFile(t, dir, "routes/v1.ts", `
import express from "express";
import orgs from "./orgs";
const router = express.Router();
router.use("/orgs", orgs);
export default router;
`)
	writeFile(t, dir, "app.ts", `
import express from "express";
import v1 from "./routes/v1";
const app = express();
app.use("/v1", v1);
`)
	writeFile(t, dir, "package.json", `{"dependencies":{"express":"^4.0.0"}}`)
	got := scanSet(t, dir)
	assertHas(t, got, "GET /v1/orgs/{orgId}/projects/{projectId}/racks")
}
