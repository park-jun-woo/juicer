//ff:func feature=scan type=test control=sequence topic=express
//ff:what E2E 테스트: 핸들러 본문에서 res.status(N).json() / res.sendStatus(N) 응답 추출
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestScan_ResponseExtraction(t *testing.T) {
	dir := t.TempDir()

	src := `
import { Router } from 'express';
import { validateRequest } from 'zod-express-middleware';
import { z } from 'zod';

const CreateUserSchema = z.object({
  name: z.string(),
  email: z.string().email(),
});

const router = Router();

async function createUser(req, res) {
  const existing = await findUser(req.body.email);
  if (existing) {
    return res.status(409).json({ error: 'User already exists' });
  }
  const user = await saveUser(req.body);
  res.status(201).json({ data: user });
}

function deleteUser(req, res) {
  res.sendStatus(204);
}

router.post('/users', validateRequest({ body: CreateUserSchema }), createUser);
router.delete('/users/:id', deleteUser);
export default router;
`
	writeFile(t, dir, "routes/users.ts", src)

	appSrc := `
import express from "express";
import usersRouter from "./routes/users";
const app = express();
app.use("/api", usersRouter);
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	postEp := findEndpoint(result.Endpoints, "POST", "/api/users")
	if postEp == nil {
		t.Fatalf("POST /api/users not found; endpoints: %v", endpointSummary(result.Endpoints))
	}
	assertResponseSet(t, postEp.Responses, []scanner.Response{
		{Status: "201", Kind: "json"},
		{Status: "409", Kind: "json"},
	})

	delEp := findEndpoint(result.Endpoints, "DELETE", "/api/users/{id}")
	if delEp == nil {
		t.Fatalf("DELETE /api/users/{id} not found; endpoints: %v", endpointSummary(result.Endpoints))
	}
	assertResponseSet(t, delEp.Responses, []scanner.Response{
		{Status: "204", Kind: "empty"},
	})
}
