//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 2단계 prefix 체이닝 테스트: app→/api→main→/users→user = /api/users
package express

import "testing"

func TestBuildPrefixMap2Level(t *testing.T) {
	dir := t.TempDir()

	userSrc := `
const express = require("express");
const router = express.Router();
router.get("/", listUsers);
export default router;
`
	writeFile(t, dir, "routes/users.ts", userSrc)

	mainSrc := `
import express from "express";
import userRouter from "./users";
const router = express.Router();
router.use("/users", userRouter);
export default router;
`
	writeFile(t, dir, "routes/index.ts", mainSrc)

	appSrc := `
import express from "express";
import mainRouter from "./routes/index";
const app = express();
app.use("/api", mainRouter);
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
	if !found["GET /api/users"] {
		t.Errorf("expected GET /api/users, got %v", found)
	}
}
