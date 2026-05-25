//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildOperation 테스트
package scanner

import (
	"testing"
)

func TestBuildOperation(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		ep := Endpoint{
			Method:  "GET",
			Path:    "/users",
			Handler: "h.GetUsers",
		}
		schemas := map[string]any{}
		op := buildOperation(ep, schemas)
		if op["operationId"] == nil {
			t.Error("expected operationId")
		}
	})

	t.Run("with query params and default", func(t *testing.T) {
		ep := Endpoint{
			Method:  "GET",
			Path:    "/users",
			Handler: "h.GetUsers",
			Request: &Request{
				Query: []Param{
					{Name: "page", Type: "string", Default: "1"},
					{Name: "limit", Type: "string"},
				},
			},
		}
		schemas := map[string]any{}
		op := buildOperation(ep, schemas)
		params, ok := op["parameters"].([]map[string]any)
		if !ok {
			t.Fatal("expected parameters")
		}
		if len(params) != 2 {
			t.Errorf("expected 2 params, got %d", len(params))
		}
	})

	t.Run("with path params", func(t *testing.T) {
		ep := Endpoint{
			Method:  "GET",
			Path:    "/users/:id",
			Handler: "h.GetUser",
			Request: &Request{
				PathParams: []Param{{Name: "id", Type: "string"}},
			},
		}
		schemas := map[string]any{}
		op := buildOperation(ep, schemas)
		params := op["parameters"].([]map[string]any)
		if len(params) != 1 || params[0]["in"] != "path" {
			t.Error("expected path param")
		}
	})

	t.Run("with request body", func(t *testing.T) {
		ep := Endpoint{
			Method:  "POST",
			Path:    "/users",
			Handler: "h.CreateUser",
			Request: &Request{
				Body: &Body{
					TypeName: "CreateReq",
					Fields:   []Field{{Name: "name", Type: "string"}},
				},
			},
		}
		schemas := map[string]any{}
		op := buildOperation(ep, schemas)
		if op["requestBody"] == nil {
			t.Error("expected requestBody")
		}
	})
}
