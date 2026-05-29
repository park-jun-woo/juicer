//ff:func feature=scan type=test control=sequence topic=express
//ff:what 같은 라우터를 두 prefix에 마운트하면 두 경로 모두 산출됨을 검증
package express

import "testing"

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
