//ff:func feature=scan type=test control=sequence topic=hono
//ff:what basePath 수집 테스트
package hono

import "testing"

func TestCollectBasePaths(t *testing.T) {
	src := []byte(`
import { Hono } from "hono"
const app = new Hono().basePath("/api/v1")
`)
	fi := mustParse(t, src)
	vars := collectHonoVars(fi)
	bp := collectBasePaths(fi, vars)
	if bp["app"] != "/api/v1" {
		t.Errorf("expected basePath /api/v1, got %s", bp["app"])
	}
}
