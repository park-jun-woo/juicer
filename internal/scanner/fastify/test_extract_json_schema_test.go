//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what JSON Schema 인라인 추출 테스트: schema.body, schema.querystring, schema.response
package fastify

import "testing"

func TestExtractJSONSchema(t *testing.T) {
	src := []byte(`
import Fastify from "fastify";
const app = Fastify();
app.post("/users", {
  schema: {
    body: {
      type: "object",
      required: ["name", "email"],
      properties: {
        name: { type: "string" },
        email: { type: "string", format: "email" },
        age: { type: "integer", minimum: 0 }
      }
    },
    querystring: {
      type: "object",
      properties: {
        page: { type: "integer" },
        limit: { type: "integer" }
      }
    },
    response: {
      200: {
        type: "object",
        properties: {
          id: { type: "string" },
          name: { type: "string" }
        }
      }
    }
  }
}, createUser);
`)
	fi := mustParse(t, src)
	instances := collectInstances(fi)
	routes := extractRoutes(fi, instances)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if r.Schema == nil {
		t.Fatal("expected schema to be non-nil")
	}

	si := extractJSONSchema(r.Schema, fi.Src)
	if si == nil {
		t.Fatal("expected schemaInfo to be non-nil")
	}
	if si.Body == nil {
		t.Error("expected Body to be non-nil")
	}
	if si.Querystring == nil {
		t.Error("expected Querystring to be non-nil")
	}
	if len(si.Response) == 0 {
		t.Error("expected at least one response")
	}
	if _, ok := si.Response["200"]; !ok {
		t.Error("expected 200 response")
	}
}

func TestExtractJSONSchema_NilOpts(t *testing.T) {
	if si := extractJSONSchema(nil, []byte("")); si != nil {
		t.Fatalf("expected nil for nil opts, got %v", si)
	}
}

func TestExtractJSONSchema_NoSchemaKey(t *testing.T) {
	obj, src := firstObject(t, `{ config: { x: 1 } }`)
	if si := extractJSONSchema(obj, src); si != nil {
		t.Fatalf("expected nil when no schema key, got %v", si)
	}
}

func TestExtractJSONSchema_SchemaNotObject(t *testing.T) {
	// schema value is not an object literal -> nil
	obj, src := firstObject(t, `{ schema: true }`)
	if si := extractJSONSchema(obj, src); si != nil {
		t.Fatalf("expected nil when schema not object, got %v", si)
	}
}
