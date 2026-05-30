//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestApplyValidator_JSON 테스트
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"testing"
)

func TestApplyValidator_JSON(t *testing.T) {
	root, src := parseTS(t, `const userSchema = z.object({ name: z.string() });`)
	schemas := CollectSchemas(root, src)
	req := &scanner.Request{}
	v := ValidatorInfo{Target: "json", SchemaName: "userSchema"}
	ok := ApplyValidator(req, v, schemas, src, map[string][]byte{"userSchema": src})
	if !ok || req.Body == nil || len(req.Body.Fields) != 1 {
		t.Fatalf("got ok=%v body=%+v", ok, req.Body)
	}
}
