//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what JSON Schema -> scanner.Field 변환 테스트
package fastify

import "testing"

func TestJSONSchemaToFields(t *testing.T) {
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
	si := extractJSONSchema(routes[0].Schema, fi.Src)
	if si == nil || si.Body == nil {
		t.Fatal("expected body schema")
	}
	fields := jsonSchemaToFields(si.Body, fi.Src)
	if len(fields) != 3 {
		t.Fatalf("expected 3 fields, got %d", len(fields))
	}

	fieldMap := make(map[string]int)
	for i, f := range fields {
		fieldMap[f.Name] = i
	}

	// Check name field
	if idx, ok := fieldMap["name"]; ok {
		if fields[idx].Type != "string" {
			t.Errorf("name.Type: want string, got %s", fields[idx].Type)
		}
		if fields[idx].Validate != "required" {
			t.Errorf("name.Validate: want required, got %s", fields[idx].Validate)
		}
	} else {
		t.Error("missing field: name")
	}

	// Check email field
	if idx, ok := fieldMap["email"]; ok {
		if fields[idx].Type != "email" {
			t.Errorf("email.Type: want email, got %s", fields[idx].Type)
		}
		if fields[idx].Validate != "required" {
			t.Errorf("email.Validate: want required, got %s", fields[idx].Validate)
		}
	} else {
		t.Error("missing field: email")
	}

	// Check age field
	if idx, ok := fieldMap["age"]; ok {
		if fields[idx].Type != "integer" {
			t.Errorf("age.Type: want integer, got %s", fields[idx].Type)
		}
		if fields[idx].Minimum == nil || *fields[idx].Minimum != 0 {
			t.Errorf("age.Minimum: want 0, got %v", fields[idx].Minimum)
		}
	} else {
		t.Error("missing field: age")
	}
}
