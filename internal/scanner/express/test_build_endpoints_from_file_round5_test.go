//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestBuildEndpointsFromFile_Round5 테스트
package express

import "testing"

func TestBuildEndpointsFromFile_Round5(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "r.ts", `
import express from "express";
const router = express.Router();
router.get("/health", listHealth);
export default router;
`)
	res, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Endpoints) == 0 {
		t.Fatalf("expected endpoints from scan, got %d", len(res.Endpoints))
	}

	fi := mustParse(t, []byte(`
const router = express.Router();
router.get("/health", listHealth);
`))
	ctx := round5Ctx()
	eps := buildEndpointsFromFile(fi, map[string]bool{"router": true}, "r.ts", "r.ts", ctx)
	if len(eps) == 0 {
		t.Fatalf("buildEndpointsFromFile: expected endpoints, got %d", len(eps))
	}
}
