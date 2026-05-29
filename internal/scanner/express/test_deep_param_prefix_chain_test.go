//ff:func feature=scan type=test control=sequence topic=express
//ff:what 4단계 + 파라미터 prefix 체인이 끝까지 합성됨을 검증
package express

import "testing"

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
