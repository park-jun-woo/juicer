//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 3단계 prefix 체이닝 테스트: app→/v1→v1→/auth→auth = /v1/auth
package express

import "testing"

func TestBuildPrefixMap3Level(t *testing.T) {
	dir := t.TempDir()

	authSrc := `
const express = require("express");
const router = express.Router();
router.post("/login", login);
export default router;
`
	writeFile(t, dir, "routes/auth.ts", authSrc)

	v1Src := `
import express from "express";
import authRouter from "./auth";
const router = express.Router();
router.use("/auth", authRouter);
export default router;
`
	writeFile(t, dir, "routes/v1.ts", v1Src)

	appSrc := `
import express from "express";
import v1Router from "./routes/v1";
const app = express();
app.use("/v1", v1Router);
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
	if !found["POST /v1/auth/login"] {
		t.Errorf("expected POST /v1/auth/login, got %v", found)
	}
}
