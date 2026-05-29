//ff:func feature=scan type=test control=sequence topic=hono
//ff:what new OpenAPIHono() 인스턴스를 Hono 변수로 인식하는지 테스트
package hono

import "testing"

func TestIsNewHonoCall_OpenAPIHono(t *testing.T) {
	src := []byte(`
import { OpenAPIHono } from "@hono/zod-openapi"
const app = new OpenAPIHono()
`)
	fi := mustParse(t, src)
	vars := collectHonoVars(fi)
	if !vars["app"] {
		t.Errorf("expected app recognized as Hono var, got %v", vars)
	}
}
