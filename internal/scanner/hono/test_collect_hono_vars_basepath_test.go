//ff:func feature=scan type=test control=sequence topic=hono
//ff:what new Hono().basePath() 체인 인스턴스 변수 수집 테스트
package hono

import "testing"

func TestCollectHonoVars_BasePath(t *testing.T) {
	src := []byte(`
import { Hono } from "hono"
const app = new Hono().basePath("/api")
`)
	fi := mustParse(t, src)
	vars := collectHonoVars(fi)
	if !vars["app"] {
		t.Error("expected app in vars")
	}
}
