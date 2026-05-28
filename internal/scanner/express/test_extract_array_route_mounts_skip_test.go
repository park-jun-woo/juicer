//ff:func feature=scan type=test control=sequence topic=express
//ff:what forEach 배열 변수 미발견 시 스킵 테스트: 배열 리터럴이 없으면 에러 없이 빈 결과
package express

import "testing"

func TestExtractArrayRouteMountsSkipMissing(t *testing.T) {
	src := []byte(`
import express from 'express';
const router = express.Router();

// importedRoutes is not declared as a literal in this file
importedRoutes.forEach((r) => {
  router.use(r.path, r.route);
});
`)
	fi := mustParse(t, src)
	routers := collectRouters(fi, nil)
	imports := map[string]string{}

	entries := extractArrayRouteMounts(fi, routers, imports, "test.ts")
	if len(entries) != 0 {
		t.Fatalf("expected 0 entries for missing array, got %d", len(entries))
	}
}
