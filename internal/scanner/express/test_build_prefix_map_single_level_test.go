//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 단일 단계 prefix 회귀 테스트: app.use('/users', router) → /users
package express

import "testing"

func TestBuildPrefixMapSingleLevel(t *testing.T) {
	dir := t.TempDir()

	routerSrc := `
const express = require("express");
const router = express.Router();
router.get("/", listUsers);
router.get("/:id", getUser);
export default router;
`
	writeFile(t, dir, "routes/users.ts", routerSrc)

	appSrc := `
import express from "express";
import usersRouter from "./routes/users";
const app = express();
app.use("/users", usersRouter);
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	if !found["GET /users"] {
		t.Errorf("expected GET /users, got %v", found)
	}
	if !found["GET /users/{id}"] {
		t.Errorf("expected GET /users/{id}, got %v", found)
	}
}
