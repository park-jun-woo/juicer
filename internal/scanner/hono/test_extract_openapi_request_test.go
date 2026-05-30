//ff:func feature=scan type=test control=sequence topic=hono
//ff:what createRoute request.body zod 스키마가 ValidatorInfo로 추출되는지 테스트
package hono

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestExtractOpenAPIRequest(t *testing.T) {
	src := []byte(`
import { OpenAPIHono, createRoute } from "@hono/zod-openapi"
import { z } from "@hono/zod-openapi"
const app = new OpenAPIHono()
app.openapi(createRoute({
  method: "post",
  path: "/signup",
  request: {
    body: {
      content: {
        "application/json": {
          schema: z.object({ email: z.string(), password: z.string() })
        }
      }
    }
  }
}), signupHandler)
`)
	fi := mustParse(t, src)
	vars := collectHonoVars(fi)
	routes := collectRoutes(fi, vars)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	r := routes[0]
	if len(r.ZodValidators) != 1 {
		t.Fatalf("expected 1 validator, got %d", len(r.ZodValidators))
	}
	if r.ZodValidators[0].Target != "json" {
		t.Errorf("expected json target, got %s", r.ZodValidators[0].Target)
	}
	if r.ZodValidators[0].SchemaNode == nil {
		t.Error("expected inline schema node")
	}
}

// createRouteObj returns the object literal node of the first createRoute({...}) call.
func createRouteObj(t *testing.T, objSrc string) (*fileInfo, *sitter.Node) {
	t.Helper()
	src := []byte("createRoute(" + objSrc + ");\n")
	fi := mustParse(t, src)
	call := findAllByType(fi.Root, "call_expression")[0]
	inner := findChildByType(call, "arguments")
	obj := findChildByType(inner, "object")
	if obj == nil {
		t.Fatal("no object node")
	}
	return fi, obj
}

func TestExtractOpenAPIRequest_NoRequest(t *testing.T) {
	fi, obj := createRouteObj(t, `{ method: "get", path: "/x" }`)
	if v := extractOpenAPIRequest(obj, fi.Src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}

func TestExtractOpenAPIRequest_RequestNotObject(t *testing.T) {
	fi, obj := createRouteObj(t, `{ request: someVar }`)
	if v := extractOpenAPIRequest(obj, fi.Src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}

func TestExtractOpenAPIRequest_QueryAndParams(t *testing.T) {
	fi, obj := createRouteObj(t, `{
  request: {
    query: z.object({ q: z.string() }),
    params: z.object({ id: z.string() })
  }
}`)
	v := extractOpenAPIRequest(obj, fi.Src)
	if len(v) != 2 {
		t.Fatalf("expected 2 validators, got %d: %+v", len(v), v)
	}
	if v[0].Target != "query" || v[1].Target != "param" {
		t.Fatalf("targets: %s %s", v[0].Target, v[1].Target)
	}
}

func TestExtractOpenAPIRequest_SectionNoZodCalls(t *testing.T) {
	// body present but no zod call inside -> skipped
	fi, obj := createRouteObj(t, `{ request: { body: { content: {} } } }`)
	if v := extractOpenAPIRequest(obj, fi.Src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}
