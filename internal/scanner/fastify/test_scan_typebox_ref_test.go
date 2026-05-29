//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TypeBox 변수 참조 body/querystring가 필드까지 채워지는지 검증한다
package fastify

import "testing"

func TestScan_TypeBoxRef(t *testing.T) {
	dir := t.TempDir()
	src := `
import Fastify from "fastify";
import { Type } from "@sinclair/typebox";
const app = Fastify();
const CredentialsSchema = Type.Object({
  email: Type.String({ format: "email" }),
  password: Type.String({ minLength: 8 }),
  remember: Type.Optional(Type.Boolean())
});
const QuerySchema = Type.Object({
  page: Type.Optional(Type.Integer()),
  q: Type.String()
});
app.post("/login", {
  schema: { body: CredentialsSchema, querystring: QuerySchema }
}, loginHandler);
`
	writeFile(t, dir, "app.ts", src)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	req := result.Endpoints[0].Request
	if req == nil || req.Body == nil {
		t.Fatal("expected request body from TypeBox ref")
	}
	if len(req.Body.Fields) != 3 {
		t.Fatalf("expected 3 body fields, got %d", len(req.Body.Fields))
	}
	fm := make(map[string]string)
	rm := make(map[string]string)
	for _, f := range req.Body.Fields {
		fm[f.Name] = f.Type
		rm[f.Name] = f.Validate
	}
	if fm["email"] != "email" {
		t.Errorf("email.Type: want email, got %s", fm["email"])
	}
	if rm["email"] != "required" {
		t.Errorf("email.Validate: want required, got %s", rm["email"])
	}
	if fm["password"] != "string" {
		t.Errorf("password.Type: want string, got %s", fm["password"])
	}
	if rm["remember"] == "required" {
		t.Error("remember should be optional (not required)")
	}
	if fm["remember"] != "boolean" {
		t.Errorf("remember.Type: want boolean, got %s", fm["remember"])
	}
	if len(req.Query) != 2 {
		t.Fatalf("expected 2 query params, got %d", len(req.Query))
	}
	qm := make(map[string]string)
	for _, q := range req.Query {
		qm[q.Name] = q.Type
	}
	if qm["page"] != "integer" {
		t.Errorf("page.Type: want integer, got %s", qm["page"])
	}
	if qm["q"] != "string" {
		t.Errorf("q.Type: want string, got %s", qm["q"])
	}
}
