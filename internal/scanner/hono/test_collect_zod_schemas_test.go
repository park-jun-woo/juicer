//ff:func feature=scan type=test control=sequence topic=hono
//ff:what Zod 스키마 const 수집 테스트
package hono

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner/zod"
)

func TestCollectZodSchemas(t *testing.T) {
	src := []byte(`
import { z } from "zod"
const createUserSchema = z.object({
  name: z.string(),
  email: z.string().email()
})
const updateUserSchema = z.object({
  name: z.string().optional()
})
const notSchema = "hello"
`)
	fi := mustParse(t, src)
	schemas := zod.CollectSchemas(fi.Root, fi.Src)
	if _, ok := schemas["createUserSchema"]; !ok {
		t.Error("expected createUserSchema")
	}
	if _, ok := schemas["updateUserSchema"]; !ok {
		t.Error("expected updateUserSchema")
	}
	if _, ok := schemas["notSchema"]; ok {
		t.Error("notSchema should not be collected")
	}
}
